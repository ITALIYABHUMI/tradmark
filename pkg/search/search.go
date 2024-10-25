package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/spf13/viper"
	"github.com/tradmark/config"
	"github.com/tradmark/public/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(db *gorm.DB) (any, error)
	FetchTradsBySerialNumber(db *gorm.DB, searialNumber string) (any, error)
	Search(db *gorm.DB, data string) ([]map[string]interface{}, error)
	UpdateTrademarkVisibility(db *gorm.DB, caseFile *model.CaseFile) error
}

type repo struct{}

var IndexName string = viper.GetString("ES_INDEXNAME")

func PostgresRepo() Repository {
	return &repo{}
}

func (r *repo) Create(db *gorm.DB) (any, error) {

	const batchSize = 500
	offset := 0

	for {
		var caseFiles []model.CaseFile
		if err := db.Order("serial_number").Limit(batchSize).Offset(offset).Find(&caseFiles).Error; err != nil {
			return nil, fmt.Errorf("error fetching data: %v", err)
		}
		if len(caseFiles) == 0 {
			break
		}

		for _, caseFile := range caseFiles {
			data, err := transformDataForElasticsearch(&caseFile)
			if err != nil {
				return nil, err
			}

			req := esapi.IndexRequest{
				Index:      IndexName,
				DocumentID: caseFile.SerialNumber,
				Body:       bytes.NewBuffer(data),
			}

			// Send the request
			resp, err := req.Do(context.Background(), config.EsClient)
			if err != nil {
				return nil, fmt.Errorf("Error getting response: %v", err)
			}

			if resp.IsError() || resp.StatusCode > 299 {
				return resp.String(), fmt.Errorf("Elasticsearch error: %v", resp.String())
			}
		}
		offset += batchSize
	}

	// defer resp.Body.Close()

	return nil, nil
}

func (r *repo) FetchTradsBySerialNumber(db *gorm.DB, searialNumber string) (any, error) {
	var resp map[string]interface{}

	var ctx context.Context

	q := &types.Query{
		Bool: &types.BoolQuery{
			Must: make([]types.Query, 0),
		},
	}

	q.Bool.Must = append(q.Bool.Must, types.Query{
		Match: map[string]types.MatchQuery{
			"serialNumber": {
				Query: searialNumber,
			},
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	search := esapi.SearchRequest{
		Index: []string{IndexName},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"bool" : {
					"must": [
						{
							"match": {
								"serial-number" : "%s"
							}
						},
						{
							"match":{
								"visible": "true"
							}
						}
					]
				}
			}
		}`, searialNumber)),
	}

	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %v", err)
	}

	if response.IsError() || response.StatusCode > 299 {
		return response.String(), fmt.Errorf("Elasticsearch error: %v", response.String())
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	if resp["hits"].(map[string]interface{})["hits"] == "" {
		return nil, nil
	} else {
		for _, hit := range resp["hits"].(map[string]interface{})["hits"].([]interface{}) {
			doc := hit.(map[string]interface{})
			source := doc["_source"]
			return source, nil
		}
	}

	return nil, nil

}

func (r *repo) Search(db *gorm.DB, data string) ([]map[string]interface{}, error) {
	var resp map[string]interface{}
	var ctx context.Context

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Elasticsearch search request
	search := esapi.SearchRequest{
		Index: []string{IndexName},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"bool": {
					"must": [
						{
							"match": {
								"visible" : { 
									"query": "true"
								}
							}
						},
						{
							"multi_match": {
								"query": "%s", 
								"fields": [
									"case-file-header.attorney-name", 
									"case-file-header.employee-name", 
									"case-file-owners.case-file-owner.city", 
									"case-file-owners.case-file-owner.state", 
									"case-file-owners.case-file-owner.country", 
									"case-file-owners.case-file-owners.party-name", 
									"case-file-owners.case-file-owner.nationality.country", 
									"case-file-owners.case-file-owner.nationality.state",
									"serial-number"
								]
							}
						}
					]
				}
			}
		}`, data)),
	}

	// Execute the search request
	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %v", err)
	}
	defer response.Body.Close()

	if response.IsError() || response.StatusCode > 299 {
		return nil, fmt.Errorf("Elasticsearch error: %v", response.String())
	}

	// Decode the response
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	hits := resp["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) == 0 {
		return nil, nil
	}

	var results []map[string]interface{}
	for _, hit := range hits {
		doc := hit.(map[string]interface{})
		source := doc["_source"].(map[string]interface{})
		results = append(results, source)
	}

	return results, nil
}

func (r *repo) UpdateTrademarkVisibility(db *gorm.DB, caseFile *model.CaseFile) error {

	updateBody := map[string]interface{}{
		"doc": caseFile,
	}

	data, err := json.Marshal(updateBody)
	if err != nil {
		return fmt.Errorf("failed to marshal CaseFile: %v", err)
	}

	req := esapi.UpdateRequest{
		Index:      IndexName,
		DocumentID: caseFile.SerialNumber,
		Body:       bytes.NewReader(data),
	}

	resp, err := req.Do(context.Background(), config.EsClient)
	if err != nil {
		return fmt.Errorf("error getting response: %v", err)
	}
	defer resp.Body.Close()

	if resp.IsError() || resp.StatusCode > 299 {
		return fmt.Errorf("elasticsearch error: %v", resp.String())
	}

	return nil
}

func transformDataForElasticsearch(caseFile *model.CaseFile) ([]byte, error) {
	// Marshal caseFile to JSON
	data, err := json.Marshal(caseFile)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CaseFile: %v", err)
	}
	return data, nil
}

package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/spf13/viper"
	"github.com/tradmark/api/model"
	"github.com/tradmark/config"
	"gorm.io/gorm"
)

type repo struct{}

var IndexName string = viper.GetString("ES_INDEXNAME")

func (r *repo) Create(db *gorm.DB) (any, error) {

	const batchSize = 500
	offset := 0

	for {
		var caseFiles []model.CaseFile
		if err := db.Limit(batchSize).Offset(offset).Find(&caseFiles).Error; err != nil {
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
			_, err = req.Do(context.Background(), config.EsClient)
			if err != nil {
				return nil, fmt.Errorf("Error getting response: %v", err)
			}

		}
		offset += batchSize
	}

	return nil, nil
}

func (r *repo) FetchTradsBySerialNumber(db *gorm.DB, serialNumber string) (any, error) {
	var resp types.ResponseBody
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	q := &types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{
				{
					Match: map[string]types.MatchQuery{
						"serial-number": {
							Query: serialNumber,
						},
					},
				},
				{
					Match: map[string]types.MatchQuery{
						"visible": {
							Query: "true",
						},
					},
				},
			},
		},
	}

	body, err := json.Marshal(map[string]interface{}{"query": q})
	if err != nil {
		return nil, fmt.Errorf("error marshaling query to JSON: %v", err)
	}

	search := esapi.SearchRequest{
		Index: []string{IndexName},
		Body:  strings.NewReader(string(body)),
	}

	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %v", err)
	}
	defer response.Body.Close()

	if response.IsError() {
		errorBody, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("Elasticsearch error: %s", string(errorBody))
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	if len(resp.Hits.Hits) == 0 {
		return nil, nil
	}

	return resp.Hits.Hits[0].Source_, nil
}

func (r *repo) Search(db *gorm.DB, data string) (any, error) {
	var resp types.ResponseBody
	var ctx context.Context

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Elasticsearch search request
	query := types.Query{
		Bool: &types.BoolQuery{
			Must: []types.Query{
				{
					Match: map[string]types.MatchQuery{
						"visible": {
							Query: "true",
						},
					},
				},
				{
					MultiMatch: &types.MultiMatchQuery{
						Query: data,
						Fields: []string{
							"serial-number",
							"case-file-header.attorney-name",
							"case-file-header.employee-name",
							"case-file-owners.case-file-owner.city",
							"case-file-owners.case-file-owner.state",
							"case-file-owners.case-file-owner.country",
							"case-file-owners.case-file-owner.party-name",
							"case-file-owners.case-file-owner.nationality.country",
							"case-file-owners.case-file-owner.nationality.state",
						},
					},
				},
			},
		},
	}

	body, err := json.Marshal(map[string]interface{}{"query": query})
	if err != nil {
		return nil, fmt.Errorf("error marshaling query to JSON: %v", err)
	}

	search := esapi.SearchRequest{
		Index: []string{IndexName},
		Body:  strings.NewReader(string(body)),
	}

	// Execute the search request
	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %v", err)
	}
	defer response.Body.Close()

	// Decode the response
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	if len(resp.Hits.Hits) == 0 {
		return nil, nil
	}

	var sources []json.RawMessage

	for _, hit := range resp.Hits.Hits {
		sources = append(sources, hit.Source_)
	}

	return sources, nil

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

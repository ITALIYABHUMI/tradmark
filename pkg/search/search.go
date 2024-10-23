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
	"github.com/tradmark/config"
	"github.com/tradmark/public/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(db *gorm.DB) (any, error)
	FetchTradsBySerialNumber(db *gorm.DB, searialNumber string) (any, error)
	Search(db *gorm.DB, data string) (any, error)
}

type repo struct{}

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
			// Prepare Elasticsearch IndexRequest
			req := esapi.IndexRequest{
				Index:      "tradmark",
				DocumentID: caseFile.SerialNumber,
				Body:       bytes.NewBuffer(data),
			}

			// Send the request
			resp, err := req.Do(context.Background(), config.EsClient)
			if err != nil {
				return nil, fmt.Errorf("Error getting response: %v", err)
			}

			if resp.IsError() || resp.StatusCode > 299 {
				fmt.Println("resp.String()", resp.String())
				fmt.Println(resp.IsError())

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
		Index: []string{"tradmark"},
		Body: strings.NewReader(fmt.Sprintf(`{
			 "query": {
				  "match": {
				  		"serial-number" : "%s"
				  }
			 },
			 "from": 0,
			 "size": 10
		}`, searialNumber)),
	}

	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %v", err)
	}

	if response.IsError() || response.StatusCode > 299 {
		fmt.Println("response.String()", response.String())
		fmt.Println(response.IsError())

		return response.String(), fmt.Errorf("Elasticsearch error: %v", response.String())
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	for _, hit := range resp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		return source, nil
	}

	return resp, nil
}

func (r *repo) Search(db *gorm.DB, data string) (any, error) {
	var resp map[string]interface{}

	var ctx context.Context

	q := &types.Query{
		Bool: &types.BoolQuery{
			Must: make([]types.Query, 0),
		},
	}

	q.Bool.Must = append(q.Bool.Must, types.Query{
		Match: map[string]types.MatchQuery{
			"serialNumber":     {Query: data},
			"description-text": {Query: data},
			"attorney-name":    {Query: data},
			"employee-name":    {Query: data},
			"text":             {Query: data},
			"address-1":        {Query: data},
			"address-2":        {Query: data},
			"address-3":        {Query: data},
			"address-4":        {Query: data},
			"address-5":        {Query: data},
			"city":             {Query: data},
			"party-name":       {Query: data},
			"country":          {Query: data},
			"entity-statement": {Query: data},
			"serial-number":    {Query: data},
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// search := esapi.SearchRequest{
	// 	Index: []string{"tradmark"},
	// 	Body: strings.NewReader(fmt.Sprintf(`{
	// 		 "query": {
	// 			  "match": {
	// 			  		"description-text" : "%s",
	// 					"attorney-name" : "%s",
	// 					"employee-name" : "%s",
	// 					"text" : "%s",
	// 					"description-text" : "%s",
	// 					"address-1": "%s",
	// 					"address-2": "%s",
	// 					"address-3":  "%s",
	// 					"address-4":  "%s",
	// 					"address-5":  "%s",
	// 					"city" ; "%s",
	// 					"party-name" : "%s",
	// 					"country" : "%s",
	// 					"entity-statement": "%s",
	// 			  }
	// 		 },
	// 		 "from": 0,
	// 		 "size": 10
	// 	}`, data)),
	// }

	search := esapi.SearchRequest{
		Index: []string{"tradmark"},
		Body: strings.NewReader(fmt.Sprintf(`{
				"query": {
					"multi_match": {
						"query": "%s", 
						"fields": ["description-text", "attorney-name", "city"]
					}
				}
			}`, data)),
	}

	response, err := search.Do(ctx, config.EsClient)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %v", err)
	}

	if response.IsError() || response.StatusCode > 299 {
		fmt.Println("response.String()", response.String())
		fmt.Println(response.IsError())

		return response, fmt.Errorf("Elasticsearch error: %v", response.String())
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	for _, hit := range resp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		return source, nil
	}

	return resp, nil
}

func transformDataForElasticsearch(caseFile *model.CaseFile) ([]byte, error) {
	// Marshal caseFile to JSON
	data, err := json.Marshal(caseFile)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CaseFile: %v", err)
	}
	return data, nil
}

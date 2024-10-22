package config

import (
	"context"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var EsClient *elasticsearch.Client
var SearchIndex = "tradmark"

func EsClientConnection() {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	EsClient = esClient
}

func EsCreateIndexIfNotExists() {
	_, err := esapi.IndicesExistsRequest{
		Index: []string{SearchIndex},
	}.Do(context.Background(), EsClient)

	if err != nil {
		EsClient.Indices.Create(SearchIndex)
	}
}
	
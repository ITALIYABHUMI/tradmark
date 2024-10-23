package config

import (
	"context"
	"fmt"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/spf13/viper"
)

var EsClient *elasticsearch.Client
var SearchIndex = "tradmark"

func EsCreateIndexIfNotExists() {

	if EsClient == nil {
		log.Fatal("Elasticsearch client is not initialized")
		return
	}

	_, err := esapi.IndicesExistsRequest{
		Index: []string{SearchIndex},
	}.Do(context.Background(), EsClient)

	if err != nil {
		_, err := EsClient.Indices.Create(SearchIndex)
		if err != nil {
			log.Fatalf("Error creating index: %s", err)
		} else {
			fmt.Printf("Created index %s\n", SearchIndex)
		}
	} else {
		fmt.Printf("Index %s already exists\n", SearchIndex)
	}
}

func EsClientConnection() {
	Username := viper.GetString("ES_USERNAME")
	password := viper.GetString("ES_PASSWORD")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: Username,
		Password: password,
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(fmt.Sprintf("Error creating Elasticsearch client: %s", err))
	}

	// Test the connection
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.Status())
	}

	fmt.Println("Connected to Elasticsearch")
	EsClient = esClient
}

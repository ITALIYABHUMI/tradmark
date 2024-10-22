package tradmark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tradmark/config"
	"github.com/tradmark/public/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error
	FetchTrads(db *gorm.DB) ([]model.CaseFile, error)
	FetchTradsBySerialNumber(db *gorm.DB, serialNumber string) (model.CaseFile, error)
}

type repo struct{}

func PostgresRepo() Repository {
	return &repo{}
}

func (r *repo) CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error {
	return db.Create(caseFile).Error
}

func (r *repo) FetchTrads(db *gorm.DB) ([]model.CaseFile, error) {
	var caseFiles []model.CaseFile

	err := db.Limit(5).Find(&caseFiles).Error
	if err != nil {
		return nil, err
	}
	return caseFiles, nil
}

func (r *repo) FetchTradsBySerialNumber(db *gorm.DB, serialNumber string) (model.CaseFile, error) {
	var caseFile model.CaseFile
	var buf bytes.Buffer
	var b map[string]interface{}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": serialNumber,
				"field": caseFile.TransactionDate,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return model.CaseFile{}, err
	}

	resp, err := config.EsClient.Search(
		config.EsClient.Search.WithIndex(config.SearchIndex),
		config.EsClient.Search.WithBody(&buf),
	)
	fmt.Println(resp.IsError())
	if err != nil || resp.IsError() {
		return model.CaseFile{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return model.CaseFile{}, err
	}

	var id string
	if hits, ok := b["hits"].(map[string]interface{}); ok {
		if hitsHits, ok := hits["hits"].([]interface{}); ok {
			for _, hit := range hitsHits {
				if hitMap, ok := hit.(map[string]interface{}); ok {
					if idStr, ok := hitMap["serial_number"].(string); ok {
						fmt.Println("idStr:", idStr)
					} else {
						fmt.Println("serial_number not found or not a string")
					}
				}
			}
		}
	}
	err = db.Where("serial_number = ?", id).First(&caseFile).Error
	if err != nil {
		return model.CaseFile{}, err
	}
	return caseFile, nil
}

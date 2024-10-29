package search

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/tradmark/api/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(db *gorm.DB) (any, error)
	FetchTradsBySerialNumber(db *gorm.DB, searialNumber string) (types.ResponseBody, error)
	Search(db *gorm.DB, data string) (types.ResponseBody, error)
	UpdateTrademarkVisibility(db *gorm.DB, caseFile *model.CaseFile) error
}

func PostgresRepo() Repository {
	return &repo{}
}

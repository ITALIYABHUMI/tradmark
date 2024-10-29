package search

import (
	"github.com/tradmark/api/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(db *gorm.DB) (any, error)
	FetchTradsBySerialNumber(db *gorm.DB, searialNumber string) (any, error)
	Search(db *gorm.DB, data string) (any, error)
	UpdateTrademarkVisibility(db *gorm.DB, caseFile *model.CaseFile) error
}

func PostgresRepo() Repository {
	return &repo{}
}

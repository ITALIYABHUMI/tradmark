package tradmark

import (
	"github.com/tradmark/api/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error
	UpdateTrademarkVisibility(db *gorm.DB, serialNumber string, visible string) (*model.CaseFile, error)
}


func PostgresRepo() Repository {
	return &repo{}
}


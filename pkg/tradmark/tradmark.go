package tradmark

import (
	"github.com/tradmark/public/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error
}

type repo struct{}

func PostgresRepo() Repository {
	return &repo{}
}

func (r *repo) CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error {
	return db.Create(caseFile).Error
}

package tradmark

import (
	"github.com/tradmark/api/model"
	"gorm.io/gorm"
)

type repo struct {}

func (r *repo) CreateCaseFiles(db *gorm.DB, caseFile *model.CaseFile) error {
	return db.Create(caseFile).Error
}

func (r *repo) UpdateTrademarkVisibility(db *gorm.DB, serialNumber string, visible string) (*model.CaseFile, error) {

	var caseFile model.CaseFile

	if err := db.Model(&model.CaseFile{}).Where("serial_number = ?", serialNumber).Update("visible", visible).Error; err != nil {
		return nil, err
	}

	err := db.Where("serial_number = ?", serialNumber).First(&caseFile).Error
	if err != nil {
		return nil, err
	}

	return &caseFile, nil
}

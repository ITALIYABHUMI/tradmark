package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

// type CaseFileEventStatements struct {
// 	CaseFileEventStatement []CaseFileEventStatement `json:"case-file-event-statement" gorm:"foreignKey:ID"`
// }

func (n CaseFileEventStatements) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileEventStatements) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFileEventStatement struct {
	Code            string `json:"code"`
	Type            string `json:"type"`
	DescriptionText string `json:"description-text"`
	Date            string `json:"date"`
	Number          string `json:"number"`
}

func (n CaseFileEventStatement) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileEventStatement) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

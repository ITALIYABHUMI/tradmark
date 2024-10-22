package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type CaseFileStatements struct {
	CaseFileStatement CaseFileStatement `json:"case_file_statement"`
}

func (n CaseFileStatements) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileStatements) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFileStatement struct {
	TypeCode string `json:"type_code"`
	Text     string `json:"text"`
}

func (n CaseFileStatement) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileStatement) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

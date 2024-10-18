package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type Classifications struct {
	Classification Classification `json:"classification"`
}

func (n Classifications) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Classifications) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type Classification struct {
	InternationalCodeTotalNo int   `json:"international_code_total_no"`
	UsCodeTotalNo            int   `json:"us_code_total_no"`
	InternationalCode        int   `json:"international_code"`
	UsCode                   []int `json:"us_code"`
	StatusCode               int   `json:"status_code"`
	StatusDate               int64 `json:"status_date"`
	PrimaryCode              int   `json:"primary_code"`
}

func (n Classification) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Classification) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

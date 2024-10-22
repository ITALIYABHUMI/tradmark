package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

// type Classifications struct {
// 	Classification []Classification `json:"classification" gorm:"foreignKey:ID"`
// }

func (n Classifications) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Classifications) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type Classification struct {
	InternationalCodeTotalNo string   `json:"international-code-total-no"`
	UsCodeTotalNo            string   `json:"us-code-total-no"`
	InternationalCode        string   `json:"international-code"`
	StatusCode               string   `json:"status-code"`
	StatusDate               string   `json:"status-date"`
	PrimaryCode              string   `json:"primary-code"`
	UsCode                   []string `json:"us-code"`
}

func (n Classification) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Classification) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

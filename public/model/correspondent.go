package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type Correspondent struct {
	Address1 string `json:"address-1"`
	Address2 string `json:"address-2"`
	Address3 string `json:"address-3"`
	Address4 string `json:"address-4"`
	Address5 string `json:"address-5"`
}

func (n Correspondent) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Correspondent) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

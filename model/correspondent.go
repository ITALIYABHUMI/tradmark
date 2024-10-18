package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type Correspondent struct {
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	Address3 string `json:"address_3"`
	Address4 string `json:"address_4"`
	Address5 string `json:"address_5"`
}

func (n Correspondent) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Correspondent) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

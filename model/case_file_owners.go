package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

// type CaseFileOwners struct {
// 	CaseFileOwner CaseFileOwner `json:"case-file-owner" gorm:"foreignKey:ID"`
// }

func (n CaseFileOwners) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileOwners) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFileOwner struct {
	EntryNumber         string      `json:"entry-number"`
	PartyType           string      `json:"party-type"`
	Nationality         Nationality `json:"nationality"`
	LegalEntityTypeCode string      `json:"legal-entity-type-code"`
	PartyName           string      `json:"party-name"`
	Address1            string      `json:"address-1"`
	City                string      `json:"city"`
	Country             string      `json:"country"`
	Postcode            string      `json:"postcode"`
}

func (n CaseFileOwner) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileOwner) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type Nationality struct {
	Country string `json:"country"`
}

func (n Nationality) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Nationality) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

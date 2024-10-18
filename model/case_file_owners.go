package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type CaseFileOwners struct {
	CaseFileOwner CaseFileOwner `json:"case_file_owner"`
}

func (n CaseFileOwners) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileOwners) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFileOwner struct {
	EntryNumber         int         `json:"entry_number"`
	PartyType           int         `json:"party_type"`
	Nationality         Nationality `json:"nationality"`
	LegalEntityTypeCode int         `json:"legal_entity_type_code"`
	PartyName           string      `json:"party_name"`
	Address1            string      `json:"address_1"`
	City                string      `json:"city"`
	Country             string      `json:"country"`
	Postcode            int         `json:"postcode"`
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

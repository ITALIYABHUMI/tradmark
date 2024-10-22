package model

import (
	"encoding/json"
	"fmt"
)

type CaseFileOwners []CaseFileOwner

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

type Nationality struct {
	Country string `json:"country"`
}

func (c *CaseFileOwners) UnmarshalJSON(data []byte) error {
	// Try unmarshaling into a slice first
	var caseFileOwners []CaseFileOwner
	if err := json.Unmarshal(data, &caseFileOwners); err == nil {
		*c = caseFileOwners
		return nil
	}

	// If it's not a slice, try unmarshaling into a single object
	var singleCaseFileOwner CaseFileOwner
	if err := json.Unmarshal(data, &singleCaseFileOwner); err == nil {
		*c = []CaseFileOwner{singleCaseFileOwner}
		return nil
	}

	// Return error if neither worked
	return fmt.Errorf("failed to unmarshal CaseFileOwners: %s", data)
}

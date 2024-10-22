package model

import (
	"encoding/json"
	"fmt"
)

type CaseFileEventStatements []CaseFileEventStatement

type CaseFileEventStatement struct {
	Code            string `json:"code"`
	Type            string `json:"type"`
	DescriptionText string `json:"description-text"`
	Date            string `json:"date"`
	Number          string `json:"number"`
}

func (c *CaseFileEventStatements) UnmarshalJSON(data []byte) error {
	// Try unmarshaling into a slice first
	var CaseFileEventStatements []CaseFileEventStatement
	if err := json.Unmarshal(data, &CaseFileEventStatements); err == nil {
		*c = CaseFileEventStatements
		return nil
	}

	// If it's not a slice, try unmarshaling into a single object
	var singleCaseFileEventStatements CaseFileEventStatement
	if err := json.Unmarshal(data, &singleCaseFileEventStatements); err == nil {
		*c = []CaseFileEventStatement{singleCaseFileEventStatements}
		return nil
	}

	// Return error if neither worked
	return fmt.Errorf("failed to unmarshal Classifications: %s", data)
}

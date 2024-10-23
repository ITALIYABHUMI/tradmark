package model

import (
	"encoding/json"
	"fmt"
)

type Classifications []Classification

type Classification struct {
	InternationalCodeTotalNo string   `json:"international-code-total-no"`
	UsCodeTotalNo            string   `json:"us-code-total-no"`
	InternationalCode        string   `json:"international-code"`
	StatusCode               string   `json:"status-code"`
	StatusDate               string   `json:"status-date"`
	PrimaryCode              string   `json:"primary-code"`
	UsCode                   []string `json:"us-code"`
}

func (c *Classifications) UnmarshalJSON(data []byte) error {
	// Try unmarshaling into a slice first
	var classifications []Classification
	if err := json.Unmarshal(data, &classifications); err == nil {
		*c = classifications
		return nil
	}

	// If it's not a slice, try unmarshaling into a single object
	var singleClassification Classification
	if err := json.Unmarshal(data, &singleClassification); err == nil {
		*c = []Classification{singleClassification}
		return nil
	}

	return fmt.Errorf("failed to unmarshal Classifications: %s", data)
}

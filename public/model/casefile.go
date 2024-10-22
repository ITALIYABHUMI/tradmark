package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
	"gorm.io/datatypes"
)

type TrademarkApplicationsDailyWrapper struct {
	TrademarkApplicationsDaily TrademarkApplicationsDaily `json:"trademark-applications-daily"`
}

func (n TrademarkApplicationsDailyWrapper) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *TrademarkApplicationsDailyWrapper) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type TrademarkApplicationsDaily struct {
	Version                Version                `json:"version"`
	CreationDateTime       string                 `json:"creation-datetime"`
	ApplicationInformation ApplicationInformation `json:"application-information"`
}

func (n TrademarkApplicationsDaily) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *TrademarkApplicationsDaily) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type Version struct {
	VersionNo   string `json:"version-no"`
	VersionDate string `json:"version-date"`
}

func (n Version) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Version) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type ApplicationInformation struct {
	FileSegments FileSegments `json:"file-segments"`
}

func (n ApplicationInformation) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *ApplicationInformation) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type FileSegments struct {
	FileSegment string       `json:"file-segment"`
	ActionKeys  []ActionKeys `json:"action-keys"`
}

func (n FileSegments) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *FileSegments) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type ActionKeys struct {
	ActionKey string     `json:"action-key"`
	CaseFile  []CaseFile `json:"case-file"`
}

func (n ActionKeys) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *ActionKeys) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFile struct {
	SerialNumber            string         `json:"serial-number"`
	RegistrationNumber      string         `json:"registration-number"`
	TransactionDate         string         `json:"transaction-date"`
	CaseFileHeader          datatypes.JSON `json:"case-file-header"`
	CaseFileStatements      datatypes.JSON `json:"case-file-statements"`
	CaseFileEventStatements datatypes.JSON `json:"case-file-event-statements"`
	Classifications         datatypes.JSON `json:"classifications"`
	Correspondent           datatypes.JSON `json:"correspondent"`
	CaseFileOwners          datatypes.JSON `json:"case-file-owners"`
	Visible                 bool           `json:"visible" gorm:"default:true"`
}

func (n CaseFile) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFile) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

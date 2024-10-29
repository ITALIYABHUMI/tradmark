package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
	"gorm.io/datatypes"
)

type TrademarkApplicationsDailyWrapper struct {
	TrademarkApplicationsDaily TrademarkApplicationsDaily `json:"trademark-applications-daily"  xml:"<trademark-applications-daily>"`
}

func (n TrademarkApplicationsDailyWrapper) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *TrademarkApplicationsDailyWrapper) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type TrademarkApplicationsDaily struct {
	Version                Version                `json:"version" xml:"<version>"`
	CreationDateTime       string                 `json:"creation-datetime" xml:"<creation-datetime>"`
	ApplicationInformation ApplicationInformation `json:"application-information" xml:"<application-information>"`
}

func (n TrademarkApplicationsDaily) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *TrademarkApplicationsDaily) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type Version struct {
	VersionNo   string `json:"version-no" xml:"<version-no>"`
	VersionDate string `json:"version-date" xml:"<version-date>"`
}

func (n Version) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *Version) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type ApplicationInformation struct {
	FileSegments FileSegments `json:"file-segments" xml:"<file-segments>"`
}

func (n ApplicationInformation) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *ApplicationInformation) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type FileSegments struct {
	FileSegment string       `json:"file-segment" xml:"<file-segment>"`
	ActionKeys  []ActionKeys `json:"action-keys" xml:"<action-keys>"`
}

func (n FileSegments) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *FileSegments) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type ActionKeys struct {
	ActionKey string     `json:"action-key" xml:"<action-key>"`
	CaseFile  []CaseFile `json:"case-file" xml:"<case-file>"`
}

func (n ActionKeys) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *ActionKeys) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

type CaseFile struct {
	SerialNumber            string         `json:"serial-number" xml:"<serial-number>" gorm:"primarykey"`
	RegistrationNumber      string         `json:"registration-number" xml:"<registration-number>"`
	TransactionDate         string         `json:"transaction-date" xml:"<transaction-date>"`
	CaseFileHeader          datatypes.JSON `json:"case-file-header" xml:"<case-file-header>"`
	CaseFileStatements      datatypes.JSON `json:"case-file-statements" xml:"<case-file-statements>"`
	CaseFileEventStatements datatypes.JSON `json:"case-file-event-statements" xml:"<case-file-event-statements>"`
	Classifications         datatypes.JSON `json:"classifications" xml:"<classifications>"`
	Correspondent           datatypes.JSON `json:"correspondent" xml:"<correspondent>"`
	CaseFileOwners          datatypes.JSON `json:"case-file-owners" xml:"<case-file-owners>"`
	Visible                 bool           `json:"visible" gorm:"default:true" xml:"<visible>"`
}

func (n CaseFile) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFile) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

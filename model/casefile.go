package model

type TrademarkApplicationsDaily struct {
	Version                Version                `json:"version"`
	CreationDateTime       int64                  `json:"creation-datetime"`
	ApplicationInformation ApplicationInformation `json:"application-information"`
}

type Version struct {
	VersionNo   int   `json:"version-no"`
	VersionDate int64 `json:"version-date"`
}

type ApplicationInformation struct {
	FileSegments FileSegments `json:"file-segments"`
}

type FileSegments struct {
	FileSegment string     `json:"file-segment"`
	ActionKeys  ActionKeys `json:"action-keys"`
}

type ActionKeys struct {
	ActionKey string     `json:"action-key"`
	CaseFile  []CaseFile `json:"case-file"`
}

type CaseFile struct {
	SerialNumber            int64                   `json:"serial-number"`
	RegistrationNumber      int64                   `json:"registration-number"`
	TransactionDate         int64                   `json:"transaction-date"`
	CaseFileHeader          CaseFileHeader          `json:"case-file-header"`
	CaseFileStatements      CaseFileStatements      `json:"case-file-statements"`
	CaseFileEventStatements CaseFileEventStatements `json:"case-file-event-statements"`
	Classifications         Classifications         `json:"classifications"`
	Correspondent           Correspondent           `json:"correspondent"`
	CaseFileOwners          CaseFileOwners          `json:"case-file-owners"`
	Visible                 bool                    `json:"visible"`
}

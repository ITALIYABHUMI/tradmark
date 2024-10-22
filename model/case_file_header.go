package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type CaseFileHeader struct {
	FilingDate                    string `json:"filing-date"`
	StatusCode                    string `json:"status-code"`
	StatusDate                    string `json:"status-date"`
	MarkIdentification            string `json:"mark-identification"`
	MarkDrawingCode               string `json:"mark-drawing-code"`
	AttorneyDocketNumber          string `json:"attorney-docket-number"`
	AttorneyName                  string `json:"attorney-name"`
	PrincipalRegisterAmendedIn    string `json:"principal-register-amended-in"`
	SupplementalRegisterAmendedIn string `json:"supplemental-register-amended-in"`
	SrademarkIn                   string `json:"srademark-in"`
	CollectiveTrademarkIn         string `json:"collective-trademark-in"`
	ServiceMarkIn                 string `json:"service-mark-in"`
	CollectiveServiceMarkIn       string `json:"collective-service-mark-in"`
	CollectiveMembershipMarkIn    string `json:"collective-membership-mark-in"`
	CertificationMarkIn           string `json:"certification-mark-in"`
	CancellationPendingIn         string `json:"cancellation-pending-in"`
	PublishedConcurrentIn         string `json:"published-concurrent-in"`
	ConcurrentUseIn               string `json:"concurrent-use-in"`
	ConcurrentUseProceedingIn     string `json:"concurrent-use-proceeding-in"`
	InterferencePendingIn         string `json:"interference-pending-in"`
	OppositionPendingIn           string `json:"opposition-pending-in"`
	Section12cIn                  string `json:"section-12c-in"`
	Section2fIn                   string `json:"section-2f-in"`
	Section2fInPartIn             string `json:"section-2f-in-part-in"`
	RenewalFiledIn                string `json:"renewal-filed-in"`
	Section8FiledIn               string `json:"section-8-filed-in"`
	Section8PartialAcceptIn       string `json:"section-8-partial-accept-in"`
	Section8AcceptedIn            string `json:"section-8-accepted-in"`
	Section15AcknowledgedIn       string `json:"section-15-acknowledged-in"`
	Section15FiledIn              string `json:"section-15-filed-in"`
	SupplementalRegisterIn        string `json:"supplemental-register-in"`
	ForeignPriorityIn             string `json:"foreign-priority-in"`
	ChangeRegistrationIn          string `json:"change-registration-in"`
	IntentToUseIn                 string `json:"intent-to-use-in"`
	IntentToUseCurrentIn          string `json:"intent-to-use-current-in"`
	FiledAsUseApplicationIn       string `json:"filed-as-use-application-in"`
	AmendedToUseApplicationIn     string `json:"amended-to-use-application-in"`
	UseApplicationCurrentlyIn     string `json:"use-application-currently-in"`
	AmendedToITUApplicationIn     string `json:"amended-to-itu-application-in"`
	FilingBasisFiledAs44dIn       string `json:"filing-basis-filed-as-44d-in"`
	AmendedTo44dApplicationIn     string `json:"amended-to-44d-application-in"`
	FilingBasisCurrent44dIn       string `json:"filing-basis-current-44d-in"`
	FilingBasisFiledAs44eIn       string `json:"filing-basis-filed-as-44e-in"`
	FilingBasisCurrent44eIn       string `json:"filing-basis-current-44e-in"`
	AmendedTo44eApplicationIn     string `json:"amended-to-44e-application-in"`
	WithoutBasisCurrentlyIn       string `json:"without-basis-currently-in"`
	FilingCurrentNoBasisIn        string `json:"filing-current-no-basis-in"`
	ColorDrawingFiledIn           string `json:"color-drawing-filed-in"`
	ColorDrawingCurrentIn         string `json:"color-drawing-current-in"`
	Drawing3DFiledIn              string `json:"drawing-3d-filed-in"`
	Drawing3DCurrentIn            string `json:"drawing-3d-current-in"`
	StandardCharactersClaimedIn   string `json:"standard-characters-claimed-in"`
	FilingBasisFiledAs66aIn       string `json:"filing-basis-filed-as-66a-in"`
	FilingBasisCurrent66aIn       string `json:"filing-basis-current-66a-in"`
	LawOfficeAssignedLocationCode string `json:"law-office-assigned-location-code"`
	CurrentLocation               string `json:"current-location"`
	LocationDate                  string `json:"location-date"`
	EmployeeName                  string `json:"employee-name"`
}

func (n CaseFileHeader) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileHeader) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

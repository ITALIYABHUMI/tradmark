package model

import (
	"database/sql/driver"

	"github.com/tradmark/common"
)

type CaseFileHeader struct {
	ID                            uint   `gorm:"primaryKey"`
	FilingDate                    int64  `json:"filing_date"`
	StatusCode                    int    `json:"status_code"`
	StatusDate                    int64  `json:"status_date"`
	MarkIdentification            string `json:"mark_identification"`
	MarkDrawingCode               int64  `json:"mark_drawing_code"`
	AttorneyDocketNumber          string `json:"attorney_docket_number"`
	AttorneyName                  string `json:"attorney_name"`
	PrincipalRegisterAmendedIn    string `json:"principal_register_amended_in"`
	SupplementalRegisterAmendedIn string `json:"supplemental_register_amended_in"`
	SrademarkIn                   string `json:"srademark_in"`
	CollectiveTrademarkIn         string `json:"collective_trademark_in"`
	ServiceMarkIn                 string `json:"service_mark_in"`
	CollectiveServiceMarkIn       string `json:"collective_service_mark_in"`
	CollectiveMembershipMarkIn    string `json:"collective_membership_mark_in"`
	CertificationMarkIn           string `json:"certification_mark_in"`
	CancellationPendingIn         string `json:"cancellation_pending_in"`
	PublishedConcurrentIn         string `json:"published_concurrent_in"`
	ConcurrentUseIn               string `json:"concurrent_use_in"`
	ConcurrentUseProceedingIn     string `json:"concurrent_use_proceeding_in"`
	InterferencePendingIn         string `json:"interference_pending_in"`
	OppositionPendingIn           string `json:"opposition_pending_in"`
	Section12cIn                  string `json:"section_12c_in"`
	Section2fIn                   string `json:"section_2f_in"`
	Section2fInPartIn             string `json:"section_2f_in_part_in"`
	RenewalFiledIn                string `json:"renewal_filed_in"`
	Section8FiledIn               string `json:"section_8_filed_in"`
	Section8PartialAcceptIn       string `json:"section_8_partial_accept_in"`
	Section8AcceptedIn            string `json:"section_8_accepted_in"`
	Section15AcknowledgedIn       string `json:"section_15_acknowledged_in"`
	Section15FiledIn              string `json:"section_15_filed_in"`
	SupplementalRegisterIn        string `json:"supplemental_register_in"`
	ForeignPriorityIn             string `json:"foreign_priority_in"`
	ChangeRegistrationIn          string `json:"change_registration_in"`
	IntentToUseIn                 string `json:"intent_to_use_in"`
	IntentToUseCurrentIn          string `json:"intent_to_use_current_in"`
	FiledAsUseApplicationIn       string `json:"filed_as_use_application_in"`
	AmendedToUseApplicationIn     string `json:"amended_to_use_application_in"`
	UseApplicationCurrentlyIn     string `json:"use_application_currently_in"`
	AmendedToITUApplicationIn     string `json:"amended_to_itu_application_in"`
	FilingBasisFiledAs44dIn       string `json:"filing_basis_filed_as_44d_in"`
	AmendedTo44dApplicationIn     string `json:"amended_to_44d_application_in"`
	FilingBasisCurrent44dIn       string `json:"filing_basis_current_44d_in"`
	FilingBasisFiledAs44eIn       string `json:"filing_basis_filed_as_44e_in"`
	FilingBasisCurrent44eIn       string `json:"filing_basis_current_44e_in"`
	AmendedTo44eApplicationIn     string `json:"amended_to_44e_application_in"`
	WithoutBasisCurrentlyIn       string `json:"without_basis_currently_in"`
	FilingCurrentNoBasisIn        string `json:"filing_current_no_basis_in"`
	ColorDrawingFiledIn           string `json:"color_drawing_filed_in"`
	ColorDrawingCurrentIn         string `json:"color_drawing_current_in"`
	Drawing3DFiledIn              string `json:"drawing_3d_filed_in"`
	Drawing3DCurrentIn            string `json:"drawing_3d_current_in"`
	StandardCharactersClaimedIn   string `json:"standard_characters_claimed_in"`
	FilingBasisFiledAs66aIn       string `json:"filing_basis_filed_as_66a_in"`
	FilingBasisCurrent66aIn       string `json:"filing_basis_current_66a_in"`
	LawOfficeAssignedLocationCode string `json:"law_office_assigned_location_code"`
	CurrentLocation               string `json:"current_location"`
	LocationDate                  int64  `json:"location_date"`
	EmployeeName                  string `json:"employee_name"`
}

func (n CaseFileHeader) Value() (driver.Value, error) {
	return common.MarshalJSONHelper(n)
}

func (n *CaseFileHeader) Scan(value interface{}) error {
	return common.UnmarshalJSONHelper(value, n)
}

package abnlookup

import (
	xmldatetime "github.com/datainq/xml-date-time"
)

// ABN holds information about an ABN
type ABN struct {
	IdentiferValue          string                 `xml:"identifierValue"`
	IsCurrentIndicator      string                 `xml:"isCurrentIndicator"`
	ReplacedIdentifierValue string                 `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            xmldatetime.CustomTime `xml:"ReplacedFrom"`
}

// EntityStatus holds information on an Entities current status
type EntityStatus struct {
	StatusCode    string                 `xml:"entityStatusCode"`
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// EntityType holds information on an Entity
type EntityType struct {
	Code        string `xml:"entityTypeCode"`
	Description string `xml:"entityDescription"`
}

// GoodsAndServicesTax holds the effectiveness a GoodsAndServicesTax
type GoodsAndServicesTax struct {
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// DGREndorsement holds information about a DGREndoresment
type DGREndorsement struct {
	EndorsedFrom      xmldatetime.CustomTime `xml:"endorsedFrom"`
	EndorsedTo        xmldatetime.CustomTime `xml:"endorsedTo"`
	EntityEndorsement string                 `xml:"entityEndorement"`
	ItemNumber        string                 `xml:"itemNumber"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string                 `xml:"givenName"`
	OtherGivenName string                 `xml:"otherGivenName,omitempty"`
	FamilyName     string                 `xml:"familyName"`
	EffectiveFrom  xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo    xmldatetime.CustomTime `xml:"effectiveTo"`
}

// MainName holds a main name of a business
type MainName struct {
	OrganisationName string                 `xml:"organisationName"`
	EffectiveFrom    xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo      xmldatetime.CustomTime `xml:"effectiveTo,omitempty"`
}

// MainTradingName holds information about a businesses main trading name
type MainTradingName struct {
	OrganisationName string                 `xml:"organisationName"`
	EffectiveFrom    xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo      xmldatetime.CustomTime `xml:"effectiveTo,omitempty"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string                 `xml:"stateCode"`
	Postcode      uint                   `xml:"postcode"`
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// CharityType holds a information about a charity
type CharityType struct {
	Description   string                 `xml:"charityTypeDescription"`
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// TaxConcessionCharityEndorsement holds information about an endoresment type
type TaxConcessionCharityEndorsement struct {
	EndorsementType string                 `xml:"endorsementType"`
	EffectiveFrom   xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo     xmldatetime.CustomTime `xml:"effectiveTo"`
}

// PBIName holds information about a PBIName
type PBIName struct {
	OrganisationName string `xml:"organisationName"`
	Score            uint   `xml:"score"`
}

// PublicBenevolentInstitutionEmployer holds information about a PBIName and the effectiveness date
type PublicBenevolentInstitutionEmployer struct {
	PBIName       PBIName                `xml:"pbiName"`
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// BusinessName holds information about a businesses name
type BusinessName struct {
	OrganisationName string                 `xml:"organisationName"`
	EffectiveFrom    xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo      xmldatetime.CustomTime `xml:"effectiveTo"`
}

// DGRFundName holds information about a DGRFund's name
type DGRFundName struct {
	OrganisationName   string `xml:"orginisationName"`
	Score              uint   `xml:"score"`
	IsCurrentIndicator string `xml:"isCurrentIndicator"`
}

// DGRFund holds information about a DGRFund
type DGRFund struct {
	DGRFundName  DGRFundName            `xml:"dgrFundName"`
	EndorsedFrom xmldatetime.CustomTime `xml:"endorsedFrom"`
	EndorsedTo   xmldatetime.CustomTime `xml:"endorsedTo"`
	ItemNumber   string                 `xml:"itemNumber"`
}

// ACNCRegistration holds the status and effectiveness of a ACNCRegistration
type ACNCRegistration struct {
	Status        string                 `xml:"status"`
	EffectiveFrom xmldatetime.CustomTime `xml:"effectiveFrom"`
	EffectiveTo   xmldatetime.CustomTime `xml:"effectiveTo"`
}

// BusinessEntity holds all history records and information attached to a search query
type BusinessEntity struct {
	RecordLastUpdatedDate               xmldatetime.CustomTime                 `xml:"recordLastUpdatedDate"`
	ABN                                 []*ABN                                 `xml:"ABN"`
	EntityStatus                        []*EntityStatus                        `xml:"entityStatus"`
	ASICNumber                          string                                 `xml:"ASICNumber,omitempty"`
	EntityType                          EntityType                             `xml:"entityType,omitempty"`
	GoodsAndServicesTax                 []*GoodsAndServicesTax                 `xml:"goodsAndServicesTax,omitempty"`
	DGREndorsement                      []*DGREndorsement                      `xml:"dgrEndorsement,omitempty"`
	LegalName                           []*LegalName                           `xml:"legalName,omitempty"`
	MainName                            []*MainName                            `xml:"mainName,omitempty"`
	MainTradingName                     []*MainTradingName                     `xml:"mainTradingName,omitempty"`
	MainBusinessPhysicalAddress         []*MainBusinessPhysicalAddress         `xml:"mainBusinessPhysicalAddress"`
	CharityType                         CharityType                            `xml:"charityType,omitempty"`
	TaxConcessionCharityEndorsement     []*TaxConcessionCharityEndorsement     `xml:"taxConcessionCharityEndorsement,omitempty"`
	PublicBenevolentInstitutionEmployer []*PublicBenevolentInstitutionEmployer `xml:"publicBenevolentInstitutionEmployer,omitempty"`
	BusinessName                        []*BusinessName                        `xml:"businessName,omitempty"`
	DGRFund                             []*DGRFund                             `xml:"dgrFund,omitempty"`
	ACNCRegistration                    ACNCRegistration                       `xml:"ACNCRegistration,omitempty"`
}

// IdentifierSearchRequest holds information about the data wanting to be queried
type IdentifierSearchRequest struct {
	AuthenticationGUID string `xml:"authenticationGUID"`
	IdentiferType      string `xml:"identifierType"`
	IdentiferValue     string `xml:"identifierValue"`
	History            string `xml:"history"`
}

// Request holds information about the search request
type Request struct {
	IdentifierSearchRequest IdentifierSearchRequest `xml:"identifierSearchRequest"`
}

// Response is what the API has sent back to the client
type Response struct {
	UsageStatement          string                 `xml:"usageStatement"`
	DateRegisterLastUpdated xmldatetime.CustomTime `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       xmldatetime.CustomTime `xml:"dateTimeRetrieved"`
	// Needs to be updated when a new endpoint is released
	BusinessEntity BusinessEntity `xml:"businessEntity201408"`
}

// ABRPayloadSearchResults is the PayLoad the API has sent to the client
type ABRPayloadSearchResults struct {
	Request  Request  `xml:"request"`
	Response Response `xml:"response"`
}

// AuthenticationGUID is the GUID that was used to make the request to the API
func (psr *ABRPayloadSearchResults) AuthenticationGUID() string {
	return psr.Request.IdentifierSearchRequest.AuthenticationGUID
}

// QueryType is the type of query etc. 'ABN', 'ASIC'
func (psr *ABRPayloadSearchResults) QueryType() string {
	return psr.Request.IdentifierSearchRequest.IdentiferType
}

// QueriedString is the string we tried to query etc. '093453453'
func (psr *ABRPayloadSearchResults) QueriedString() string {
	return psr.Request.IdentifierSearchRequest.IdentiferValue
}

// HistoryIncluded checks if the SearchResults include history
func (psr *ABRPayloadSearchResults) HistoryIncluded() bool {
	switch psr.Request.IdentifierSearchRequest.History {
	case "Y":
		return true
	case "y":
		return true
	case "True":
		return true
	case "true":
		return true
	}
	return false
}

// GetACN will get an ACN from the search results if it exists
// Returns an empty string and false if doesn't exist
func (psr *ABRPayloadSearchResults) GetACN() (string, bool) {
	if psr.Response.BusinessEntity.ASICNumber != "" {
		return psr.Response.BusinessEntity.ASICNumber, true
	}
	return "", false
}

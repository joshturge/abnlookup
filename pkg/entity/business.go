package entity

import "encoding/xml"

// Status holds information on an Entities current status
type Status struct {
	Code          string `xml:"entityStatusCode"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

// Type holds information on an Entity
type Type struct {
	Code        string `xml:"entityTypeCode"`
	Description string `xml:"entityDescription"`
}

// GoodsAndServicesTax holds the effectiveness a GoodsAndServicesTax
type GoodsAndServicesTax struct {
	EffectiveFrom Date `xml:"effectiveFrom"`
	EffectiveTo   Date `xml:"effectiveTo"`
}

// DGREndorsement holds information about a DGREndoresment
type DGREndorsement struct {
	EndorsedFrom      Date   `xml:"endorsedFrom"`
	EndorsedTo        Date   `xml:"endorsedTo"`
	EntityEndorsement string `xml:"entityEndorement"`
	ItemNumber        string `xml:"itemNumber"`
}

// MainName holds a main name of a business
type MainName struct {
	Name          string `xml:"organisationName"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// MainTradingName holds information about a businesses main trading name
type MainTradingName struct {
	Name          string `xml:"organisationName"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// Charity holds a information about a charity
type Charity struct {
	Description   string `xml:"charityTypeDescription"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

// TaxConcessionCharityEndorsement holds information about an endoresment type
type TaxConcessionCharityEndorsement struct {
	Type          string `xml:"endorsementType"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

// PBIOrginisation holds information about a PBIName
type PBIOrginisation struct {
	Name  string `xml:"organisationName"`
	Score uint   `xml:"score"`
}

// PublicBenevolentInstitutionEmployer holds information about a PBIName and the effectiveness date
type PublicBenevolentInstitutionEmployer struct {
	Orginisation  PBIOrginisation `xml:"pbiName"`
	EffectiveFrom Date            `xml:"effectiveFrom"`
	EffectiveTo   Date            `xml:"effectiveTo"`
}

// BusinessOrginisation holds information about a businesses name
type BusinessOrginisation struct {
	Name          string `xml:"organisationName"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

// DGRFundOrginisation holds information about a DGRFund's name
type DGRFundOrginisation struct {
	Name               string `xml:"orginisationName"`
	Score              uint   `xml:"score"`
	IsCurrentIndicator Bool   `xml:"isCurrentIndicator"`
}

// DGRFund holds information about a DGRFund
type DGRFund struct {
	Orginisation DGRFundOrginisation `xml:"dgrFundName"`
	EndorsedFrom Date                `xml:"endorsedFrom"`
	EndorsedTo   Date                `xml:"endorsedTo"`
	ItemNumber   string              `xml:"itemNumber"`
}

// ACNCRegistration holds the status and effectiveness of a ACNCRegistration
type ACNCRegistration struct {
	Status        string `xml:"status"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

// Business holds all history records and information attached to a search query
type Business struct {
	LastUpdatedDate                     Date                                   `xml:"recordLastUpdatedDate"`
	ABN                                 []*ABN                                 `xml:"ABN"`
	Status                              []*Status                              `xml:"entityStatus"`
	ACN                                 string                                 `xml:"ASICNumber,omitempty"`
	Type                                Type                                   `xml:"entityType,omitempty"`
	GoodsAndServicesTax                 []*GoodsAndServicesTax                 `xml:"goodsAndServicesTax,omitempty"`
	DGREndorsement                      []*DGREndorsement                      `xml:"dgrEndorsement,omitempty"`
	LegalName                           []*LegalName                           `xml:"legalName,omitempty"`
	MainName                            []*MainName                            `xml:"mainName,omitempty"`
	MainTradingName                     []*MainTradingName                     `xml:"mainTradingName,omitempty"`
	MainBusinessPhysicalAddress         []*MainBusinessPhysicalAddress         `xml:"mainBusinessPhysicalAddress"`
	Charity                             Charity                                `xml:"charityType,omitempty"`
	TaxConcessionCharityEndorsement     []*TaxConcessionCharityEndorsement     `xml:"taxConcessionCharityEndorsement,omitempty"`
	PublicBenevolentInstitutionEmployer []*PublicBenevolentInstitutionEmployer `xml:"publicBenevolentInstitutionEmployer,omitempty"`
	BusinessOrginisation                []*BusinessOrginisation                `xml:"businessName,omitempty"`
	DGRFund                             []*DGRFund                             `xml:"dgrFund,omitempty"`
	ACNCRegistration                    ACNCRegistration                       `xml:"ACNCRegistration,omitempty"`
}

// BusinessResultResponse is what the API has sent back to the client
type BusinessResultResponse struct {
	UsageStatement string `xml:"usageStatement"`
	// Needs to be updated when a new endpoint is released
	Business Business `xml:"businessEntity201408"`
}

// ABRPayloadBusinessResult is the PayLoad the API has sent to the client
type ABRPayloadBusinessResult struct {
	XMLName          xml.Name               `xml:"ABRPayloadSearchResults"`
	BusinessResponse BusinessResultResponse `xml:"response"`
}

// HasUsageStatement will check if a BusinessEntityResponse has a usage statement
func (br *ABRPayloadBusinessResult) HasUsageStatement() bool {
	if br.BusinessResponse.UsageStatement != "" {
		return true
	}
	return false
}

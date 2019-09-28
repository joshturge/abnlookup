package abnlookup

import (
	"encoding/xml"
)

type ABN struct {
	IdentiferValue          string `xml:"identifierValue"`
	IsCurrentIndicator      string `xml:"isCurrentIndicator"`
	ReplacedIdentifierValue string `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            string `xml:"ReplacedFrom"`
}

type EntityStatus struct {
	EntityStatusCode string `xml:"entityStatusCode"`
	EffectiveFrom    string `xml:"effectiveFrom"`
	EffectiveTo      string `xml:"effectiveTo"`
}

type EntityType struct {
	EntityTypeCode    string `xml:"entityTypeCode"`
	EntityDescription string `xml:"entityDescription"`
}

type GoodsAndServicesTax struct {
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

type DGREndorsement struct {
	EndorsedFrom      string `xml:"endorsedFrom"`
	EndorsedTo        string `xml:"endorsedTo"`
	EntityEndorsement string `xml:"entityEndorement"`
	ItemNumber        string `xml:"itemNumber"`
}

type LegalName struct {
	GivenName      string `xml:"givenName"`
	OtherGivenName string `xml:"otherGivenName,omitempty"`
	FamilyName     string `xml:"familyName"`
	EffectiveFrom  string `xml:"effectiveFrom"`
	EffectiveTo    string `xml:"effectiveTo"`
}

type MainName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
}

type MainTradingName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
}

type MainBusinessPhysicalAddress struct {
	StateCode     string `xml:"stateCode"`
	Postcode      uint   `xml:"postcode"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

type CharityType struct {
	CharityTypeDescription string `xml:"charityTypeDescription"`
	EffectiveFrom          string `xml:"effectiveFrom"`
	EffectiveTo            string `xml:"effectiveTo"`
}

type TaxConcessionCharityEndorsement struct {
	EndorsementType string `xml:"endorsementType"`
	EffectiveFrom   string `xml:"effectiveFrom"`
	EffectiveTo     string `xml:"effectiveTo"`
}

type PBIName struct {
	OrganisationName string `xml:"organisationName"`
	Score            uint   `xml:"score"`
}

type PublicBenevolentInstitutionEmployer struct {
	XMLName       xml.Name `xml:"publicBenevolentInstitutionEmployer"`
	PBIName       PBIName  `xml:"pbiName"`
	EffectiveFrom string   `xml:"effectiveFrom"`
	EffectiveTo   string   `xml:"effectiveTo"`
}

type BusinessName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
}

type DGRFundName struct {
	OrganisationName   string `xml:"orginisationName"`
	Score              uint   `xml:"score"`
	IsCurrentIndicator string `xml:"isCurrentIndicator"`
}

type DGRFund struct {
	XMLName      xml.Name    `xml:"dgrFund"`
	DGRFundName  DGRFundName `xml:"dgrFundName"`
	EndorsedFrom string      `xml:"endorsedFrom"`
	EndorsedTo   string      `xml:"endorsedTo"`
	ItemNumber   string      `xml:"itemNumber"`
}

type ACNCRegistration struct {
	Status        string `xml:"status"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

type BusinessEntity struct {
	// Needs to be updated when a new endpoint is released
	XMLName xml.Name `xml:"businessEntity201408"`

	RecordLastUpdatedDate               string                                 `xml:"recordLastUpdatedDate"`
	ABN                                 []ABN                                  `xml:"ABN"`
	EntityStatus                        []EntityStatus                         `xml:"entityStatus"`
	ASICNumber                          string                                 `xml:"ASICNumber,omitempty"`
	EntityType                          *EntityType                            `xml:"entityType,omitempty"`
	GoodsAndServicesTax                 []*GoodsAndServicesTax                 `xml:"goodsAndServicesTax,omitempty"`
	DGREndorsement                      []*DGREndorsement                      `xml:"dgrEndorsement,omitempty"`
	LegalName                           []*LegalName                           `xml:"legalName,omitempty"`
	MainName                            []*MainName                            `xml:"mainName,omitempty"`
	MainTradingName                     []*MainTradingName                     `xml:"mainTradingName,omitempty"`
	MainBusinessPhysicalAddress         []MainBusinessPhysicalAddress          `xml:"mainBusinessPhysicalAddress"`
	CharityType                         *CharityType                           `xml:"charityType,omitempty"`
	TaxConcessionCharityEndorsement     []*TaxConcessionCharityEndorsement     `xml:"taxConcessionCharityEndorsement,omitempty"`
	PublicBenevolentInstitutionEmployer []*PublicBenevolentInstitutionEmployer `xml:"publicBenevolentInstitutionEmployer,omitempty"`
	BusinessName                        []*BusinessName                        `xml:"businessName,omitempty"`
	DGRFund                             []*DGRFund                             `xml:"dgrFund,omitempty"`
	ACNCRegistration                    ACNCRegistration                       `xml:"ACNCRegistration,omitempty"`
}

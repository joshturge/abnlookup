package abnlookup

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Date is a time.Time
type Date struct {
	time.Time
}

// UnmarshalXML will unmarshal a xml date to a time.Time date
func (xdt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02", v)
	if err != nil {
		return fmt.Errorf("couldn't pass time: %s", err.Error())
	}

	*xdt = Date{parse}
	return nil
}

// Bool is a bool
type Bool bool

// UnmarshalXML will unmarshal a "Y" or "N" xml string to a bool evaluation
func (xb *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	*xb = Bool(v == "Y" || v == "N")

	return nil
}

// ABN holds information about an ABN
type ABN struct {
	Value         string `xml:"identifierValue,omitempty"`
	Status        string `xml:"identifierStatus,omitempty"`
	IsCurrent     Bool   `xml:"isCurrentIndicator,omitempty"`
	ReplacedValue string `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom  Date   `xml:"replacedFrom,omitempty"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string `xml:"givenName,omitempty"`
	OtherGivenName string `xml:"otherGivenName,omitempty"`
	FamilyName     string `xml:"familyName,omitempty"`
	EffectiveFrom  Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo    Date   `xml:"effectiveTo,omitempty"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string `xml:"stateCode,omitempty"`
	Postcode      string `xml:"postcode,omitempty"`
	EffectiveFrom Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// Status holds information on an Entities current status
type Status struct {
	Code          string `xml:"entityStatusCode,omitempty"`
	EffectiveFrom Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// Type holds information on an Entity
type Type struct {
	Code        string `xml:"entityTypeCode,omitempty"`
	Description string `xml:"entityDescription,omitempty"`
}

// GoodsAndServicesTax holds the effectiveness a GoodsAndServicesTax
type GoodsAndServicesTax struct {
	EffectiveFrom Date `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date `xml:"effectiveTo,omitempty"`
}

// DGREndorsement holds information about a DGREndoresment
type DGREndorsement struct {
	EndorsedFrom      Date   `xml:"endorsedFrom,omitempty"`
	EndorsedTo        Date   `xml:"endorsedTo,omitempty"`
	EntityEndorsement string `xml:"entityEndorement,omitempty"`
	ItemNumber        string `xml:"itemNumber,omitempty"`
}

// BusinessName holds a name of a business
type BusinessName struct {
	Name          string `xml:"organisationName,omitempty"`
	EffectiveFrom Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// Charity holds a information about a charity
type Charity struct {
	Description   string `xml:"charityTypeDescription,omitempty"`
	EffectiveFrom Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// TaxConcessionCharityEndorsement holds information about an endoresment type
type TaxConcessionCharityEndorsement struct {
	Type          string `xml:"endorsementType,omitempty"`
	EffectiveFrom Date   `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date   `xml:"effectiveTo,omitempty"`
}

// PBIOrginisation holds information about a PBIName
type PBIOrginisation struct {
	Name  string `xml:"organisationName,omitempty"`
	Score uint   `xml:"score,omitempty"`
}

// PublicBenevolentInstitutionEmployer holds information about a PBIName and the effectiveness date
type PublicBenevolentInstitutionEmployer struct {
	Orginisation  *PBIOrginisation `xml:"pbiName,omitempty"`
	EffectiveFrom Date             `xml:"effectiveFrom,omitempty"`
	EffectiveTo   Date             `xml:"effectiveTo,omitempty"`
}

// DGRFundOrginisation holds information about a DGRFund's name
type DGRFundOrginisation struct {
	Name               string `xml:"orginisationName,omitempty"`
	Score              uint   `xml:"score,omitempty"`
	IsCurrentIndicator Bool   `xml:"isCurrentIndicator,omitempty"`
}

// DGRFund holds information about a DGRFund
type DGRFund struct {
	Orginisation *DGRFundOrginisation `xml:"dgrFundName,omitempty"`
	EndorsedFrom Date                 `xml:"endorsedFrom,omitempty"`
	EndorsedTo   Date                 `xml:"endorsedTo,omitempty"`
	ItemNumber   string               `xml:"itemNumber,omitempty"`
}

// Business holds all history records and information attached to a search query
type Business struct {
	LastUpdated                         Date                                   `xml:"recordLastUpdatedDate,omitempty"`
	ABN                                 []*ABN                                 `xml:"ABN,omitempty"`
	Status                              []*Status                              `xml:"entityStatus,omitempty"`
	ACN                                 string                                 `xml:"ASICNumber,omitempty"`
	Type                                *Type                                  `xml:"entityType,omitempty"`
	GoodsAndServicesTax                 []*GoodsAndServicesTax                 `xml:"goodsAndServicesTax,omitempty"`
	DGREndorsement                      []*DGREndorsement                      `xml:"dgrEndorsement,omitempty"`
	LegalName                           []*LegalName                           `xml:"legalName,omitempty"`
	MainName                            []*BusinessName                        `xml:"mainName,omitempty"`
	MainTradingName                     []*BusinessName                        `xml:"mainTradingName,omitempty"`
	MainBusinessPhysicalAddress         []*MainBusinessPhysicalAddress         `xml:"mainBusinessPhysicalAddress,omitempty"`
	Charity                             *Charity                               `xml:"charityType,omitempty"`
	TaxConcessionCharityEndorsement     []*TaxConcessionCharityEndorsement     `xml:"taxConcessionCharityEndorsement,omitempty"`
	PublicBenevolentInstitutionEmployer []*PublicBenevolentInstitutionEmployer `xml:"publicBenevolentInstitutionEmployer,omitempty"`
	BusinessOrginisation                []*BusinessName                        `xml:"businessName,omitempty"`
	DGRFund                             []*DGRFund                             `xml:"dgrFund,omitempty"`
	ACNCRegistration                    *Status                                `xml:"ACNCRegistration,omitempty"`
}

// CurrentABN will return the current ABN a business is using
func (b *Business) CurrentABN() string {
	for _, abn := range b.ABN {
		if abn.IsCurrent && abn.Status != "Cancelled" {
			return abn.Value
		}
	}

	return ""
}

// CurrentStatus will return the most current status
func (b *Business) CurrentStatus() *Status {
	return b.Status[0]
}

// CurrentGST will return the current GST on a business
func (b *Business) CurrentGST() *GoodsAndServicesTax {
	return b.GoodsAndServicesTax[0]
}

// TODO: get current of every struct slice

// Exception describes an exception and provides an exception code.
// More information about exceptions and there meaning can be found
// here: https://api.gov.au/service/5b639f0f63f18432cd0e1a66/Exceptions#exception-codes-and-descriptions
type Exception struct {
	Description string `xml:"exceptionDescription"`
	Code        string `xml:"exceptionCode"`
}

// Error will return a formatted string with information about an API exception
func (e *Exception) Error() error {
	return fmt.Errorf("%s: %s", e.Code, e.Description)

}

// PersonName holds fields on a persons name and search score
type PersonName struct {
	OrganisationName    string `xml:"organisationName,omitempty"`
	FullName            string `xml:"fullName,omitempty"`
	Score               int    `xml:"score,omitempty"`
	IsCurrentIdentifier Bool   `xml:"isCurrentIdentifier,omitempty"`
}

// Person holds information about a person
type Person struct {
	ABN                         ABN                          `xml:"ABN,omitempty"`
	LegalName                   *PersonName                  `xml:"legalName,omitempty"`
	BusinessName                *PersonName                  `xml:"businessName,omitempty"`
	MainName                    *PersonName                  `xml:"mainName,omitempty"`
	MainBusinessPhysicalAddress *MainBusinessPhysicalAddress `xml:"mainBusinessPhysicalAddress,omitempty"`
}

// PeopleResults holds a slice of people
type PeopleResults struct {
	People []*Person `xml:"searchResultsRecord,omitempty"`
}

// ABNList holds a slice of abn's
type ABNList struct {
	ABN []string `xml:"abn"`
}

// Response is an API response
type Response struct {
	UsageStatement string         `xml:"usageStatement,omitempty"`
	Business       *Business      `xml:"businessEntity201408,omitempty"`
	Exception      *Exception     `xml:"exception,omitempty"`
	ABNlist        *ABNList       `xml:"abnList,omitempty"`
	PeopleResults  *PeopleResults `xml:"searchResultsList,omitempty"`
}

// Payload holds the API response
type Payload struct {
	Response *Response `xml:"response,omitempty"`
}

// IsException will check if a payload response is an exception
func (p *Payload) IsException() bool {
	if p.Response.UsageStatement == "" && p.Response.Exception != nil {
		return true
	}
	return false
}

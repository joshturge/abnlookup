package abnlookup

// ABN holds information about an ABN
type ABN struct {
	IdentiferValue          string `xml:"identifierValue"`
	IsCurrentIndicator      string `xml:"isCurrentIndicator"`
	ReplacedIdentifierValue string `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            string `xml:"ReplacedFrom"`
}

// EntityStatus holds information on an Entities current status
type EntityStatus struct {
	StatusCode    string `xml:"entityStatusCode"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

// EntityType holds information on an Entity
type EntityType struct {
	Code        string `xml:"entityTypeCode"`
	Description string `xml:"entityDescription"`
}

// GoodsAndServicesTax holds the effectiveness a GoodsAndServicesTax
type GoodsAndServicesTax struct {
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

// DGREndorsement holds information about a DGREndoresment
type DGREndorsement struct {
	EndorsedFrom      string `xml:"endorsedFrom"`
	EndorsedTo        string `xml:"endorsedTo"`
	EntityEndorsement string `xml:"entityEndorement"`
	ItemNumber        string `xml:"itemNumber"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string `xml:"givenName"`
	OtherGivenName string `xml:"otherGivenName,omitempty"`
	FamilyName     string `xml:"familyName"`
	EffectiveFrom  string `xml:"effectiveFrom"`
	EffectiveTo    string `xml:"effectiveTo"`
}

// MainName holds a main name of a business
type MainName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
	EffectiveTo      string `xml:"effectiveTo,omitempty"`
}

// MainTradingName holds information about a businesses main trading name
type MainTradingName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
	EffectiveTo      string `xml:"effectiveTo,omitempty"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string `xml:"stateCode"`
	Postcode      uint   `xml:"postcode"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

// CharityType holds a information about a charity
type CharityType struct {
	Description   string `xml:"charityTypeDescription"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

// TaxConcessionCharityEndorsement holds information about an endoresment type
type TaxConcessionCharityEndorsement struct {
	EndorsementType string `xml:"endorsementType"`
	EffectiveFrom   string `xml:"effectiveFrom"`
	EffectiveTo     string `xml:"effectiveTo"`
}

// PBIName holds information about a PBIName
type PBIName struct {
	OrganisationName string `xml:"organisationName"`
	Score            uint   `xml:"score"`
}

// PublicBenevolentInstitutionEmployer holds information about a PBIName and the effectiveness date
type PublicBenevolentInstitutionEmployer struct {
	PBIName       PBIName `xml:"pbiName"`
	EffectiveFrom string  `xml:"effectiveFrom"`
	EffectiveTo   string  `xml:"effectiveTo"`
}

// BusinessName holds information about a businesses name
type BusinessName struct {
	OrganisationName string `xml:"organisationName"`
	EffectiveFrom    string `xml:"effectiveFrom"`
	EffectiveTo      string `xml:"effectiveTo"`
}

// DGRFundName holds information about a DGRFund's name
type DGRFundName struct {
	OrganisationName   string `xml:"orginisationName"`
	Score              uint   `xml:"score"`
	IsCurrentIndicator string `xml:"isCurrentIndicator"`
}

// DGRFund holds information about a DGRFund
type DGRFund struct {
	DGRFundName  DGRFundName `xml:"dgrFundName"`
	EndorsedFrom string      `xml:"endorsedFrom"`
	EndorsedTo   string      `xml:"endorsedTo"`
	ItemNumber   string      `xml:"itemNumber"`
}

// ACNCRegistration holds the status and effectiveness of a ACNCRegistration
type ACNCRegistration struct {
	Status        string `xml:"status"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

// BusinessEntity holds all history records and information attached to a search query
type BusinessEntity struct {
	RecordLastUpdatedDate               string                                 `xml:"recordLastUpdatedDate"`
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

// BusinessEntityResponse is what the API has sent back to the client
type BusinessEntityResponse struct {
	UsageStatement          string `xml:"usageStatement"`
	DateRegisterLastUpdated string `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       string `xml:"dateTimeRetrieved"`
	// Needs to be updated when a new endpoint is released
	BusinessEntity BusinessEntity `xml:"businessEntity201408"`
}

// HasUsageStatement will check if a BusinessEntityResponse has a usage statement
func (ber *BusinessEntityResponse) HasUsageStatement() bool {
	if ber.UsageStatement != "" {
		return true
	}
	return false
}

// ABRPayloadBusinessResults is the PayLoad the API has sent to the client
type ABRPayloadBusinessResults struct {
	BusinessEntityResponse BusinessEntityResponse `xml:"response"`
}

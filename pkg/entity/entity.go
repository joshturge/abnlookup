package entity

// Response holds methods to get related data out of
// different response structs
type Response interface {
	HasUsageStatement() bool
}

// ABN holds information about an ABN
type ABN struct {
	Value              string `xml:"identifierValue"`
	IsCurrentIndicator string `xml:"isCurrentIndicator"`
	ReplacedValue      string `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom       string `xml:"ReplacedFrom"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string `xml:"givenName"`
	OtherGivenName string `xml:"otherGivenName,omitempty"`
	FamilyName     string `xml:"familyName"`
	EffectiveFrom  string `xml:"effectiveFrom"`
	EffectiveTo    string `xml:"effectiveTo"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string `xml:"stateCode"`
	Postcode      uint   `xml:"postcode"`
	EffectiveFrom string `xml:"effectiveFrom"`
	EffectiveTo   string `xml:"effectiveTo"`
}

package abnlookup

// PersonEntity holds information about a person
type PersonEntity struct {
	ABN                         ABN                         `xml:"ABN"`
	LegalName                   LegalName                   `xml:"legalName"`
	MainBusinessPhysicalAddress MainBusinessPhysicalAddress `xml:"mainBusinessPhysicalAddress"`
}

// SearchResults holds information on how many records there are
// and if it exceeds the maximum that was set
type SearchResults struct {
	NumberOfRecords int            `xml:"numberOfRecords"`
	ExceedsMaximum  string         `xml:"exceedsMaximum"`
	PersonEntity    []PersonEntity `xml:"searchResultsRecord"`
}

// PersonEntityResponse holds information about an API response
type PersonEntityResponse struct {
	UsageStatement          string        `xml:"usageStatement"`
	DateRegisterLastUpdated string        `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       string        `xml:"dateTimeRetrieved"`
	SearchResults           SearchResults `xml:"searchResultsList"`
}

// HasUsageStatement will check if a PersonEntityResponse has a usage statement
func (per *PersonEntityResponse) HasUsageStatement() bool {
	if per.UsageStatement != "" {
		return true
	}
	return false
}

// ABRPayloadPersonResults holds all the results for a person response
type ABRPayloadPersonResults struct {
	PersonEntityResponse PersonEntityResponse `xml:"response"`
}

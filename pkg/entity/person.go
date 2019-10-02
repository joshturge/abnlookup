package entity

// Person holds information about a person
type Person struct {
	ABN                         ABN                         `xml:"ABN"`
	LegalName                   LegalName                   `xml:"legalName"`
	MainBusinessPhysicalAddress MainBusinessPhysicalAddress `xml:"mainBusinessPhysicalAddress"`
}

// PersonResults holds information on how many records there are
// and if it exceeds the maximum that was set
type PersonResults struct {
	Person []*Person `xml:"searchResultsRecord"`
}

// PersonResultsResponse holds information about an API response
type PersonResultsResponse struct {
	UsageStatement          string        `xml:"usageStatement"`
	DateRegisterLastUpdated string        `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       string        `xml:"dateTimeRetrieved"`
	PersonResults           PersonResults `xml:"searchResultsList"`
}

// ABRPayloadPersonResults holds all the results for a person response
type ABRPayloadPersonResults struct {
	PersonEntityResponse PersonResultsResponse `xml:"response"`
}

// HasUsageStatement will check if a PersonEntityResponse has a usage statement
func (pr *ABRPayloadPersonResults) HasUsageStatement() bool {
	if pr.PersonEntityResponse.UsageStatement != "" {
		return true
	}
	return false
}

package abnlookup

// ABNResults holds a slice of abn's
type ABNResults struct {
	NumberOfRecords int      `xml:"numberOfRecords"`
	ABNList         []string `xml:"abn"`
}

// ABNEntityResponse holds a response with ABN results
type ABNEntityResponse struct {
	UsageStatement            string     `xml:"usageStatement"`
	DateRegisteredLastUpdated string     `xml:"dateRegisteredLastUpdated"`
	DateTimeRetrieved         string     `xml:"dateTimeRetrieved"`
	ABNResults                ABNResults `xml:"abnList"`
}

// HasUsageStatement will check if a ABNEntityResponse has a usage statement
func (abner *ABNEntityResponse) HasUsageStatement() bool {
	if abner.UsageStatement != "" {
		return true
	}
	return false
}

// ABRPayloadABNResults holds a response from the ABN Lookup API
type ABRPayloadABNResults struct {
	ABNEntityResponse ABNEntityResponse `xml:"response"`
}

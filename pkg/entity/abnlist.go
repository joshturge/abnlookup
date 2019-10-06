package entity

// ABNResults holds a slice of abn's
type ABNResults struct {
	ABNList []string `xml:"abn"`
}

// ABNResultsResponse holds a response with ABN results
type ABNResultsResponse struct {
	UsageStatement string     `xml:"usageStatement"`
	ABNResults     ABNResults `xml:"abnList"`
}

// ABRPayloadABNResults holds a response from the ABN Lookup API
type ABRPayloadABNResults struct {
	ABNEntityResponse ABNResultsResponse `xml:"response"`
}

// HasUsageStatement will check if a ABNEntityResponse has a usage statement
func (abnr *ABRPayloadABNResults) HasUsageStatement() bool {
	if abnr.ABNEntityResponse.UsageStatement != "" {
		return true
	}
	return false
}

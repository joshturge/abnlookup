package abnlookup

import (
	"encoding/xml"
	"fmt"
)

type Exception struct {
	Description string `xml:"exceptionDescription"`
	Code        string `xml:"exceptionCode"`
}

type ExceptionResponse struct {
	XMLName        xml.Name `xml:"response"`
	UsageStatement string   `xml:"usageStatement"`
	// Both could cause an error whilst unmarshalling
	DateRegisterLastUpdated string    `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       string    `xml:"dateTimeRetrieved"`
	Exception               Exception `xml:"exception"`
}

type ABRPayloadSearchResultsException struct {
	XMLName           xml.Name          `xml:"ABRPayloadSearchResults"`
	Request           Request           `xml:"request"`
	ExceptionResponse ExceptionResponse `xml:"response"`
}

func (exc Exception) String() string {
	return fmt.Sprintf("exception code: %s description: %s", exc.Code, exc.Description)
}

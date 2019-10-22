package entity

import (
	"fmt"
)

// Exception describes an exception and provides an exception code.
// More information about exceptions and there meaning can be found
// here: https://api.gov.au/service/5b639f0f63f18432cd0e1a66/Exceptions#exception-codes-and-descriptions
type Exception struct {
	Description string `xml:"exceptionDescription"`
	Code        string `xml:"exceptionCode"`
}

// Error will return a formatted string with information about an API exception
func (e *Exception) Error() error {
	return fmt.Errorf("exception code: %s description: %s", e.Code, e.Description)
}

// ExceptionResponse is a response received from the API that has
// an Exception tag
type ExceptionResponse struct {
	UsageStatement string    `xml:"usageStatement"`
	Exception      Exception `xml:"exception"`
}

// ABRPayloadException holds a Request and an ExceptionResponse
// which holds information on a request
type ABRPayloadException struct {
	ExceptionResponse ExceptionResponse `xml:"response"`
}

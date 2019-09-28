package abnlookup

import (
	"encoding/xml"
)

type IdentifierSearchRequest struct {
	AuthenticationGUID string `xml:"authenticationGUID"`
	IdentiferType      string `xml:"identifierType"`
	IdentiferValue     string `xml:"identifierValue"`
	History            string `xml:"history"`
}

type Request struct {
	XMLName                 xml.Name                `xml:"request"`
	IdentifierSearchRequest IdentifierSearchRequest `xml:"identifierSearchRequest"`
}

type Response struct {
	XMLName        xml.Name `xml:"response"`
	UsageStatement string   `xml:"usageStatement"`
	// Both could cause an error whilst unmarshalling
	DateRegisterLastUpdated string `xml:"dateRegisterLastUpdated"`
	DateTimeRetrieved       string `xml:"dateTimeRetrieved"`
	// Needs to be updated when a new endpoint is released
	BusinessEntity BusinessEntity `xml:"businessEntity201408"`
}

type ABRPayloadSearchResults struct {
	XMLName  xml.Name `xml:"ABRPayloadSearchResults"`
	Request  Request  `xml:"request"`
	Response Response `xml:"response"`
}

func (psr ABRPayloadSearchResults) AuthenticationGUID() string {
	return psr.Request.IdentifierSearchRequest.AuthenticationGUID
}

func (psr ABRPayloadSearchResults) QueriedType() string {
	return psr.Request.IdentifierSearchRequest.IdentiferType
}

func (psr ABRPayloadSearchResults) QueriedString() string {
	return psr.Request.IdentifierSearchRequest.IdentiferValue
}

func (psr ABRPayloadSearchResults) History() bool {
	switch psr.Request.IdentifierSearchRequest.History {
	case "Y":
		return true
	case "True":
		return true
	case "true":
		return true
	}
	return false
}

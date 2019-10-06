package entity

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Response holds methods to get related data out of
// different response structs
type Response interface {
	HasUsageStatement() bool
}

type xmlDate struct {
	time.Time
}

func (xdt *xmlDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02", v)
	if err != nil {
		return fmt.Errorf("couldn't pass time: %s", err.Error())
	}

	*xdt = xmlDate{parse}
	return nil
}

// ABN holds information about an ABN
type ABN struct {
	Value              string  `xml:"identifierValue"`
	IsCurrentIndicator string  `xml:"isCurrentIndicator"`
	ReplacedValue      string  `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom       xmlDate `xml:"replacedFrom"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string  `xml:"givenName"`
	OtherGivenName string  `xml:"otherGivenName,omitempty"`
	FamilyName     string  `xml:"familyName"`
	EffectiveFrom  xmlDate `xml:"effectiveFrom"`
	EffectiveTo    xmlDate `xml:"effectiveTo"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string  `xml:"stateCode"`
	Postcode      uint    `xml:"postcode"`
	EffectiveFrom xmlDate `xml:"effectiveFrom"`
	EffectiveTo   xmlDate `xml:"effectiveTo"`
}

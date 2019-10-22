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

// Date is a time.Time
type Date struct {
	time.Time
}

// UnmarshalXML will unmarshal a xml date to a time.Time date
func (xdt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse("2006-01-02", v)
	if err != nil {
		return fmt.Errorf("couldn't pass time: %s", err.Error())
	}

	*xdt = Date{parse}
	return nil
}

// Bool is a bool
type Bool bool

// UnmarshalXML will unmarshal a "Y" or "N" xml string to a bool evaluation
func (xb *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)

	*xb = Bool(v == "Y" || v == "N")

	return nil
}

// ABN holds information about an ABN
type ABN struct {
	Value              string `xml:"identifierValue"`
	IsCurrentIndicator Bool   `xml:"isCurrentIndicator"`
	ReplacedValue      string `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom       Date   `xml:"replacedFrom"`
}

// LegalName holds information about a person's legal name
type LegalName struct {
	GivenName      string `xml:"givenName"`
	OtherGivenName string `xml:"otherGivenName,omitempty"`
	FamilyName     string `xml:"familyName"`
	EffectiveFrom  Date   `xml:"effectiveFrom"`
	EffectiveTo    Date   `xml:"effectiveTo"`
}

// MainBusinessPhysicalAddress holds information for a businesses main physical address
type MainBusinessPhysicalAddress struct {
	StateCode     string `xml:"stateCode"`
	Postcode      string `xml:"postcode"`
	EffectiveFrom Date   `xml:"effectiveFrom"`
	EffectiveTo   Date   `xml:"effectiveTo"`
}

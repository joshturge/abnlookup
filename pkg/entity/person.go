package entity

import "encoding/xml"

// PersonMainName holds fields on a persons main name and search score
type PersonMainName struct {
	Name                string `xml:"organisationName"`
	Score               int    `xml:"score"`
	IsCurrentIdentifier Bool   `xml:"isCurrentIdentifier"`
}

// PersonBusinessOrganisation holds fields on a persons business name and search score
type PersonBusinessOrganisation struct {
	Name                string `xml:"organisationName"`
	Score               int    `xml:"score"`
	IsCurrentIdentifier Bool   `xml:"isCurrentIdentifier"`
}

// PersonLegalName holds fields on a persons name and search score
type PersonLegalName struct {
	FullName            string `xml:"fullName"`
	Score               int    `xml:"score,omitempty"`
	IsCurrentIdentifier Bool   `xml:"isCurrentIdentifier"`
}

// Person holds information about a person
type Person struct {
	ABN                         ABN                         `xml:"ABN"`
	LegalName                   PersonLegalName             `xml:"legalName"`
	BusinessOrginisation        PersonBusinessOrganisation  `xml:"businessName"`
	MainName                    PersonMainName              `xml:"mainName"`
	MainBusinessPhysicalAddress MainBusinessPhysicalAddress `xml:"mainBusinessPhysicalAddress"`
}

// PersonResults holds information on how many records there are
// and if it exceeds the maximum that was set
type PersonResults struct {
	People []*Person `xml:"searchResultsRecord"`
}

// PersonResultsResponse holds information about an API response
type PersonResultsResponse struct {
	UsageStatement string        `xml:"usageStatement"`
	PersonResults  PersonResults `xml:"searchResultsList"`
}

// ABRPayloadPersonResults holds all the results for a person response
type ABRPayloadPersonResults struct {
	XMLName              xml.Name              `xml:"ABRPayloadSearchResults"`
	PersonEntityResponse PersonResultsResponse `xml:"response"`
}

// HasUsageStatement will check if a PersonEntityResponse has a usage statement
func (pr *ABRPayloadPersonResults) HasUsageStatement() bool {
	if pr.PersonEntityResponse.UsageStatement != "" {
		return true
	}
	return false
}

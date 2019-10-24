package abnlookup

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

var (
	stateCodes = []string{
		"NSW",
		"SA",
		"ACT",
		"VIC",
		"WA",
		"NT",
		"QLD",
		"TAS",
	}
)

// SearchByABN will return a ABRPayload with search results for a specified ABN or will return an error
// Also does a validity check on the provided ABN
func (c *Client) SearchByABN(abn string, history bool) (*Business, error) {
	if ValidateABN(abn) {
		return c.searchByNumber("SearchByABNv201408", abn, history)
	}
	return nil, fmt.Errorf("ABN provided is not valid")
}

// SearchByASIC will return a ABRPayload with search results for a specified ASIC or return an error
func (c *Client) SearchByASIC(asic string, history bool) (*Business, error) {
	if ValidateACN(asic) {
		return c.searchByNumber("SearchByASICv201408", asic, history)
	}

	return nil, fmt.Errorf("ASIC provided is not valid")
}

// searchByNumber will make a request to the ABN Lookup API and attempt to decode the response body
// into a Payload struct
func (c *Client) searchByNumber(searchType string, query string, history bool) (*Business, error) {
	// Add url values
	// NOTE: the authentication GUID is added in NewRequest for you
	v := url.Values{}
	v.Add("searchString", query)
	v.Add("includeHistoricalDetails", returnYorNString(history))

	req, err := c.NewRequest(searchType, v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	// Do the request and decode the response body into a payload struct
	var payload Payload
	resp, err := c.Do(req, &payload)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}
	defer resp.Body.Close()

	if payload.IsException() {
		return nil, payload.Response.Exception.Error()
	}

	return payload.Response.Business, nil
}

// NameQuery is a query that is used to search for an ABN by a persons name
type NameQuery struct {
	StateCodes     []string
	Postcode       string
	LegalName      bool
	TradingName    bool
	BusinessName   bool
	ActiveABNsOnly bool
	SearchWidth    string
	MinimumScore   int
	MaxResults     int
}

// SearchByName allows you to lookup an ABN/s via a name
func (c *Client) SearchByName(name string, nq *NameQuery) ([]*Person, error) {
	v := url.Values{}
	v.Add("name", name)

	if len(nq.StateCodes) == 0 {
		for _, stateCode := range stateCodes {
			v.Add(stateCode, "Y")
		}
	}

	for _, stateCode := range stateCodes {
		if elementExists(nq.StateCodes, stateCode) {
			v.Add(stateCode, "Y")
			continue
		}
		v.Add(stateCode, "N")
	}

	v.Add("postcode", nq.Postcode)
	v.Add("legalName", returnYorNString(nq.LegalName))
	v.Add("tradingName", returnYorNString(nq.TradingName))
	v.Add("businessName", returnYorNString(nq.BusinessName))
	v.Add("activeABNsOnly", returnYorNString(nq.ActiveABNsOnly))
	v.Add("searchWidth", nq.SearchWidth)
	v.Add("minimumScore", strconv.Itoa(nq.MinimumScore))
	v.Add("maxSearchResults", strconv.Itoa(nq.MaxResults))

	req, err := c.NewRequest("ABRSearchByNameAdvancedSimpleProtocol2017", v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	// Do the request and decode the response body into an ABRPayloadBusinessResults struct
	var payload Payload
	resp, err := c.Do(req, &payload)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}
	defer resp.Body.Close()

	if payload.IsException() {
		return nil, payload.Response.Exception.Error()
	}

	return payload.Response.PeopleResults.People, nil
}

// ABNStatusQuery holds fields for a ABNStatus query
type ABNStatusQuery struct {
	Postcode                   string
	ActiveABNsOnly             bool
	CurrentGSTRegistrationOnly bool
	EntityTypeCode             string
}

// SearchByABNStatus allows you to search for ABN/s via ABN Status
func (c *Client) SearchByABNStatus(asq *ABNStatusQuery) ([]string, error) {
	v := url.Values{}
	v.Add("postcode", asq.Postcode)
	v.Add("activeABNsOnly", returnYorNString(asq.ActiveABNsOnly))
	v.Add("currentGSTRegistrationOnly", returnYorNString(asq.CurrentGSTRegistrationOnly))
	v.Add("entityTypeCode", asq.EntityTypeCode)

	ABNList, err := c.filterSearch("SearchByABNStatus", v)
	if err != nil {
		return nil, err
	}

	return ABNList, nil
}

// CharityQuery holds fields for a charity query
type CharityQuery struct {
	Postcode           string
	StateCode          string
	CharityTypeCode    string
	ConcessionTypeCode string
}

// SearchByCharity allows you to search for ABN/s via a charity
func (c *Client) SearchByCharity(cq *CharityQuery) ([]string, error) {
	v := url.Values{}
	v.Add("postcode", cq.Postcode)
	v.Add("state", cq.StateCode)
	v.Add("charityTypeCode", cq.CharityTypeCode)
	v.Add("concessionTypeCode", cq.ConcessionTypeCode)

	ABNList, err := c.filterSearch("SearchByCharity", v)
	if err != nil {
		return nil, err
	}

	return ABNList, nil
}

// SearchByPostcode allows you to search for ABN/s via a specified postcode
func (c *Client) SearchByPostcode(postcode string) ([]string, error) {
	v := url.Values{}
	v.Add("postcode", postcode)

	ABNList, err := c.filterSearch("SearchByCharity", v)
	if err != nil {
		return nil, err
	}

	return ABNList, nil
}

// RegistrationEventQuery holds fields for a registration event query
type RegistrationEventQuery struct {
	Postcode       string
	EntityTypeCode string
	StateCode      string
	Date           time.Time
}

// SearchByRegistrationEvent allows you to search for ABN/s via a registration event query
func (c *Client) SearchByRegistrationEvent(req *RegistrationEventQuery) ([]string, error) {
	v := url.Values{}
	v.Add("postcode", req.Postcode)
	v.Add("entityTypeCode", req.EntityTypeCode)
	v.Add("state", req.StateCode)
	v.Add("month", strconv.Itoa(int(req.Date.Month())))
	v.Add("year", strconv.Itoa(req.Date.Year()))

	ABNList, err := c.filterSearch("SearchByRegistrationEvent", v)
	if err != nil {
		return nil, err
	}

	return ABNList, nil
}

// UpdateEventQuery holds fields for a update event query
type UpdateEventQuery struct {
	Postcode       string
	EntityTypeCode string
	StateCode      string
	UpdateDate     time.Time
}

// SearchByUpdateEvent allows you to search for ABN/s via a update event query
func (c *Client) SearchByUpdateEvent(ueq *UpdateEventQuery) ([]string, error) {
	v := url.Values{}
	v.Add("postcode", ueq.Postcode)
	v.Add("entityTypeCode", ueq.EntityTypeCode)
	v.Add("state", ueq.StateCode)
	v.Add("updateDate", ueq.UpdateDate.Format("2006-01-02T15:04:05"))

	ABNList, err := c.filterSearch("SearchByUpdateEvent", v)
	if err != nil {
		return nil, err
	}

	return ABNList, nil
}

func (c *Client) filterSearch(path string, v url.Values) ([]string, error) {
	req, err := c.NewRequest(path, v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	var payload Payload
	resp, err := c.Do(req, &payload)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}
	defer resp.Body.Close()

	if payload.IsException() {
		return nil, payload.Response.Exception.Error()
	}

	return payload.Response.ABNlist.ABN, nil
}

func returnYorNString(b bool) string {
	if b {
		return "Y"
	}

	return "N"
}

func elementExists(slice []string, element string) bool {
	s := reflect.ValueOf(slice)

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == element {
			return true
		}
	}

	return false
}

package abnlookup

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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
func (c *Client) SearchByABN(abn string, history bool) (*ABRPayloadBusinessResults, error) {
	if ValidateABN(abn) {
		return c.searchBy("SearchByABNv201408", abn, history)
	}
	return nil, fmt.Errorf("ABN provided is not valid")
}

// SearchByASIC will return a ABRPayload with search results for a specified ASIC or return an error
func (c *Client) SearchByASIC(asic string, history bool) (*ABRPayloadBusinessResults, error) {
	if ValidateACN(asic) {
		return c.searchBy("SearchByASICv201408", asic, history)
	}

	return nil, fmt.Errorf("ASIC provided is not valid")
}

// searchBy will make a request to the ABN Lookup API and attempt to decode the response body
// into a ABRPayloadSearchResults struct, if the Response UsageStatement is not set then this
// function will then try to decode the response body into a ABRPayloadException struct and
// return an error.
func (c *Client) searchBy(searchType string, query string, history bool) (*ABRPayloadBusinessResults, error) {
	// Add url values
	// NOTE: the authentication GUID is added in NewRequest for you
	v := url.Values{}
	v.Add("searchString", query)
	v.Add("includeHistoricalDetails", returnYorNString(history))

	req, err := c.NewRequest(searchType, v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	// Do the request and decode the response body into an ABRPayloadBusinessResults struct
	var ABRPBR ABRPayloadBusinessResults
	resp, err := c.Do(req, &ABRPBR)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}

	if err = checkExceptionResponse(resp, &ABRPBR.BusinessEntityResponse); err != nil {
		return nil, err
	}

	return &ABRPBR, nil
}

// NameQuery is a query that is used to search for an ABN by a persons name
type NameQuery struct {
	Name           string
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

func (c *Client) SearchByName(nq NameQuery) (*ABRPayloadPersonResults, error) {
	v := url.Values{}
	v.Add("name", nq.Name)
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
	var ABRPPR ABRPayloadPersonResults
	resp, err := c.Do(req, &ABRPPR)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}

	if err = checkExceptionResponse(resp, &ABRPPR.PersonEntityResponse); err != nil {
		return nil, err
	}

	return &ABRPPR, nil
}

func checkExceptionResponse(resp *http.Response, entityResp EntityResponse) error {
	// If the usage statement isn't defined then there was probably an exception
	if !entityResp.HasUsageStatement() {
		var ABRPException ABRPayloadException
		if err := xml.NewDecoder(resp.Body).Decode(&ABRPException); err != nil {
			return fmt.Errorf("couldn't decode response body into ABRPayloadException: %s", err)
		}

		return fmt.Errorf(ABRPException.ExceptionResponse.Exception.Error())
	}

	return nil
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

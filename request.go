package abnlookup

import (
	"encoding/xml"
	"fmt"
	"net/url"
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
	// The Lookup API requires a string for includeHistoricalDetails that can either be 'Y' or 'N'
	// true == "Y", false == "N".
	var includeHistory string
	if history {
		includeHistory = "Y"
	} else {
		includeHistory = "N"
	}

	// Add url values
	// NOTE: the authentication GUID is added in NewRequest for you
	v := url.Values{}
	v.Add("searchString", query)
	v.Add("includeHistoricalDetails", includeHistory)

	req, err := c.NewRequest(searchType, v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	// Do the request and decode the response body into an ABRPayloadSearchResults struct
	var ABRPBR ABRPayloadBusinessResults
	resp, err := c.Do(req, &ABRPBR)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}

	// If the usage statement isn't defined then there was probably an exception
	if ABRPBR.Response.UsageStatement == "" {
		var ABRPException ABRPayloadException
		if err = xml.NewDecoder(resp.Body).Decode(&ABRPException); err != nil {
			return nil, fmt.Errorf("couldn't decode response body into ABRPayloadException: %s", err)
		}

		return nil, fmt.Errorf(ABRPException.Error())
	}

	return &ABRPBR, nil
}

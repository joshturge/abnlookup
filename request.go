package abnlookup

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

func (c *Client) SearchByABN(abn string, history bool) (*ABRPayloadSearchResults, error) {
	var includeHistory string
	if history {
		includeHistory = "Y"
	} else {
		includeHistory = "N"
	}

	v := url.Values{}
	v.Add("searchString", abn)
	v.Add("includeHistoricalDetails", includeHistory)

	req, err := c.NewRequest("SearchByABNv201408", v)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err)
	}

	var ABRPSR ABRPayloadSearchResults
	resp, err := c.Do(req, &ABRPSR)
	if err != nil {
		return nil, err
	}

	// If the usage statement isn't defined then there was probably an exception
	if ABRPSR.Response.UsageStatement == "" {
		var ABRPException ABRPayloadException
		if err = xml.NewDecoder(resp.Body).Decode(&ABRPException); err != nil {
			return nil, fmt.Errorf("couldn't decode response body into ResponseException: %s", err)
		}

		return nil, fmt.Errorf(ABRPException.ExceptionResponse.Exception.String())
	}

	return &ABRPSR, nil
}

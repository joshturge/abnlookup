package abnlookup

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Client holds information about a http client and GUID
// which is used for API authentication
type Client struct {
	BaseURL    *url.URL
	GUID       string
	httpClient http.Client
	logger     *log.Logger
}

// NewClient will create a instance of client with GUID you provide.
// If you don't have a GUID yet, you can register for one here:
// https://api.gov.au/service/5b639f0f63f18432cd0e1a66/Registration
func NewClient(guid string) (*Client, error) {
	base, err := url.Parse("https://abr.business.gov.au/abrxmlsearch/AbrXmlSearch.asmx/")
	if err != nil {
		return nil, fmt.Errorf("couldn't pass the base url: %s", err.Error())
	}

	logger := log.New(os.Stdout, "[ABN-LOOKUP] ", 0)

	client := &Client{
		BaseURL: base,
		GUID:    guid,
		logger:  logger,
	}

	return client, nil
}

// SetTimeout will set a timeout for requests
func (c *Client) SetTimeout(duration time.Duration) {
	c.httpClient.Timeout = duration
}

// NewRequest creates a new request instance, your GUID will be added
// to your url.Values and the host header is set within this function.
func (c *Client) NewRequest(path string, urlVal url.Values) (*http.Request, error) {

	urlVal.Add("authenticationGuid", c.GUID)

	// Generate the url
	var url bytes.Buffer
	url.WriteString(c.BaseURL.String())
	url.WriteString(path)
	url.WriteString("?")
	url.WriteString(urlVal.Encode())

	// Create a new GET request, body is nil as the values needed for an API request are encoded within the url
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Host", "abr.business.gov.au")

	return req, nil
}

// Do will execute a http request and decode a response body into a valid struct
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// For debugging purposes
	c.logger.Println(req.Method, req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, fmt.Errorf("couldn't read from response body: %s", err.Error())
		}
		return resp, fmt.Errorf("API response status was not 200: Got %d: %s: Response.Body: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
	}

	// This is definently not ideal for big response bodies but the API
	// responses I've seen so far are only 1-2KB max in size.
	// A better option would be to use io.Copy() but I'm not too sure
	// how I would implement it.
	// Get a copy of the response body as I need to have resp.Body available
	var bodyByte []byte
	bodyByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read from response body: %s", err.Error())
	}
	// Set resp.Body to a new ReadCloser with same response body
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))

	// Create a Reader from bodyByte so it can be decoded
	body := bytes.NewReader(bodyByte)

	// Decode response body into struct
	if err = xml.NewDecoder(body).Decode(v); err != nil {
		return nil, fmt.Errorf("couldn't decode response body into struct: %s", err.Error())
	}

	return resp, nil
}

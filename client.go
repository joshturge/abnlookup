package abnlookup

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Client holds information about a http client
type Client struct {
	BaseURL    *url.URL
	GUID       string
	httpClient http.Client
	logger     *log.Logger
}

// NewClient will create a new instance of client with GUID provided
func NewClient(guid string) (*Client, error) {
	base, err := url.Parse("http://abr.business.gov.au/ABRXMLSearch/")
	if err != nil {
		return nil, err
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

// NewRequest creates a new request instance
func (c *Client) NewRequest(path string, form url.Values) (*http.Request, error) {
	ref := &url.URL{Path: path}
	uri := c.BaseURL.ResolveReference(ref)

	req, err := http.NewRequest("POST", uri.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.PostForm = form
	req.Header.Set("Host", "abr.business.gov.au")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// Do will execute a http request and decode the response body into a struct
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	c.logger.Println(req.Method, req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Printf("http.Client.Do: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: Find a way to check for errors, status code is always set to 200 for some reason

	if err := xml.NewDecoder(resp.Body).Decode(v); err != nil {
		c.logger.Printf("xml.Decoder.Decode: %s", err)
		return nil, err
	}

	return resp, nil
}

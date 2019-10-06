package abnlookup

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	// LogDiscard will log to os.Discard
	LogDiscard = 0
	// LogDebug will log everything
	LogDebug = 1
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
func NewClient(guid string, logLevel int) (*Client, error) {
	base, err := url.Parse("https://abr.business.gov.au/abrxmlsearch/AbrXmlSearch.asmx/")
	if err != nil {
		return nil, fmt.Errorf("couldn't pass the base url: %s", err.Error())
	}

	var logger *log.Logger
	switch logLevel {
	case LogDiscard:
		logger = log.New(ioutil.Discard, "", 0)
	case LogDebug:
		logger = log.New(os.Stdout, "[DEBUG] ", 1)
	}

	client := &Client{
		BaseURL: base,
		GUID:    guid,
		logger:  logger,
	}

	return client, nil
}

// log a message if debug logging is enabled for the client
func (c *Client) log(format string, v ...interface{}) {
	if c.logger.Flags() == LogDebug {
		c.logger.Printf(format, v...)
	}
}

// SetTimeout will set a timeout for requests
func (c *Client) SetTimeout(duration time.Duration) {
	c.log("client timeout set to %s\n", duration.String())
	c.httpClient.Timeout = duration
}

// NewRequest creates a new request instance, your GUID will be added
// to your url.Values and the host header is set within this function.
func (c *Client) NewRequest(path string, urlVal url.Values) (*http.Request, error) {

	urlVal.Add("authenticationGuid", c.GUID)

	// Generate the url
	var url bytes.Buffer
	fmt.Fprintf(&url, "%s%s?%s", c.BaseURL.String(), path, urlVal.Encode())

	c.log("creating new request with url values: %s\n", urlVal.Encode())

	// Create a new GET request, body is nil as the values needed for an API request are encoded within the url
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't create new request: %s", err.Error())
	}

	req.Header.Set("Host", "abr.business.gov.au")

	return req, nil
}

func drainBody(b io.ReadCloser) (*bytes.Buffer, error) {
	var body bytes.Buffer
	if _, err := io.Copy(&body, b); err != nil {
		return nil, fmt.Errorf("couldn't drain response body: %s", err.Error())
	}
	return &body, nil
}

// Do will execute a http request and decode a response body into a valid struct
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	c.log("%s %s\n", req.Method, req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't do request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("API response status was not 200: Got %d: %s, check resp.Body for more info", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	c.log("copying response body")

	// Get a copy of the response body
	var body *bytes.Buffer
	body, err = drainBody(resp.Body)
	if err != nil {
		return resp, err
	}

	// Create a new bytes reader and add a closer
	resp.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))

	// Decode response body into struct
	if err = xml.NewDecoder(body).Decode(v); err != nil {
		return resp, fmt.Errorf("couldn't decode response body into struct: %s", err.Error())
	}

	return resp, nil
}

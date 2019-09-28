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

// Client holds information about a http client
type Client struct {
	BaseURL    *url.URL
	GUID       string
	httpClient http.Client
	logger     *log.Logger
}

// NewClient will create a new instance of client with GUID provided
func NewClient(guid string) (*Client, error) {
	base, err := url.Parse("https://abr.business.gov.au/abrxmlsearch/AbrXmlSearch.asmx/")
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
	form.Add("authenticationGuid", c.GUID)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?%s", c.BaseURL.String(), path, form.Encode()), nil)
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
		return nil, err
	}

	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("client response status was not 200: Got %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// This is definently not ideal for big response bodies but it's all I know at this point
	// Get a copy of the response body as I need to have resp.Body available
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	body := bytes.NewReader(bodyByte)

	if err = xml.NewDecoder(body).Decode(v); err != nil {
		return nil, err
	}

	return resp, nil
}

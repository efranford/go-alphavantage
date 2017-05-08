package alphavantage

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

const (
	defaultBaseURL = "http://www.alphavantage.co"
)

type Client struct {
	client	*http.Client

	BaseURL	*url.URL
	ApiKey string
	TimeSeries *TimeSeriesService

	common service
}

type service struct {
	client *Client
}

func NewClient(apiKey string) *Client {

	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, ApiKey: apiKey}
	c.common.client = c

	c.TimeSeries = (*TimeSeriesService)(&c.common)

	return c
}

func (c *Client) NewGetRequest(urlStr string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	q := req.URL.Query()
	q.Add("apikey", c.ApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err :=c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
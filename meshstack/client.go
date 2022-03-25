package meshstack

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var HttpClient *http.Client

type Client struct {
	apiVersion string
	BaseUrl    *url.URL
	httpClient *http.Client
	headers    http.Header
}

func NewClient(h http.Header, port int, hostname, version string) *Client {
	client := &Client{apiVersion: version}

	client.httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	urlstr := fmt.Sprintf("%s:%d", hostname, port)
	if u, err := url.Parse(urlstr); err != nil {
		panic("Could not init Provider client to meshStack API")
	} else {
		client.BaseUrl = u
	}

	if h != nil {
		client.headers = h
	}

	return client
}

// HTTP GET - /details?name=${name}
func (c *Client) executeGetAPI(baseURL string, apiUri string, resourceName string, resourceHeaders http.Header) (b []byte, err error) {
	path := fmt.Sprintf("%s/%s/%s", baseURL, apiUri, resourceName)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	for name, value := range c.headers {
		req.Header.Add(name, strings.Join(value, ""))
	}

	for name, value := range resourceHeaders {
		req.Header.Add(name, strings.Join(value, ""))
	}

	log.Printf("Calling %s\n", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to issue get details API call: %s", err.Error())
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response from name allocation API: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Name allocation API returned HTTP status code %d", resp.StatusCode)
	}

	log.Printf("Raw output: %v\n", string(b))
	return
}

// HTTP PUT - /name?resource_type=${type}&region=${region}
func (c *Client) doAllocateName(baseURL, resourceType, region string) (b []byte, err error) {
	path := fmt.Sprintf("http://%s/name", baseURL)
	req, err := http.NewRequest(http.MethodPut, path, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = fmt.Sprintf("resource_type=%s&region=%s", resourceType, region)
	for name, value := range c.headers {
		req.Header.Add(name, strings.Join(value, ""))
	}

	log.Printf("Calling %s\n", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to issue name allocation API call: %s", err.Error())
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type of HTTP response is invalid")
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response from name allocation API: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Name allocation API returned HTTP status code %d", resp.StatusCode)
	}

	log.Printf("Raw output: %v\n", string(b))
	return
}

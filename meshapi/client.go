package meshapi

import (
	"bytes"
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
	BaseUrl    *url.URL
	httpClient *http.Client
	headers    http.Header
}

func NewClient(hostname string, port int, h http.Header) *Client {
	client := &Client{}

	client.httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	urlstr := fmt.Sprintf("%s:%d", hostname, port)
	if u, err := url.Parse(urlstr); err != nil {
		panic("Could not init Provider client to meshapi API")
	} else {
		client.BaseUrl = u
	}

	if h != nil {
		client.headers = h
	}

	return client
}

func (c *Client) executeGetAPI(baseURL string, apiUri string, resourceName string, resourceHeaders http.Header, resourceQueries map[string]string) (b []byte, err error) {
	var path string
	if resourceName != "" {
		path = fmt.Sprintf("%s/%s/%s", baseURL, apiUri, resourceName)
	} else {
		path = fmt.Sprintf("%s/%s", baseURL, apiUri)
	}
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range resourceQueries {
		req.URL.Query().Add(key, value)
	}

	for key, value := range c.headers {
		req.Header.Add(key, strings.Join(value, ""))
	}

	for key, value := range resourceHeaders {
		req.Header.Add(key, strings.Join(value, ""))
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

func (c *Client) executePutAPI(baseURL string, jsonBody string, resourceHeaders http.Header) (b []byte, err error) {
	path := fmt.Sprintf("%s/api/meshobjects", baseURL)
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		return nil, err
	}

	for key, value := range c.headers {
		req.Header.Add(key, strings.Join(value, ""))
	}

	for key, value := range resourceHeaders {
		req.Header.Add(key, strings.Join(value, ""))
	}

	log.Printf("Calling %s\n", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to meshObject Declarative Import API call: %s", err.Error())
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response from Declarative Import API: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Declarative Import API returned HTTP status code %d", resp.StatusCode)
	}

	log.Printf("Raw output: %v\n", string(b))
	return
}

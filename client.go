package gojikan

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientHelper interface {
	Get(endpoint string, params url.Values) ([]byte, error)
}

type clientHelper struct {
	httpClient HttpClient
}

func NewClientHelper(httpClient HttpClient) *clientHelper {
	return &clientHelper{
		httpClient: httpClient,
	}
}

func (api *clientHelper) Get(endpoint string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = params.Encode()
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

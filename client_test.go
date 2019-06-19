package gojikan

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

type MockHttpClient struct{}

func NewMockHttpClient() HttpClient {
	return &MockHttpClient{}
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	body := []byte(req.URL.RawQuery)
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}, nil
}

func TestClientHelper_Get(t *testing.T) {
	clientHelper := NewClientHelper(NewMockHttpClient())

	q := url.Values{}
	q.Add("q", "Guk")
	q.Add("page", strconv.Itoa(1))
	q.Add("limit", strconv.Itoa(10))

	bytesArr, err := clientHelper.Get("abc", q)
	if err != nil {
		t.Errorf("Return error: %v", err.Error())
	}

	exp := "limit=10&page=1&q=Guk"
	resp := string(bytesArr[:])
	if resp != exp {
		t.Errorf("Response mismatch. Expected: %v, actual %v", exp, resp)
	}
}

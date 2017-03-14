package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCheck(t *testing.T) {
	expectedResponse := "Snappy Test Device Imager"
	s := httptest.NewServer(http.HandlerFunc(Check))
	defer s.Close()
	resp, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != expectedResponse {
		t.Error("Expected", expectedResponse, "got", string(body))
	}
}

func TestRuncmd(t *testing.T) {
    expectedResponse := "testworked\n"
    s := httptest.NewServer(http.HandlerFunc(Runcmd))
    defer s.Close()
    req, err := http.NewRequest("GET", s.URL, nil)
    q := req.URL.Query()
    q.Add("cmd", "echo testworked")
    req.URL.RawQuery = q.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != expectedResponse {
		t.Error("Expected", expectedResponse, "got", string(body))
	}
}

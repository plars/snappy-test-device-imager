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

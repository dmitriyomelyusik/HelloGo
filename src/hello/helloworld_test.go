package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBodyData(t *testing.T) {
	//ts := httptest.NewServer(http.HandlerFunc(HandleHelloWorld))
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	got, _ := ioutil.ReadAll(res.Body)
	expected := "Hello, Go!"

	if string(got) != expected {
		t.Errorf("unexpected body: got %v expected %v", string(got), expected)
	}
}

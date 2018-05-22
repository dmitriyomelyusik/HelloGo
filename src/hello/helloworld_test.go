package main

import (
  //"fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  //"io/ioutil"
)

func TestBodyData(t *testing.T) {
  req, err := http.NewRequest("GET", "http://localhost:8080", nil)
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(HandleHelloWorld)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  expected := "Hello, Go!"
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

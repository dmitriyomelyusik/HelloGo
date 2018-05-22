package main

import (
  "fmt"
  "net/http"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Go!")
}

func main() {
  http.HandleFunc("/", HandleHelloWorld)
  http.ListenAndServe(":8080", nil)
}

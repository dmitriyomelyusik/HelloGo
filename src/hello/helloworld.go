package main

import (
  "fmt"
  "net/http"
  "log"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Go!")
}

func main() {
  http.HandleFunc("/", HandleHelloWorld)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

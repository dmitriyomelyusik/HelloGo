package main

import (
	"fmt"
	"log"
	"net/http"
)

/*//HandleHelloWorld func writes string "Hello, Go!" into the ResponseWriter
func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}*/

func main() {
	//http.HandleFunc("/", HandleHelloWorld)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

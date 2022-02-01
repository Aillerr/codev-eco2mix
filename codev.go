package main

import (
	E2M "codev/eco2mix"
	"fmt"
	"log"
	"net/http"
)

var apiKey string = "751d9da6351bcb1d0a3710004096de3c5b0cb94d36e3264cb6d1d5f4"

func main() {
	httpServer()
}

func callE2M(apiKey string) {
	E2M.MakeRequest(apiKey)
}

func httpServer() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler type function for a http call
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//Special case, depends on what the server gets
	callE2M(apiKey)
}

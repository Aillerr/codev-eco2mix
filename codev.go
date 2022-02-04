package main

import (
	E2M "codev/eco2mix"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var apiKey string = "751d9da6351bcb1d0a3710004096de3c5b0cb94d36e3264cb6d1d5f4"
var port int = 8000
var portStr = ":" + strconv.Itoa(port)

func main() {
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "")
	os.Setenv("NET", "tcp")
	os.Setenv("ADDR", "127.0.0.1")
	callE2M(apiKey)
	//httpServer()
}

func callE2M(apiKey string) {
	E2M.Maineco2mix(apiKey)
}

func httpServer() {
	//List all handlers
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(portStr, nil))
}

// Handler type function for a http call
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "E2M")
	//Special case, depends on what the server gets
	callE2M(apiKey)
}

// Simple program which prints the hostname of your webserver
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	listen = "0.0.0.0"
	port   = "8888"
)

var returnCodeError bool = false

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	if returnCodeError {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "The hostname is: "+hostname+"\n")
}

func trigger(w http.ResponseWriter, r *http.Request) {
	returnCodeError = true
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/trigger", trigger)
	log.Fatal(http.ListenAndServe(listen+":"+port, nil))
}

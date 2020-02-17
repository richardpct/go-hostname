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

var count int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	if count >= 5 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error\n")
	} else {
		fmt.Fprintf(w, "Version: v3\nThe hostname is: "+hostname+"\n")
	}

	count++
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen+":"+port, nil))
}

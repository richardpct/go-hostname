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

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "The hostname is: "+hostname+"\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen+":"+port, nil))
}

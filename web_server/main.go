package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, "Error occurred", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Some error occurred"))
			return
		}
		log.Printf("Request received, data: %s", payload)
		fmt.Fprintf(rw, "Hello %s\n", payload) // returning response, writing to ResponseWriter
	})

	http.HandleFunc("/hello", func(http.ResponseWriter, *http.Request) {
		log.Println("/hello Request received...")
	})

	http.ListenAndServe(":8080", nil)
}

// Client request: curl -d "some data" localhost:8080

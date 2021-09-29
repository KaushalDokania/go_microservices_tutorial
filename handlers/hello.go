package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.logger.Println("Hello World!!")

	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Error occurred", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", payload) // returning response, writing to ResponseWriter
}

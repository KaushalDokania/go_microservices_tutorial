package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KaushalDokania/go_microservices_tutorial/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	products := data.GetProducts()
	d, err := json.Marshal(products)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}

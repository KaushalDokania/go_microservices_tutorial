package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KaushalDokania/go_microservices_tutorial/handlers"
)

func main() {
	logger := log.New(os.Stdout, "microservice-api", log.LstdFlags)
	helloHander := handlers.NewHello(logger)
	goodByeHandler := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", helloHander)
	sm.Handle("/bye", goodByeHandler)

	http.ListenAndServe(":8080", sm)
}

// Client request: curl -d "some data" localhost:8080

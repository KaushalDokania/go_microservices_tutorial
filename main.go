package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KaushalDokania/go_microservices_tutorial/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "microservice-api", log.LstdFlags)
	productHandler := handlers.NewProducts(logger)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProducts)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProducts)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port 8080...")
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received Terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	// http.ListenAndServe(":8080", sm)
}

// Client request: curl -d "some data" localhost:8080

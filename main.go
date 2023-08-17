package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	address := ":8080"

	mux := http.NewServeMux() //handler

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Starting server at port: %v", address)

	log.Fatal(s.ListenAndServe())
}

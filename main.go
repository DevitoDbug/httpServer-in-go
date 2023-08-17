package main

import (
	"log"
	"net/http"
	"simpleWebServer/pkg/server"
	"time"
)

func main() {

	address := ":8080"

	mux := http.NewServeMux() //handler

	myServer := server.New()
	mux.HandleFunc("/", myServer.HandleIndex)

	//Adding the user route
	mux.HandleFunc("/users", myServer.HandleUser)

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

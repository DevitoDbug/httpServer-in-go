package main

import (
	"log"
	"net/http"
	"time"
)

var indexPage = `
<!DOCTYPE html>
<html>
<body style = "width: 500px ; margin: 0 auto;">
<h1>Our first server</h2>
<p>Every http handler handles a request and a response</p>

</body>
</html>
`
var userInfo = `{
"name":"testUser",
"age":21
}`

func main() {

	address := ":8080"

	mux := http.NewServeMux() //handler

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-Type", "text/html")
		w.WriteHeader(http.StatusAccepted)
		_, err := w.Write([]byte(indexPage))
		if err != nil {
			log.Printf("%v", err)
		}

	})

	//Adding the user route
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-Type", "text/html")
		w.WriteHeader(http.StatusAccepted)
		_, err := w.Write([]byte(userInfo))
		if err != nil {
			log.Printf("%v", err)
		}

	})

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

package server

import (
	"log"
	"net/http"
)

// html to render
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

// Server is an HTTP sever
type Server struct {
}

// Creating a new server
func New() *Server {
	return &Server{}
}

// HandleIndex handles the index path "/"
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "text/html")
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(indexPage))
	if err != nil {
		log.Printf("%v", err)
	}
}

// HandleUser handles the user path "/"
func (s *Server) HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(userInfo))
	if err != nil {
		log.Printf("%v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP/1.1 server
	// srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// using pusher handler
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(serverPushHandler)}

	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	// Exactly how you would run an HTTP/1.1 server with TLS connection.
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server-cert.pem", "server-key.pem"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)
	fmt.Println("req: \n", r)
	// Send a message back to the client
	w.Write([]byte("Hello"))
}

// server push handler
func serverPushHandler(w http.ResponseWriter, r *http.Request) {
	// log the request protocol
	log.Printf("Got connection: %s", r.Proto)

	// Handle 2nd request
	if r.URL.Path == "/2nd" {
		log.Println("Handling 2nd")
		w.Write([]byte("hello again!"))
		return
	}

	log.Println("Handling 1st request")
	if pusher, ok := w.(http.Pusher); ok {
		log.Println("cant push to client")
	} else {
		if err := pusher.Push("/2nd", nil); err != nil {
			log.Printf("Failed push: %v", err)
		}
	}

	//send response body
	w.Write([]byte("Hellow"))
}

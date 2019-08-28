package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorldServer Basic server that returns hello world string
func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	handler := http.HandlerFunc(HelloWorldServer)
	if err := http.ListenAndServe(":80", handler); err != nil {
		log.Fatalf("could not listen on port 80 %v", err)
	}
}

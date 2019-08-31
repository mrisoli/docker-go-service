package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorldServer Basic server that returns hello world string
func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, v2!")
}

func main() {
	handler := http.HandlerFunc(HelloWorldServer)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}

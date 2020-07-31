package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Server Error", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("start listening...")
	http.ListenAndServe(":9090", mux)
}
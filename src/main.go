package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love you and \n%+v!", r)
}

func main() {
	listen := ":8080"
	mode := os.Getenv("GO_ENV")
	if mode == "production" {
		listen = ":80"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(listen, nil)
}

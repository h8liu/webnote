package main

import (
	"net/http"
	"log"
	"flag"
	"fmt"
	"html"
)

var addr = ":8000"

func handleApi (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	flag.StringVar(&addr, "addr", addr, "listen address")
	flag.Parse()

	server := http.FileServer(http.Dir("."))
	http.Handle("/", server)
	http.HandleFunc("/api/", handleApi)
	for {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

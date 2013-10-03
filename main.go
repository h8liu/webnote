package main

import "net/http"
import "log"
import "flag"

var addr = ":8000"

func main() {
	flag.StringVar(&addr, "addr", addr, "listen address")
	flag.Parse()

	server := http.FileServer(http.Dir("."))
	http.Handle("/", server)
	for {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

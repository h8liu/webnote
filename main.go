package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	// "html"
	"io/ioutil"
	"strings"
)

var addr = ":8000"
var verbose = false

func pathFor(name string) string {
	return fmt.Sprintf("dat/%s", name)
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/")
	if verbose {
		fmt.Println(r.Method, name)
	}
	switch r.Method {
	case "POST":
		dat, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error on ready body: ", err)
		} else {
			err = ioutil.WriteFile(pathFor(name), dat, 0600)
			if err != nil {
				fmt.Println("error on saving: ", err)
			}
		}
	case "GET":
		dat, err := ioutil.ReadFile(pathFor(name))
		if err != nil {
			fmt.Println("error on reading: ", err)
		}
		_, err = w.Write(dat)
		if err != nil {
			fmt.Println("error on returning: ", err)
		}
	}
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

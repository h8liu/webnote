package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
		path := pathFor(name)
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			break
		}
		if info.IsDir() {
			break
		}
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("error on reading: ", err)
		}
		_, err = w.Write(dat)
		if err != nil {
			fmt.Println("error on returning: ", err)
		}
	}
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/jquery.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write([]byte(jquery))
	case "/main.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write([]byte(script))
	case "/":
		fallthrough
	case "/index.html":
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(index))
	case "/style.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(style))
	case "/favicon.ico":
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(favicon)
	default:
		w.WriteHeader(404)
	}
}

func main() {
	flag.StringVar(&addr, "addr", addr, "listen address")
	flag.Parse()

	// server := http.FileServer(http.Dir("."))
	// http.Handle("/", server)
	http.HandleFunc("/", handleFile)
	http.HandleFunc("/api/", handleApi)
	for {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

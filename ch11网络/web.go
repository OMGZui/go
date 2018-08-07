package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.Header)
	remPartOfURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shoutHelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shoutHello/"):]
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shoutHello/", shoutHelloHandler)
	http.ListenAndServe("localhost:9999", nil)
}

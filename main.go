package main

import (
	"./lib"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("capturer online")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	phantomjs := lib.PhantomCapture{}
	c := make(chan string)
	go phantomjs.Capture(r.URL.Path[1:], c)
	fmt.Fprintf(w, "%s", <-c)
}

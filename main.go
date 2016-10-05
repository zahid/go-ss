package main

import (
	"./lib"
	"fmt"
	"net/http"
)

// func main() {
// 	phantomjs := lib.PhantomCapture{}
// 	c := make(chan string)
// 	go phantomjs.Capture("https://yahoo.com", c)
// 	fmt.Printf("Got: %s\n", <-c)
// }

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	phantomjs := lib.PhantomCapture{}
	c := make(chan string)
	go phantomjs.Capture("https://yahoo.com", c)

	fmt.Fprintf(w, "<img src=\"data:image/png;base64,%s\">", <-c)
}

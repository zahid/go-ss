package main

import (
	"./lib"
	"fmt"
)

func main() {
	phantomjs := lib.PhantomCapture{}

	c := make(chan string)

	go phantomjs.Capture("https://yahoo.com", c)

	fmt.Printf("Got: %s\n", <-c)

}

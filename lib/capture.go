package lib

import (
	"fmt"
)

type Capturer interface {
	Capture()
}

type PhantomCapture struct{}

func (*PhantomCapture) Capture(url string, c chan string) {
	fmt.Printf("Capturing %s\n", url)
	c <- "[an image streammmm........]"
}

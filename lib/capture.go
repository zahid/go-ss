package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Capturer interface {
	Capture()
}

type PhantomCapture struct{}

func (*PhantomCapture) Capture(url string, c chan string) {
	cmd := exec.Command("phantomjs", "lib/rasterize.js", url)

	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	// printOutput(output)
	c <- string(output)
}

/*
http://www.darrencoxall.com/golang/executing-commands-in-go/
*/
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

package main

import (
	"bytes"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	buf := new(bytes.Buffer)
	buf.ReadFrom(os.Stdin)

	str := buf.String()

	time.Sleep(1 * time.Second)
	robotgo.TypeStr(str)
}

package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Existing("Hello World")
	msg.Hi()
	display.Display("Hello from main")
}

package main

import (
	"fmt"
	color "github.com/mouhamedsylla/term-color"
)

func main() {
	text := "Hello, World!"
	color := color.Color(text)
	fmt.Println(color.Colorize(color.Green))
}

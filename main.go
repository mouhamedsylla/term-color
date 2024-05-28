package main

import (
	"fmt"
	"github.com/mouhamedsylla/term-color/color"
)

func main() {
	text := "Hello, World!"
	color := color.Color(text)
	fmt.Println(color.Colorize(color.Green))
}

package main

import (
	"fmt"
	"term-color/color"
)

func main() {
	text := "Hello, World!"
	color := color.Color(text)
	fmt.Println(color.Colorize(color.Green))
}

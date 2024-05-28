package main

import (
	"fmt"

	"github.com/mouhamedsylla/term-color/color"
)

func main() {
	text := "Hello, World!"
	clr := color.Color().SetText(text)
	clr.ColorTextPattern("Hello", clr.HEX_to_ANSI("#AF0123"))
	clr.ColorTextPattern("World", clr.Blue)
	clr.ColorTextPattern("Hello", clr.Underline)

	fmt.Println(clr.ToString())


}

# TERM-COLOR

This project provides functionalities to handle and manipulate text colors in the terminal using Go. It includes features to apply color, bold, italic, and underline styles, as well as convert HEX and RGB codes to ANSI codes.

## Project Structure

```
.
├── color
│   ├── color.go
│   └── colors.go
├── go.mod
├── main.go
└── README.md
```


## File Descriptions


## File Descriptions

### color/colors.go

This file defines the `Colors` struct and a function to initialize it with ANSI color codes. The struct contains fields for various colors and text styles, and a constructor function initializes these fields with corresponding ANSI escape codes.

### color/color.go

This file includes methods to manipulate the text within the `Colors` struct. It provides functionalities to:
- Set the text to bold, italic, or underline.
- Convert HEX color codes to ANSI codes.
- Convert RGB color codes to ANSI codes.
- Apply custom HEX or RGB colors to the text.
- Color specific portions of the text based on a pattern.
- Output the colored and styled text as a string.

## Usage

Here's how you can use the functionalities provided by this project:

1. **Initialize a Colors Object**:
    ```go
    clr := color.Color()
    ```

2. **Set the Text**:
    ```go
    clr.SetText("Your text here")
    ```

3. **Apply Basic Colors**:
    ```go
    coloredText := clr.Colorize(clr.Red)
    fmt.Println(coloredText)
    ```

4. **Apply Bold, Italic, and Underline Styles**:
    ```go
    boldText := clr.SetBold().ToString()
    fmt.Println(boldText)

    italicText := clr.SetItalic().ToString()
    fmt.Println(italicText)

    underlinedText := clr.SetUnderline().ToString()
    fmt.Println(underlinedText)
    ```

5. **Apply Custom HEX Color**:
    ```go
    hexColoredText := clr.HexCustom("#FF5733").ToString()
    fmt.Println(hexColoredText)
    ```

6. **Apply Custom RGB Color**:
    ```go
    rgbColoredText := clr.RGBCustom("rgb(255,87,51)").ToString()
    fmt.Println(rgbColoredText)
    ```

7. **Color Specific Portions of the Text**:
    ```go
    clr.SetText("Hello, this is a test string with multiple colors.")
    
    clr.ColorTextPattern("Hello", clr.Red)
    clr.ColorTextPattern("test", clr.Green)
    clr.ColorTextPattern("multiple", clr.Blue)
    
    fmt.Println(clr.ToString())
    ```

## Example

Here's an example to demonstrate the use of the library:

```go
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
```

## Authors

- [@mouhamedsylla](https://www.github.com/mouhamedsylla)
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

### color/colors.go

This file defines the `Colors` struct and a function to initialize it with ANSI color codes. The struct contains fields for various colors and text styles, and a constructor function initializes these fields with corresponding ANSI escape codes.

### color/color.go

This file includes methods to manipulate the text within the `Colors` struct. It provides functionalities to:
- Set the text to bold, italic, or underline.
- Convert HEX color codes to ANSI codes.
- Convert RGB color codes to ANSI codes.
- Apply custom HEX or RGB colors to the text.
- Output the colored and styled text as a string.

## Usage

Here's how you can use the functionalities provided by this project:

1. **Initialize a Colors Object**:
    ```go
    clr := color.Color("Your text here")
    ```

2. **Apply Basic Colors**:
    ```go
    coloredText := clr.Colorize(clr.Red).ToString()
    fmt.Println(coloredText)
    ```

3. **Apply Bold, Italic, and Underline Styles**:
    ```go
    boldText := clr.SetBold().ToString()
    fmt.Println(boldText)

    italicText := clr.SetItalic().ToString()
    fmt.Println(italicText)

    underlinedText := clr.SetUnderline().ToString()
    fmt.Println(underlinedText)
    ```

4. **Apply Custom HEX Color**:
    ```go
    hexColoredText := clr.HexCustom("#FF5733").ToString()
    fmt.Println(hexColoredText)
    ```

5. **Apply Custom RGB Color**:
    ```go
    rgbColoredText := clr.RGBCustom("rgb(255,87,51)").ToString()
    fmt.Println(rgbColoredText)
    ```

These examples demonstrate how to use the various features provided by the `Colors` struct to style and colorize text output in a Go application.

## Authors

- [@mouhamedsylla](https://www.github.com/mouhamedsylla)
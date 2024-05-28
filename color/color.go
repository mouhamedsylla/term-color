package color

import (
	"errors"
	"strconv"
	"strings"
)

func (color *Colors) SetBold() *Colors {
	color.text = color.Bold + color.text + color.Reset
	return color
}

func (color *Colors) SetItalic() *Colors {
	color.text = color.Italic + color.text + color.Reset
	return color
}

func (color *Colors) SetUnderline() *Colors {
	color.text = color.Underline + color.text + color.Reset
	return color
}

func (color *Colors) SetText(text string) *Colors {
	color.text = text
	return color
}


func (color *Colors) HEX_to_ANSI(hex string) string {
	r := []rune(hex[1:])
	Ansi_Format := ""
	pos := 0
	for i := 0; i < 3; i++ {
		code_str := string(r[pos]) + string(r[pos+1])
		n, _ := strconv.ParseInt(code_str, 16, 64)
		if i != 2 {
			Ansi_Format += strconv.Itoa(int(n)) + ";"
		} else {
			Ansi_Format += strconv.Itoa(int(n)) + "m"
		}
		pos += 2
	}
	Ansi_Format = "\033[38;2;" + Ansi_Format
	return Ansi_Format
}

func (color *Colors) RGB_to_ANSI(rgb string) (string, error) {
	rgb = rgb[4 : len(rgb)-1]
	tabRGB := strings.Split(strings.TrimSpace(rgb), ",")
	if IsValide_RGB(tabRGB) {
		return "", errors.New("invalid RGB code")
	}
	m := strings.Fields(tabRGB[0] + ";" + tabRGB[1] + ";" + tabRGB[2] + "m")
	Ansi_Format := "\033[38;2;" + strings.Join(m, "")
	return Ansi_Format, nil
}

func IsValide_RGB(tab []string) bool {
	for _, v := range tab {
		n, _ := strconv.Atoi(strings.TrimSpace(v))
		if n < 0 || n > 255 {
			return true
		}
	}
	return false
}

func (clr *Colors) Colorize(color string) string{
	clr.text = color + clr.text + clr.Reset
	return clr.text
}

func (color *Colors) ColorTextPattern(textPattern string, colorCode string) {
	textColored := colorCode + textPattern + color.Reset
	color.text = strings.ReplaceAll(color.text, textPattern, textColored)
}


func (color *Colors) HexCustom(hexCode string) *Colors {
	color.custom = color.HEX_to_ANSI(hexCode)
	color.text = color.custom + color.text + color.Reset
	return color
}

func (color *Colors) RGBCustom(rgbCode string) *Colors {
	custom, err := color.RGB_to_ANSI(rgbCode)
	if err != nil {
		return color
	}
	color.custom = custom
	color.text = color.custom + color.text + color.Reset
	return color
}

func (color *Colors) ToString() string {
	return color.text
}

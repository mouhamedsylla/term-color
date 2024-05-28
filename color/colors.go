package color

type Colors struct {
	text      string
	custom    string
	Black     string
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Purple    string
	Cyan      string
	White     string
	Reset     string
	Bold      string
	Italic    string
	Underline string
}

func Color() *Colors {
	return &Colors{
		Black:     "\033[0;30m",
		Red:       "\033[0;31m",
		Green:     "\033[0;32m",
		Yellow:    "\033[0;33m",
		Blue:      "\033[0;34m",
		Purple:    "\033[0;35m",
		Cyan:      "\033[0;36m",
		White:     "\033[0;37m",
		Reset:     "\033[0m",
		Bold:      "\033[1m",
		Italic:    "\033[3m",
		Underline: "\033[4m",
	}
}

package utils

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
)

var OutputTypeColor = map[string]string{
	"error":   ColorRed,
	"warning": ColorYellow,
	"success": ColorGreen,
	"info":    ColorBlue,
	"debug":   ColorCyan,
	"notice":  ColorMagenta,
	"reset":   ColorReset,
}

// ColorText returns the input text wrapped in the color code for the given output type.
func ColorText(text, outputType string) string {
	color, ok := OutputTypeColor[outputType]
	if !ok {
		color = ""
	}
	return color + text + ColorReset
}

package log

import (
	"github.com/OhYee/rainbow/color"
)

var (
	// Error logger
	Error = New().SetNewLine(false).SetColor(color.New().SetFrontRed()).SetPrefix(
		func(s string) string { return color.New().SetFrontRed().SetFontBold().Colorful("ERROR ") },
	).SetOutputToStderr()
	// Info logger
	Info = New().SetNewLine(false).SetPrefix(
		func(s string) string { return color.New().SetFontBold().Colorful("INFO ") },
	).SetOutputToStdout()
	// Debug logger
	Debug = New().SetNewLine(false).SetColor(color.New().SetFrontYellow()).SetPrefix(
		func(s string) string { return color.New().SetFrontYellow().SetFontBold().Colorful("DEBUG ") },
	).SetOutputToStdout()
)

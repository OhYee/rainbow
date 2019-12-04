package log

import (
	"fmt"
	"github.com/OhYee/rainbow/color"
	"io"
	"os"
)

type Logger struct {
	prefix func(string) string
	suffix func(string) string
	color  *color.Color
	output io.Writer
}

func New() *Logger {
	return new(Logger)
}

func (log *Logger) SetPrefix(prefix func(string) string) {
	log.prefix = prefix
}

func (log *Logger) SetSuffix(suffix func(string) string) {
	log.suffix = suffix
}

func (log *Logger) SetColor(color *color.Color) {
	log.color = color
}

func (log *Logger) SetOutput(output io.Writer) {
	log.output = output
}

func (log *Logger) SetOutputToStdout() {
	log.output = os.Stdout
}

func (log *Logger) SetOutputToStderr() {
	log.output = os.Stderr
}

func (log *Logger) SetOutputToNil() {
	log.output = nil
}

func (log *Logger) print(text string) {
	if log.output != nil {
		var prefix, suffix, color string
		if log.prefix != nil {
			prefix = log.prefix(text)
		}
		if log.suffix != nil {
			suffix = log.suffix(text)
		}
		if log.color != nil {
			color = log.color.Colorful(text)
		} else {
			color = text
		}

		log.output.Write([]byte(fmt.Sprintf("%s%s%s", prefix, color, suffix)))
	}
}

func (log *Logger) Print(args ...interface{}) {
	print(fmt.Sprint(args...))
}

func (log *Logger) Println(args ...interface{}) {
	print(fmt.Sprintln(args...))
}

func (log *Logger) Printf(format string, args ...interface{}) {
	print(fmt.Sprintf(format, args...))
}

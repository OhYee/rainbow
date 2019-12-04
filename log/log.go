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

func (log *Logger) SetPrefix(prefix func(string) string) *Logger {
	log.prefix = prefix
	return log
}

func (log *Logger) SetSuffix(suffix func(string) string) *Logger {
	log.suffix = suffix
	return log
}

func (log *Logger) SetColor(color *color.Color) *Logger {
	log.color = color
	return log
}

func (log *Logger) SetOutput(output io.Writer) *Logger {
	log.output = output
	return log
}

func (log *Logger) SetOutputToStdout() *Logger {
	log.output = os.Stdout
	return log
}

func (log *Logger) SetOutputToStderr() *Logger {
	log.output = os.Stderr
	return log
}

func (log *Logger) SetOutputToNil() *Logger {
	log.output = nil
	return log
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

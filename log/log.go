package log

import (
	"fmt"
	"io"
	"os"

	"github.com/OhYee/rainbow/color"
)

// Logger object
type Logger struct {
	prefix  func(string) string
	suffix  func(string) string
	color   *color.Color
	output  io.Writer
	newline bool
}

// New initial a new Logger
func New() *Logger {
	return &Logger{
		prefix:  nil,
		suffix:  nil,
		color:   nil,
		output:  os.Stdout,
		newline: true,
	}
}

// SetPrefix output text prefix generate function
func (log *Logger) SetPrefix(prefix func(string) string) *Logger {
	log.prefix = prefix
	return log
}

// SetSuffix output text suffix generate function
func (log *Logger) SetSuffix(suffix func(string) string) *Logger {
	log.suffix = suffix
	return log
}

// SetColor output text color
func (log *Logger) SetColor(color *color.Color) *Logger {
	log.color = color
	return log
}

// SetNewLine add a newline at the end of the log
func (log *Logger) SetNewLine(newline bool) *Logger {
	log.newline = newline
	return log
}

// SetOutput decide there to output
func (log *Logger) SetOutput(output io.Writer) *Logger {
	log.output = output
	return log
}

// SetOutputToStdout output to stdout
func (log *Logger) SetOutputToStdout() *Logger {
	log.output = os.Stdout
	return log
}

// SetOutputToStderr output to stderr
func (log *Logger) SetOutputToStderr() *Logger {
	log.output = os.Stderr
	return log
}

// SetOutputToNil do not output
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

		newline := ""
		if log.newline {
			newline = "\n"
		}
		log.output.Write([]byte(fmt.Sprintf("%s%s%s%s", prefix, color, suffix, newline)))
	}
}

// Print log something
func (log *Logger) Print(args ...interface{}) {
	log.print(fmt.Sprint(args...))
}

// Println log something with newline
func (log *Logger) Println(args ...interface{}) {
	log.print(fmt.Sprintln(args...))
}

// Printf log something with format
func (log *Logger) Printf(format string, args ...interface{}) {
	log.print(fmt.Sprintf(format, args...))
}

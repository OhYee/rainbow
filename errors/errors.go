package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Error struct
type Error struct {
	msg string
}

func getStack() string {
	pcs := make([]uintptr, 100)
	count := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs)

	s := make([]string, count)

	for idx := 0; idx < count; idx++ {
		frame, more := frames.Next()
		s[idx] = fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return strings.Join(s, "\n")
}

// New error with call stack
func New(format string, args ...interface{}) error {
	err := &Error{
		msg: fmt.Sprintf(
			"%s\n%s",
			fmt.Sprintf(format, args...),
			getStack(),
		),
	}
	return err
}

// NewErr new error with error stack from error
func NewErr(err error) error {
	if err != nil {
		return &Error{
			msg: fmt.Sprintf(
				"%s\n%s",
				err.Error(),
				getStack(),
			),
		}
	}
	return nil
}

// Error return the error message
func (err *Error) Error() string {
	return err.msg
}

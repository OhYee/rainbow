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

func getStack(_skip ...int) string {
	pcs := make([]uintptr, 100)
	skip := 3
	if len(_skip) != 0 {
		skip += _skip[0]
	}
	count := runtime.Callers(skip, pcs)
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
func NewErr(err error, skip ...int) error {
	if !IsError(err) && err != nil {
		return &Error{
			msg: fmt.Sprintf(
				"%s\n%s",
				err.Error(),
				getStack(skip...),
			),
		}
	}
	return err
}

// Error return the error message
func (err *Error) Error() string {
	return err.msg
}

// Wrapper wrap error with stack.
// func foo() (err error) {
//     defer errors.Wrapper(&err)
//     ... some code
// }
func Wrapper(err *error) {
	if !IsError(*err) && err != nil {
		*err = NewErr(*err, 1)
	}
}

// IsError error is Rainbow Error or not
func IsError(err error) bool {
	_, isRainbowError := (err).(*Error)
	return isRainbowError
}

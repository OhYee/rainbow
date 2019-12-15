package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	msg string
}

func New(format string, args ...interface{}) *Error {
	pcs := make([]uintptr, 100)
	count := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs)

	s := make([]string, count)

	for idx := 0; idx < count; idx++ {
		frame, more := frames.Next()
		s[idx] = fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}

	err := &Error{
		msg: fmt.Sprintf(
			"%s\n%s",
			fmt.Sprintf(format, args...),
			strings.Join(s, "\n"),
		),
	}
	return err
}

func (err *Error) Error() string {
	return err.msg
}

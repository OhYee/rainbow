package errors

import (
	"fmt"
	"runtime"
	"testing"
)

func ErrCause() (err error) {
	return New("err")
}

func TestError(t *testing.T) {
	err := ErrCause()
	_, file, _, _ := runtime.Caller(0)
	wantErr := fmt.Sprintf(
		"%s\n%s:10 github.com/OhYee/rainbow/errors.ErrCause\n%s:14 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
		file,
	)

	if err.Error()[:len(wantErr)] != wantErr {
		t.Errorf("want err %s, but got %s\n", wantErr, err)
	}

}

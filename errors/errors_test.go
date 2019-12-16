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
	var err error

	err = ErrCause()
	_, file, _, _ := runtime.Caller(0)
	wantErr := fmt.Sprintf(
		"%s\n%s:10 github.com/OhYee/rainbow/errors.ErrCause\n%s:16 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
		file,
	)
	wantErr2 := fmt.Sprintf(
		"%s\n%s:39 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
	)

	if err.Error()[:len(wantErr)] != wantErr {
		t.Errorf("want err %s, but got %s\n", wantErr, err)
	}

	err = NewErr(nil)
	if err != nil {
		t.Errorf("nil error returns %v", err)
	}

	err = NewErr(fmt.Errorf("err"))
	if err.Error()[:len(wantErr2)] != wantErr2 {
		t.Errorf("want err %s, but got %s\n", wantErr2, err)
	}

}

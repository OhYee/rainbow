package errors

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func ErrCause() (err error) {
	return New("err")
}

func TestError(t *testing.T) {
	var err error

	_, file, _, _ := runtime.Caller(0)
	wantErr := fmt.Sprintf(
		"%s\n%s:11 github.com/OhYee/rainbow/errors.ErrCause\n%s:42 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
		file,
	)
	wantErr2 := fmt.Sprintf(
		"%s\n%s:47 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
	)
	wantErr3 := fmt.Sprintf(
		"%s\n%s:66 github.com/OhYee/rainbow/errors.TestError\n",
		"err",
		file,
	)

	// nil error
	err = NewErr(nil)
	if err != nil {
		t.Errorf("nil error returns %v", err)
	}

	// stack test
	err = ErrCause()
	if !strings.HasPrefix(ShowStack(err), wantErr) {
		t.Errorf("want err %s, but got %s\n", wantErr, err)
	}

	err = NewErr(fmt.Errorf("err"))
	if !strings.HasPrefix(ShowStack(err), wantErr2) {
		t.Errorf("want err %s, but got %s\n", wantErr2, err)
	}

	// NewErr
	err2 := NewErr(err)

	if err != err2 {
		t.Errorf("want err %s, but got %s\n", err, err2)
	}

	// Wrapper
	err2 = err
	Wrapper(&err2)
	if err != err2 {
		t.Errorf("want err %s, but got %s\n", err, err2)
	}
	err2 = fmt.Errorf("err")
	Wrapper(&err2)
	if !strings.HasPrefix(ShowStack(err2), wantErr3) {
		t.Errorf("want err %s, but got %s\n", wantErr3, err2)
	}

}

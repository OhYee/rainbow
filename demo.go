package main

import (
	"fmt"
	"time"

	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/errors"
	"github.com/OhYee/rainbow/log"
)

var Debugger = log.New().SetColor(
	color.New().SetFrontYellow(),
).SetPrefix(func(s string) string {
	return fmt.Sprintf("%s: ", time.Now().Format("2006-01-02 15:04:05"))
}).SetSuffix(func(s string) string {
	return fmt.Sprintf(" (The message length is %d)\n", len(s))
})

var Error = log.New().SetColor(color.New().SetFrontRed()).SetPrefix(
	func(s string) string {
		return color.New().SetFontBold().SetFrontRed().Colorful(
			"ERROR: ",
		)
	},
).SetSuffix(func(string) string { return "\n" })

func CauseError() (err error) {
	// You can use wrapper to transfer an error to a rainbow error
	defer errors.Wrapper(&err)
	// cause error
	return fmt.Errorf("A error")
}

func main() {
	Debugger.Print("begin")
	err := CauseError()
	if err != nil {
		Error.Print(errors.ShowStack(err))
	}
	Debugger.Print("end")
}

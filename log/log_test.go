package log

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/OhYee/rainbow/color"
)

func TestLog(t *testing.T) {
	want := bytes.NewBuffer([]byte{})
	bold := color.New().SetFontBold()
	red := color.New().SetFrontRed()
	want.WriteString("1st. Line")
	want.WriteRune('\n')
	want.WriteString("2nd. Line\n")
	want.WriteRune('\n')
	want.WriteString("3rd. Line 3")
	want.WriteRune('\n')
	want.WriteString(red.Colorful("4th. Line"))
	want.WriteRune('\n')
	want.WriteString(bold.Colorful("[PREFIX] "))
	want.WriteString(red.Colorful("5th. Line"))
	want.WriteString(" 9")
	want.WriteRune('\n')
	want.WriteString(bold.Colorful("INFO "))
	want.WriteString("test info\n")
	want.WriteString(color.New().SetFrontYellow().SetFontBold().Colorful("WARNING "))
	want.WriteString(color.New().SetFrontYellow().Colorful("test warning\n"))
	want.WriteString(color.New().SetFrontBlue().SetFontBold().Colorful("DEBUG "))
	want.WriteString(color.New().SetFrontBlue().Colorful("test debug\n"))
	want.WriteString(color.New().SetFrontRed().SetFontBold().Colorful("ERROR "))
	want.WriteString(color.New().SetFrontRed().Colorful("test error\n"))

	buf := bytes.NewBuffer([]byte{})
	logger := New().SetOutput(buf)

	logger.Print("1st. Line")
	logger.Println("2nd. Line")
	logger.Printf("3rd. Line %d", 3)
	logger.SetColor(color.New().SetFrontRed())
	logger.Print("4th. Line")
	logger.SetPrefix(func(s string) string {
		return color.New().SetFontBold().Colorful("[PREFIX] ")
	})
	logger.SetSuffix(func(s string) string {
		return fmt.Sprintf(" %d", len(s))
	})
	logger.Print("5th. Line")

	logger.SetOutputToNil()
	logger.Print("6th. Line")

	Info.SetOutput(buf)
	Info.Println("test info")

	Warning.SetOutput(buf)
	Warning.Println("test warning")

	Debug.SetOutput(buf)
	Debug.Println("test debug")

	Error.SetOutput(buf)
	Error.Println("test error")

	b := buf.Bytes()
	w := want.Bytes()

	if !bytes.Equal(b, w) {
		l := Min(len(b), len(w))
		pos := -1
		for i := 0; i < l; i++ {
			if b[i] != w[i] {
				pos = i
			}
		}
		t.Errorf("Want:\n%s\nGot\n%s\n\nDiff at %d %s", w, b, pos, b[Max(0, pos-5):Min(l, pos+5)])
	}

}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

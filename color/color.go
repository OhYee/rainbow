package color

import (
	"fmt"
	"github.com/OhYee/goutils/functional"
	"strings"
)

//go:generate gcg ./template/data.json

// Color surround with color
type Color struct {
	front string
	back  string
	font  string
}

// New return a pointer to a Color
func New() *Color {
	return new(Color)
}

// Reset the color
func (c *Color) Reset() *Color {
	c.front = ""
	c.back = ""
	c.font = ""
	return c
}

// SetFrontColor256 set the front color using 256-bit color, not all terminals support
func (c *Color) SetFrontColor256(color byte) *Color {
	c.front = fmt.Sprintf("38;5;%d", color)
	return c
}

// SetBackColor256 set the background color using 256-bit color, not all terminals support
func (c *Color) SetBackColor256(color byte) *Color {
	c.front = fmt.Sprintf("48;5;%d", color)
	return c
}

// SetFrontColorRGB set the front color using rgb color, not all terminals support
func (c *Color) SetFrontColorRGB(r byte, g byte, b byte) *Color {
	c.front = fmt.Sprintf("48;2;%d;%d;%d", r, g, b)
	return c
}

// SetBackColorRGB set the background color using rgb color, not all terminals support
func (c *Color) SetBackColorRGB(r byte, g byte, b byte) *Color {
	c.front = fmt.Sprintf("48;2;%d;%d;%d", r, g, b)
	return c
}

// Colorful the text
func (c *Color) Colorful(text string) string {
	prefix := strings.Join(fp.FilterString(func(s string) bool {
		return s != ""
	}, []string{c.font, c.front, c.back}), ";")
	if prefix != "" {
		return fmt.Sprintf("\x1b[%sm%s\x1b[0m", prefix, text)
	}
	return text
}

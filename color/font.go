// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.7 (https://github.com/OhYee/gcg)

package color

import (
	"fmt"
)

const (
	bold       = 1
	weak       = 2
	italic     = 3
	underline  = 4
	blink      = 5
	quickBlink = 6
	reverse    = 7
	hide       = 8
	del        = 9
)

// SetFontBold set the font to bold, not all terminals support
func (c *Color) SetFontBold() *Color {
	c.font = fmt.Sprintf("%d", bold)
	return c
}

// SetFontWeak set the font to weak, not all terminals support
func (c *Color) SetFontWeak() *Color {
	c.font = fmt.Sprintf("%d", weak)
	return c
}

// SetFontItalic set the font to italic, not all terminals support
func (c *Color) SetFontItalic() *Color {
	c.font = fmt.Sprintf("%d", italic)
	return c
}

// SetFontUnderline set the font to underline, not all terminals support
func (c *Color) SetFontUnderline() *Color {
	c.font = fmt.Sprintf("%d", underline)
	return c
}

// SetFontBlink set the font to blink, not all terminals support
func (c *Color) SetFontBlink() *Color {
	c.font = fmt.Sprintf("%d", blink)
	return c
}

// SetFontQuickBlink set the font to quickBlink, not all terminals support
func (c *Color) SetFontQuickBlink() *Color {
	c.font = fmt.Sprintf("%d", quickBlink)
	return c
}

// SetFontReverse set the font to reverse, not all terminals support
func (c *Color) SetFontReverse() *Color {
	c.font = fmt.Sprintf("%d", reverse)
	return c
}

// SetFontHide set the font to hide, not all terminals support
func (c *Color) SetFontHide() *Color {
	c.font = fmt.Sprintf("%d", hide)
	return c
}

// SetFontDel set the font to del, not all terminals support
func (c *Color) SetFontDel() *Color {
	c.font = fmt.Sprintf("%d", del)
	return c
}

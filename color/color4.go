// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package color

import (
	"fmt"
)

const (
	black   = 0
	red     = 1
	green   = 2
	yellow  = 3
	blue    = 4
	magenta = 5
	cyan    = 6
	white   = 7

	lightBlack   = 0 + 60
	lightRed     = 1 + 60
	lightGreen   = 2 + 60
	lightYellow  = 3 + 60
	lightBlue    = 4 + 60
	lightMagenta = 5 + 60
	lightCyan    = 6 + 60
	lightWhite   = 7 + 60
)

// SetFrontBlack set the front color to {.name}
func (c *Color) SetFrontBlack() *Color {
	c.front = fmt.Sprintf("%d", black+30)
	return c
}

// SetLightFrontBlack set the front color to light {.name}
func (c *Color) SetFrontLightBlack() *Color {
	c.front = fmt.Sprintf("%d", lightBlack+30)
	return c
}

// SetLightBackBlack set the background color to light {.name}
func (c *Color) SetBackBlack() *Color {
	c.back = fmt.Sprintf("%d", black+40)
	return c
}

// SetLightBackBlack set the background color to light {.name}
func (c *Color) SetBackLightBlack() *Color {
	c.back = fmt.Sprintf("%d", lightBlack+40)
	return c
}

// SetFrontRed set the front color to {.name}
func (c *Color) SetFrontRed() *Color {
	c.front = fmt.Sprintf("%d", red+30)
	return c
}

// SetLightFrontRed set the front color to light {.name}
func (c *Color) SetFrontLightRed() *Color {
	c.front = fmt.Sprintf("%d", lightRed+30)
	return c
}

// SetLightBackRed set the background color to light {.name}
func (c *Color) SetBackRed() *Color {
	c.back = fmt.Sprintf("%d", red+40)
	return c
}

// SetLightBackRed set the background color to light {.name}
func (c *Color) SetBackLightRed() *Color {
	c.back = fmt.Sprintf("%d", lightRed+40)
	return c
}

// SetFrontGreen set the front color to {.name}
func (c *Color) SetFrontGreen() *Color {
	c.front = fmt.Sprintf("%d", green+30)
	return c
}

// SetLightFrontGreen set the front color to light {.name}
func (c *Color) SetFrontLightGreen() *Color {
	c.front = fmt.Sprintf("%d", lightGreen+30)
	return c
}

// SetLightBackGreen set the background color to light {.name}
func (c *Color) SetBackGreen() *Color {
	c.back = fmt.Sprintf("%d", green+40)
	return c
}

// SetLightBackGreen set the background color to light {.name}
func (c *Color) SetBackLightGreen() *Color {
	c.back = fmt.Sprintf("%d", lightGreen+40)
	return c
}

// SetFrontYellow set the front color to {.name}
func (c *Color) SetFrontYellow() *Color {
	c.front = fmt.Sprintf("%d", yellow+30)
	return c
}

// SetLightFrontYellow set the front color to light {.name}
func (c *Color) SetFrontLightYellow() *Color {
	c.front = fmt.Sprintf("%d", lightYellow+30)
	return c
}

// SetLightBackYellow set the background color to light {.name}
func (c *Color) SetBackYellow() *Color {
	c.back = fmt.Sprintf("%d", yellow+40)
	return c
}

// SetLightBackYellow set the background color to light {.name}
func (c *Color) SetBackLightYellow() *Color {
	c.back = fmt.Sprintf("%d", lightYellow+40)
	return c
}

// SetFrontBlue set the front color to {.name}
func (c *Color) SetFrontBlue() *Color {
	c.front = fmt.Sprintf("%d", blue+30)
	return c
}

// SetLightFrontBlue set the front color to light {.name}
func (c *Color) SetFrontLightBlue() *Color {
	c.front = fmt.Sprintf("%d", lightBlue+30)
	return c
}

// SetLightBackBlue set the background color to light {.name}
func (c *Color) SetBackBlue() *Color {
	c.back = fmt.Sprintf("%d", blue+40)
	return c
}

// SetLightBackBlue set the background color to light {.name}
func (c *Color) SetBackLightBlue() *Color {
	c.back = fmt.Sprintf("%d", lightBlue+40)
	return c
}

// SetFrontMagenta set the front color to {.name}
func (c *Color) SetFrontMagenta() *Color {
	c.front = fmt.Sprintf("%d", magenta+30)
	return c
}

// SetLightFrontMagenta set the front color to light {.name}
func (c *Color) SetFrontLightMagenta() *Color {
	c.front = fmt.Sprintf("%d", lightMagenta+30)
	return c
}

// SetLightBackMagenta set the background color to light {.name}
func (c *Color) SetBackMagenta() *Color {
	c.back = fmt.Sprintf("%d", magenta+40)
	return c
}

// SetLightBackMagenta set the background color to light {.name}
func (c *Color) SetBackLightMagenta() *Color {
	c.back = fmt.Sprintf("%d", lightMagenta+40)
	return c
}

// SetFrontCyan set the front color to {.name}
func (c *Color) SetFrontCyan() *Color {
	c.front = fmt.Sprintf("%d", cyan+30)
	return c
}

// SetLightFrontCyan set the front color to light {.name}
func (c *Color) SetFrontLightCyan() *Color {
	c.front = fmt.Sprintf("%d", lightCyan+30)
	return c
}

// SetLightBackCyan set the background color to light {.name}
func (c *Color) SetBackCyan() *Color {
	c.back = fmt.Sprintf("%d", cyan+40)
	return c
}

// SetLightBackCyan set the background color to light {.name}
func (c *Color) SetBackLightCyan() *Color {
	c.back = fmt.Sprintf("%d", lightCyan+40)
	return c
}

// SetFrontWhite set the front color to {.name}
func (c *Color) SetFrontWhite() *Color {
	c.front = fmt.Sprintf("%d", white+30)
	return c
}

// SetLightFrontWhite set the front color to light {.name}
func (c *Color) SetFrontLightWhite() *Color {
	c.front = fmt.Sprintf("%d", lightWhite+30)
	return c
}

// SetLightBackWhite set the background color to light {.name}
func (c *Color) SetBackWhite() *Color {
	c.back = fmt.Sprintf("%d", white+40)
	return c
}

// SetLightBackWhite set the background color to light {.name}
func (c *Color) SetBackLightWhite() *Color {
	c.back = fmt.Sprintf("%d", lightWhite+40)
	return c
}

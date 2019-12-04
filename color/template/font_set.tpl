
// SetFont{{upperFirstChar .name}} set the font to {{.name}}, not all terminals support
func (c *Color) SetFont{{upperFirstChar .name}}() *Color {
    c.font = fmt.Sprintf("%d", {{.name}})
    return c
}

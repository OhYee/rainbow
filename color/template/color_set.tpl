
// SetFront{{upperFirstChar .name}} set the front color to {.name}
func (c *Color) SetFront{{upperFirstChar .name}}() *Color {
    c.front = fmt.Sprintf("%d", {{.name}}+30)
    return c
}

// SetLightFront{{upperFirstChar .name}} set the front color to light {.name}
func (c *Color) SetFrontLight{{upperFirstChar .name}}() *Color {
    c.front = fmt.Sprintf("%d", light{{upperFirstChar .name}}+30)
    return c
}

// SetLightBack{{upperFirstChar .name}} set the background color to light {.name}
func (c *Color) SetBack{{upperFirstChar .name}}() *Color {
    c.back = fmt.Sprintf("%d", {{.name}}+40)
    return c
}

// SetLightBack{{upperFirstChar .name}} set the background color to light {.name}
func (c *Color) SetBackLight{{upperFirstChar .name}}() *Color {
    c.back = fmt.Sprintf("%d", light{{upperFirstChar .name}}+40)
    return c
}

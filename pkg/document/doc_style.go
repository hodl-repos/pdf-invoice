package document

type Color struct {
	r int // red 0-255
	g int // green 0-255
	b int // blue 0-255
}

func (c *Color) R() int {
	return colorLimit(c.r)
}
func (c *Color) G() int {
	return colorLimit(c.g)
}
func (c *Color) B() int {
	return colorLimit(c.b)
}

func colorLimit(i int) int {
	if i < 0 {
		return 0
	}
	if i > 255 {
		return 255
	}
	return i
}

type DocStyle struct {
	// TODO: use const
	fontFamily     string
	fontStyle      []FontStyleType
	fontSize       float64
	fontLineHeight float64

	fillColor Color

	drawColor     Color
	lineWidth     float64
	lineCapStyle  LineCapType
	lineJoinStyle LineJoinType
}

package document

type Block interface {
	GetWidth() float64
	GetHeight() float64
	Render() error
}

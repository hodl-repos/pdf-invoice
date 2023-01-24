package document

type Box struct {
	doc      *Doc
	x        float64
	y        float64
	w        float64
	h        float64
	boxType  BoxType
	position PositionType
}

func (d *Doc) NewBox(x, y, w, h float64, args ...interface{}) *Box {
	b := &Box{doc: d, x: x, y: y, w: w, h: h}
	for _, a := range args {
		switch param := a.(type) {
		case BoxType:
			b.boxType = param
		case PositionType:
			b.position = param
		}
	}
	b.setDefaults()
	return b
}

func (b *Box) setDefaults() {
	if b.boxType == BoxUnset {
		b.boxType = BoxClosed
	}
	if b.position == PositionUnset {
		b.position = PositionAbsolute
	}
}

func (b *Box) GetType() BoxType {
	return b.boxType
}

func (b *Box) SetFocus() {
	if b.doc.debug {
		b.doc.Rect(b.x, b.y, b.w, b.h, "D")
	}
	b.doc.SetXY(b.x, b.y)
	pWidth, pHeight := b.doc.GetPageSize()
	ml := pWidth - b.x - b.w
	mb := pHeight - b.y - b.h
	b.doc.SetMargins(b.x, b.y, ml)
	b.doc.SetAutoPageBreak(true, mb)
}

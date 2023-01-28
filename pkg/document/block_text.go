package document

type Text struct {
	doc  *Doc
	flow FlowType
	str  string
}

func NewText(doc *Doc, str string, args ...interface{}) *Text {
	return &Text{doc, FlowBlock, str}
}

func (txt *Text) GetWidth() float64 {
	return txt.doc.GetStringWidth(txt.str)
}

func (txt *Text) GetHeight() float64 {
	lines := txt.doc.SplitText(txt.str, txt.doc.GetPrintWidth())
	return float64(len(lines)) * txt.doc.GetFontLineHeight()
}

func (txt *Text) Render() error {
	txt.doc.MCell(txt.doc.GetPrintWidth(), txt.doc.GetFontLineHeight(), txt.str, "C", "", false)
	return nil
}

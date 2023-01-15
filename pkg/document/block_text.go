package document

type Text struct {
	doc  *Doc
	flow FlowType
	str  string
}

// TODO: use args with:
// func checkTypes(args ...interface{}) {
// 	for _, arg := range args {
// 		fmt.Println(reflect.TypeOf(arg))
// 	}
// }
//
// checkTypes(FlowBlock, AlignBottom, ColFixed)
//
// output:
// document.FlowType
// document.Alignment
// document.ColumnType

func NewText(doc *Doc, str string, args ...interface{}) *Text {
	return &Text{doc, FlowBlock, str}
}

func (r *Text) GetWidth() float64 {
	return r.doc.GetPrintWidth()
}

func (r *Text) GetHeight() float64 {
	return r.doc.GetLineHeight()
}

func (r *Text) Render() error {
	r.doc.MCell(r.doc.GetPrintWidth(), r.doc.GetLineHeight(), r.str, "", "", false)
	return nil
}

package document

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type PageSize struct {
	width, height float64
}

type LayoutV2 struct {
	orientation OrientationType
	unit        UnitType
	pageSize    PageSize
	box         *BoxV2
}

func (l *LayoutV2) NewPdf() *gofpdf.Fpdf {
	return gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        gofpdf.PageSizeA4,
		FontDirStr:     "",
	})
}

func (l *LayoutV2) NextBox() *BoxV2 {
	return l.box
}

func (l *LayoutV2) Render() {
	fmt.Println(l.box.Height())
}

func LayoutA4() *LayoutV2 {
	return &LayoutV2{
		orientation: OrientationPortrait,
		unit:        UnitMillimeter,
		pageSize:    PageSize{210, 297},
		box:         NewBoxV2(20, 16.9, 170, 263.2, BoxOpen),
	}
}

type BoxV2 struct {
	x        float64
	y        float64
	w        float64
	h        float64
	boxType  BoxType
	position PositionType
	blocks   []Block
}

func (b *BoxV2) Add(blk Block) {
	b.blocks = append(b.blocks, blk)
}

func (b *BoxV2) Height() float64 {
	ht := 0.
	for _, b := range b.blocks {
		ht += b.GetHeight()
	}
	return ht
}

func NewBoxV2(x, y, w, h float64, args ...interface{}) *BoxV2 {
	b := &BoxV2{x: x, y: y, w: w, h: h}
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

func (b *BoxV2) setDefaults() {
	if b.boxType == BoxUnset {
		b.boxType = BoxClosed
	}
	if b.position == PositionUnset {
		b.position = PositionAbsolute
	}
}

type DocV2 struct {
	*gofpdf.Fpdf
	layout  *LayoutV2
	currBox *BoxV2
}

func (d *DocV2) NextBox() {
	d.currBox = d.layout.NextBox()
}

func (d *DocV2) Add(blk Block) {
	d.currBox.Add(blk)
}

func (d *DocV2) Generate() {
	d.Fpdf = d.layout.NewPdf()
	d.SetFont("Arial", "", 8)
	d.layout.Render()
}

func (d *DocV2) GetPrintWidth() float64 {
	pageWidth, _ := d.Fpdf.GetPageSize()
	marginL, _, marginR, _ := d.Fpdf.GetMargins()
	return pageWidth - marginL - marginR
}

func NewDocV2() *DocV2 {
	doc := &DocV2{
		layout: LayoutA4(),
	}
	doc.NextBox()
	return doc
}

// BLOCKS

type TextV2 struct {
	doc  *DocV2
	flow FlowType
	str  string
}

func NewTextV2(doc *DocV2, str string, args ...interface{}) *TextV2 {
	return &TextV2{doc, FlowBlock, str}
}

func (txt *TextV2) GetWidth() float64 {
	return txt.doc.GetStringWidth(txt.str)
}

func (txt *TextV2) GetHeight() float64 {
	lines := txt.doc.SplitText(txt.str, txt.doc.GetPrintWidth())
	_, fontHt := txt.doc.GetFontSize()
	return float64(len(lines)) * fontHt
}

func (txt *TextV2) Render() error {
	fmt.Println(txt.str)
	return nil
}

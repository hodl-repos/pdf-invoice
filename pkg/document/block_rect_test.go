package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRect(t *testing.T) {
	doc := NewA4()

	r := NewRect(doc, 40, 30)

	// default
	err := r.Render()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "Rect.pdf")
}

func TestRectError(t *testing.T) {
	doc := NewA4()

	r := NewRect(doc, 10, 10)
	r.SetBorder(BorderNone)

	err := r.Render()
	assert.EqualError(t, err, "nothing to render of rect")
}

func TestRectFlow(t *testing.T) {
	doc := NewA4()

	r := NewRect(doc, 10, 10)

	for i := 0; i < 9; i++ {
		err := r.Render()
		assert.NoError(t, err)
	}
	r.SetFlow(FlowBlock)
	for i := 0; i < 9; i++ {
		err := r.Render()
		assert.NoError(t, err)
	}
	r.SetFlow(FlowNewline)
	err := r.Render()
	assert.NoError(t, err)

	err = r.Render()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "RectFlow.pdf")
}

func TestRectFill(t *testing.T) {
	doc := NewA4()

	a := doc.GetPrintWidth() / 255

	rect := NewRect(doc, a, a)
	rect.SetFill(true)
	rect.SetBorder(BorderNone)

	len := 255

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			doc.SetFillColor(i, 0, j)
			err := rect.Render()
			assert.NoError(t, err)
		}
		doc.Ln(a)
	}

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "RectFill.pdf")
}

func TestRectBorder(t *testing.T) {
	doc := NewA4()

	r := NewRect(doc, 10, 10)
	r.doc.SetFillColor(230, 230, 230)
	r.SetFill(true)
	r.doc.SetLineWidth(0.5)

	// default
	err := r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderInside)
	err = r.Render()
	assert.EqualError(t, err, "invalid border type BorderInside for rect")

	// Borders
	r.SetBorder(BorderTop)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderRight)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderBottom)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderLeft)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderX)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderY)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderTopAndLeft)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderTopAndRight)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderBottomAndRight)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderBottomAndLeft)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderOpenTop)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderOpenRight)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderOpenBottom)
	err = r.Render()
	assert.NoError(t, err)
	r.doc.SetX(r.doc.GetX() + 2)

	r.SetBorder(BorderOpenLeft)
	err = r.Render()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "RectBorder.pdf")
}

func TestRectSlot(t *testing.T) {
	doc := NewA4()

	r := NewRect(doc, 25, 15)
	txt := NewText(doc, "TEST")

	r.SetSlot(txt)

	err := r.Render()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "RectSlot.pdf")
}

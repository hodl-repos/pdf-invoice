package document

import "fmt"

type Rect struct {
	doc    *Doc
	flow   FlowType
	w      float64
	h      float64
	border BorderType
	fill   bool
	slot   Block
}

// TODO: Add RoundedRect

// NewRect creates a Rect block.
//
// Rect parameters & defaults:
//   - flow: FlowInline
//   - w: constructor parameter
//   - h: constructor parameter
//   - border: BorderOutline
//   - fill: false
//   - slot: nil
func NewRect(doc *Doc, w, h float64) *Rect {
	r := &Rect{
		doc: doc,
		w:   w,
		h:   h,
	}
	r.setDefaults()
	return r
}

func (r *Rect) setDefaults() {
	r.flow = FlowInline
	r.border = BorderOutside
	r.fill = false
	r.slot = nil
}

// BLOCK INTERFACE

func (r *Rect) GetWidth() float64 {
	return r.w
}

func (r *Rect) GetHeight() float64 {
	return r.h
}

func (r *Rect) Render() error {
	b := r.border
	if b == BorderInside {
		return fmt.Errorf("invalid border type BorderInside for rect")
	}
	if b == BorderNone && !r.fill {
		return fmt.Errorf("nothing to render of rect")
	}
	x1, y1 := r.doc.GetXY()
	x2, y2 := x1+r.w, y1+r.h
	// lineW := r.doc.GetLineWidth()
	// halfLineW := lineW / 2.

	if r.fill {
		r.doc.Rect(x1, y1, r.w, r.h, "F")
	}

	if b == BorderOutside {
		r.doc.Rect(x1, y1, r.w, r.h, "D")
	}

	if b != BorderOutside {
		// top border
		if b == BorderTop ||
			b == BorderTopAndLeft ||
			b == BorderTopAndRight ||
			b == BorderY ||
			b == BorderOpenRight ||
			b == BorderOpenBottom ||
			b == BorderOpenLeft {
			r.doc.Line(x1, y1, x2, y1)
		}

		// right border
		if b == BorderRight ||
			b == BorderTopAndRight ||
			b == BorderBottomAndRight ||
			b == BorderX ||
			b == BorderOpenTop ||
			b == BorderOpenBottom ||
			b == BorderOpenLeft {
			r.doc.Line(x2, y1, x2, y2)
		}

		// bottom border
		if b == BorderBottom ||
			b == BorderBottomAndLeft ||
			b == BorderBottomAndRight ||
			b == BorderY ||
			b == BorderOpenTop ||
			b == BorderOpenRight ||
			b == BorderOpenLeft {
			r.doc.Line(x1, y2, x2, y2)
		}

		// left border
		if b == BorderLeft ||
			b == BorderTopAndLeft ||
			b == BorderBottomAndLeft ||
			b == BorderX ||
			b == BorderOpenTop ||
			b == BorderOpenRight ||
			b == BorderOpenBottom {
			r.doc.Line(x1, y1, x1, y2)
		}
	}

	// ml, mt, mr, mb := r.doc.GetMargins()

	if r.slot != nil {
		// r.doc.SetMargins(x1, y1, )
		err := r.slot.Render()
		if err != nil {
			return err
		}
	}

	switch r.flow {
	case FlowInline:
		r.doc.SetXY(x2, y1)
	case FlowBlock:
		r.doc.SetXY(x1, y2)
	case FlowNewline:
		r.doc.SetY(y2)
	}

	return nil
}

// SETTERS

func (r *Rect) SetFlow(f FlowType) {
	r.flow = f
}

func (r *Rect) SetBorder(b BorderType) {
	r.border = b
}

func (r *Rect) SetFill(b bool) {
	r.fill = b
}

func (r *Rect) SetSlot(b Block) {
	r.slot = b
}

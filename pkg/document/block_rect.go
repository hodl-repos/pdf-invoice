package document

import (
	"fmt"
)

type Rect struct {
	doc        *Doc
	w          float64
	h          float64
	flow       FlowType
	padding    Padding
	border     BorderType
	fill       FillType
	slot       Block
	vAlignment VAlignmentType
}

// NewRect creates a Rect block.
//
// Rect parameters & defaults:
//   - flow: FlowInline
//   - w: constructor parameter
//   - h: constructor parameter
//   - border: BorderOutline
//   - fill: false
//   - slot: nil
func NewRect(doc *Doc, w, h float64, args ...interface{}) *Rect {
	r := &Rect{
		doc: doc,
		w:   w,
		h:   h,
	}

	// fmt.Println(r.padding)

	for _, a := range args {
		switch param := a.(type) {
		case FlowType:
			r.flow = param
		case Padding:
			r.padding = param
		case BorderType:
			if param == BorderInside {
				break
			}
			r.border = param
		case FillType:
			r.fill = param
		case VAlignmentType:
			r.vAlignment = param
		}
	}

	r.setDefaults()
	return r
}

func (r *Rect) setDefaults() {
	if r.flow == FlowUnset {
		r.flow = FlowInline
	}
	if r.border == BorderUnset {
		r.border = BorderOutside
	}
	if r.fill == FillUnset {
		r.fill = FillNone
	}

	if r.vAlignment == VAlignUnset {
		r.vAlignment = VAlignTop
	}
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
	if b == BorderNone && r.fill == FillNone {
		return fmt.Errorf("nothing to render of rect")
	}
	x1, y1 := r.doc.GetXY()
	x2, y2 := x1+r.w, y1+r.h
	// lineW := r.doc.GetLineWidth()
	// halfLineW := lineW / 2.

	if r.fill == Fill {
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

	if r.slot != nil {
		ml, mt, mr, _ := r.doc.GetMargins()
		pageWidth, _ := r.doc.GetPageSize()
		x, y := r.doc.GetXY()
		p := r.padding
		r.doc.SetMargins(x+p[paddingLeft], y+p[paddingTop], pageWidth-x-r.w+p[paddingRight])
		defer r.doc.SetMargins(ml, mt, mr)

		bHeight := r.slot.GetHeight()
		if bHeight > r.h {
			return fmt.Errorf("error rendering slot: slot is heigher that rect")
		}

		switch r.vAlignment {
		case VAlignTop:
			break
		case VAlignMiddle:
			dy := (r.h-bHeight-p[paddingTop]-p[paddingBottom])/2 + p[paddingTop]
			r.doc.SetXY(r.doc.GetX()+p[paddingLeft], r.doc.GetY()+dy)
		case VAlignBottom:
			dy := r.h - bHeight - p[paddingBottom]
			r.doc.SetXY(r.doc.GetX()+p[paddingLeft], r.doc.GetY()+dy)
		}

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

func (r *Rect) SetFill(f FillType) {
	r.fill = f
}

func (r *Rect) SetSlot(b Block) {
	r.slot = b
}

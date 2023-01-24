package document

import "fmt"

type Layout struct {
	doc   *Doc
	boxes []Box

	openBoxes  []int
	currentBox int

	headerFn func()
	footerFn func()
}

func NewLayout(d *Doc, args ...interface{}) *Layout {
	l := &Layout{doc: d}
	for _, a := range args {
		switch param := a.(type) {
		case *Box:
			l.boxes = append(l.boxes, *param)
		}
	}
	return l
}

func (l *Layout) NewPage() {
	l.doc.SetAcceptPageBreakFunc(l.continuationFunc())

	l.currentBox = 0
	for i, b := range l.boxes {
		if b.boxType == BoxOpen {
			l.openBoxes = append(l.openBoxes, i)
		}
	}
	l.boxes[l.currentBox].SetFocus()
}

func (l *Layout) continuationFunc() func() bool {
	f := func() bool {
		fmt.Println("currentBox:", l.currentBox)
		return false
	}
	return f
}

func (l *Layout) AddBox(b Box) {
	l.boxes = append(l.boxes, b)
	if b.GetType() == BoxOpen {
		l.openBoxes = append(l.openBoxes, len(l.boxes)-1)
	}
}

func (l *Layout) SetAllBoxes(boxes []Box) {
	l.boxes = boxes
	l.openBoxes = []int{}
	for i, b := range boxes {
		if b.GetType() == BoxOpen {
			l.openBoxes = append(l.openBoxes, i)
		}
	}
}

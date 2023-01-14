package document

// Alignment determines how the content is aligned inside a cell.
//
// cell:
// #----------------------------------------------#
// | AlignTopLeft      AlignTop     AlignTopRight |
// |                                              |
// | AlignLeft	      AlignCenter	     AlignRight |
// |                                              |
// | AlignBottomLeft AlignBottom AlignBottomRight |
// #----------------------------------------------#
//
// Possible values:
//   - AlignCenter
//   - AlignTop
//   - AlignRight
//   - AlignBottom
//   - AlignLeft
//   - AlignTopLeft
//   - AlignTopRight
//   - AlignBottomRight
//   - AlignBottomLeft
type Alignment int

const (
	AlignCenter Alignment = iota
	AlignTop
	AlignRight
	AlignBottom
	AlignLeft
	AlignTopLeft
	AlignTopRight
	AlignBottomRight
	AlignBottomLeft
)

// BorderType determines how a border will be rendered.
//
// Possible values:
//   - BorderNone
//   - BorderOutside
//   - BorderInside
//   - BorderTop
//   - BorderRight
//   - BorderBottom
//   - BorderLeft
//   - BorderX
//   - BorderY
//   - BorderTopAndLeft
//   - BorderTopAndRight
//   - BorderBottomAndRight
//   - BorderBottomAndLeft
//   - BorderOpenTop
//   - BorderOpenRight
//   - BorderOpenBottom
//   - BorderOpenLeft
type BorderType int

const (
	BorderNone BorderType = iota
	BorderOutside
	BorderInside // e.g. table
	BorderTop
	BorderRight
	BorderBottom
	BorderLeft
	BorderX // BorderLeft + BorderRight
	BorderY // BorderTop + BorderBottom
	BorderTopAndLeft
	BorderTopAndRight
	BorderBottomAndRight
	BorderBottomAndLeft
	BorderOpenTop    // BorderRight + BorderBottom + BorderLeft
	BorderOpenRight  // BorderBottom + BorderLeft + BorderTop
	BorderOpenBottom // BorderLeft + BorderTop + BorderRight
	BorderOpenLeft   // BorderTop + BorderRight + BorderBottom
)

// Padding is the padding inside of a cell. The indices follow the convention
// Top(0), Right(1), Bottom(2), Left(3).
//
// #-- border -------------------------#
// |            paddingTop             |
// |             #------#              |
// | paddingLeft | cell | paddingRight |
// |             #------#              |
// |          paddingBottom            |
// #-----------------------------------#
type Padding [4]float64

const (
	paddingTop = iota
	paddingRight
	paddingBottom
	paddingLeft
)

// FlowType determines where the cursor will positioned after generate.
//
// FlowInline - cursor on the same height after the element.
//
// FlowBlock - cursor under and at the beginning of the element.
//
// FlowNewline - cursor under the element and at the left margin
type FlowType int

const (
	FlowInline FlowType = iota
	FlowBlock
	FlowNewline
)

package document

type LayoutType string

const (
	LayoutTypeDIN5008A LayoutType = "DIN_5008A"
	LayoutTypeDIN5008B LayoutType = "DIN_5008B"
)

type OrientationType int

const (
	OrientationUnset OrientationType = iota
	OrientationPortrait
	OrientationLandscape
)

type UnitType int

const (
	UnitUnset UnitType = iota
	UnitMillimeter
	UnitPoint
	UnitCentimeter
	UnitInch
)

// BoxType determines if content has to fit inside the box or it will be
// continued on another box on the same or subsequent page.
//
// # BoxUnset (default) - BoxType is not set
//
// BoxOpen - content will be continued when it does not fit the box entirely.
//
// BoxClosed - content has to fit the box, otherwise this should lead to an
// error.
type BoxType int

const (
	BoxUnset BoxType = iota
	BoxOpen
	BoxClosed
)

// PositionType determines how an element gets positioned.
//
// PositionAbsolute - place an element at the coordinates x, y on the current
// page.
//
// PositionsRelative - place an element at the current cursor position.
type PositionType int

const (
	PositionUnset PositionType = iota
	PositionAbsolute
	PositionRelative
)

// AlignmentType determines how the content is aligned inside a cell.
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
type AlignmentType int

const (
	AlignCenter AlignmentType = iota
	AlignTop
	AlignRight
	AlignBottom
	AlignLeft
	AlignTopLeft
	AlignTopRight
	AlignBottomRight
	AlignBottomLeft
)

// VAlignmentType determines how content is aligned vertically in a cell:
//
// #------------------#
// |    VAlignTop     |
// |                  |
// |   VAlignMiddle	  |
// |                  |
// |   VAlignBottom   |
// #------------------#
type VAlignmentType int

const (
	VAlignUnset VAlignmentType = iota
	VAlignTop
	VAlignMiddle
	VAlignBottom
)

// HAlignmentType determines how content is aligned horizontally in a cell:
//
// #-----------------------------------------#
// |                                         |
// | HAlignLeft	  HAlignCenter	 HAlignRight |
// |                                         |
// #-----------------------------------------#
type HAlignmentType int

const (
	HAlignUnset HAlignmentType = iota
	HAlignLeft
	HAlignCenter
	HAlignRight
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
	BorderUnset BorderType = iota
	BorderNone
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

type FillType int

const (
	FillUnset FillType = iota
	FillNone
	Fill
	FillGradientLinear
	FillGradientRadial
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
	FlowUnset FlowType = iota
	FlowInline
	FlowBlock
	FlowNewline
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

// FontStyle determines the font style of a given font.
//
// Possible values:
//   - FontStyleUnset
//   - FontStyleRegular
//   - FontStyleItalic
//   - FontStyleBold
//   - FontStyleUnderscore
//   - FontStyleStrikeOut
type FontStyleType int

const (
	FontStyleUnset FontStyleType = iota
	FontStyleRegular
	FontStyleItalic
	FontStyleBold
	FontStyleUnderscore
	FontStyleStrikeOut
)

// LineCapType determines the line cap stype of a line.
//
// LineCapButt - line ends as a square at the end of the line.
//
// LineCapRound - line ends as a circle at the end of the line. The center of
// the circle is at the end of the line.
//
// LineCapSquare - line ends as a square at the end of the line. the center of
// the square is at the end of the line.
type LineCapType int

const (
	LineCapUnset LineCapType = iota
	LineCapButt
	LineCapRound
	LineCapSquare
)

// lineCapType2Fpdf converts LineCapType to Fpdf type following the convention
// seen in SetLineJoinStyle().
//
// The default value is "butt".
func lineCapType2Fpdf(t LineCapType) string {
	switch t {
	case LineCapRound:
		return "round"
	case LineCapSquare:
		return "square"
	default:
		return "butt"
	}
}

// LineJoinType determines the line join style of a line.
//
// LineJoinMiter - connects two line segments with a straight line, extending
// their width to meet at a sharp point, or "miter".
//
// LineJoinRound - rounds off the corner where two lines meet, creating a
// circular arc with a radius equal to the line width.
//
// LineJoinBevel - connects two line segments with a straight line, but
// instead of extending the width to meet at a sharp point, it cuts off the
// corner at an angle.
type LineJoinType int

const (
	LineJoinUnset LineJoinType = iota
	LineJoinMiter
	LineJoinRound
	LineJoinBevel
)

// lineJoinType2Fpdf converts LineJoinType to Fpdf type following the convention
// seen in SetLineJoinStyle().
//
// The default value is "miter".
func lineJoinType2Fpdf(t LineJoinType) string {
	switch t {
	case LineJoinRound:
		return "round"
	case LineJoinBevel:
		return "bevel"
	default:
		return "miter"
	}
}

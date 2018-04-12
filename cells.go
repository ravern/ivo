package ivo

// Cells represents a rectangular collection of Cells.
type Cells struct {
	Cols int
	Rows int

	cc [][]Cell
}

// NewCells creates a new collection of cells with the specified number of
// columns and rows.
func NewCells(cols, rows int) Cells {
	cc := make([][]Cell, rows)
	for i := range cc {
		cc[i] = make([]Cell, cols)
	}
	return Cells{
		Cols: cols,
		Rows: rows,
		cc:   cc,
	}
}

// Set sets a cell at the specified column and row. If the column and/or row
// exceeds the bounds, nothing will be set.
func (cc Cells) Set(col, row int, c Cell) {
	if row >= cc.Rows || col >= cc.Cols {
		return
	}
	cc.cc[row][col] = c
}

// Get returns the cell at the specified column and row.
func (cc Cells) Get(col, row int) (Cell, bool) {
	if row >= cc.Rows || col >= cc.Cols {
		return Cell{}, false
	}
	return cc.cc[row][col], true
}

// Cell is a cell on the terminal screen.
type Cell struct {
	Rune rune
	Fore CellColor // text color
	Back CellColor // cell color
	Attr CellAttr
}

// CellColor represents the color a cell might have.
type CellColor int

// Supported colors.
const (
	CellColorDefault CellColor = iota
	CellColorBlack
	CellColorRed
	CellColorGreen
	CellColorYellow
	CellColorBlue
	CellColorMagenta
	CellColorCyan
	CellColorWhite
)

func (c CellColor) String() string {
	switch c {
	case CellColorDefault:
		return "default"
	case CellColorBlack:
		return "black"
	case CellColorRed:
		return "red"
	case CellColorGreen:
		return "green"
	case CellColorYellow:
		return "yellow"
	case CellColorBlue:
		return "blue"
	case CellColorMagenta:
		return "magenta"
	case CellColorCyan:
		return "cyan"
	case CellColorWhite:
		return "white"
	default:
		return "invalid"
	}
}

// CellAttr represents an attribute a cell might have.
type CellAttr int

// Supported attributes.
const (
	CellAttrNone CellAttr = 1 << iota
	CellAttrBold
	CellAttrUnderline
)

func (ca CellAttr) String() string {
	switch ca {
	case CellAttrNone:
		return "none"
	case CellAttrBold:
		return "bold"
	case CellAttrUnderline:
		return "underline"
	default:
		return "invalid"
	}
}

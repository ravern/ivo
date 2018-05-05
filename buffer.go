package ivo

// Buffer represents a rectangular collection of Cells.
//
// Set and Get will panic with an out of bounds message if an
// invalid column or row is passed.
type Buffer struct {
	Cols int
	Rows int
	cc   []*Cell
}

// newBuffer creates a new collection of cells with the specified
// number of columns and rows.
func newBuffer(cols, rows int) *Buffer {
	return &Buffer{
		Cols: cols,
		Rows: rows,
		cc:   make([]*Cell, cols*rows),
	}
}

// Set sets a cell at the specified column and row.
func (b *Buffer) Set(col, row int, c *Cell) {
	b.check(col, row)

	b.cc[col+row*b.Cols] = c
}

// Get returns the cell at the specified column and row.
func (b *Buffer) Get(col, row int) *Cell {
	b.check(col, row)

	return b.cc[col+row*b.Cols]
}

// resize resizes the Buffer with the existing cells pinned to
// the top left corner of the Buffer.
func (b *Buffer) resize(cols, rows int) {
	nb := newBuffer(cols, rows)

	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Cols; col++ {
			if row < b.Rows && col < b.Cols {
				c := b.Get(col, row)
				nb.Set(col, row, c)
			}
		}
	}

	*b = *nb
}

// check checks whether the given column and row are within
// bounds and panics if they are not.
func (b *Buffer) check(col, row int) {
	if col >= b.Cols || col < 0 || row >= b.Rows || row < 0 {
		panic("runtime error: index out of bounds")
	}
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
	}
	return "invalid"
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
	}
	return "invalid"
}

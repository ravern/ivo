package text

// Raw returns the raw rune slice contained in Text.
func (txt *Text) Raw() []rune {
	return txt.rr[:len(txt.rr)-1]
}

// RegionRaw returns the raw rune slice contained in Text, within the
// Region provided.
func (txt *Text) RegionRaw(reg Region) []rune {
	return txt.Subtext(reg).Raw()
}

// RawLines returns the raw rune slices of the lines contained in Text.
func (txt *Text) RawLines() [][]rune {
	lines := make([][]rune, len(txt.lines))

	offset := 0
	for i, line := range txt.lines {
		lines[i] = txt.rr[offset : offset+line]
		offset += line
	}

	return lines
}

// RegionRawLines returns the raw runes slices of the lines contained
// in Text, within the Region provided.
func (txt *Text) RegionRawLines(reg Region) [][]rune {
	return txt.Subtext(reg).RawLines()
}

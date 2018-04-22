package text

type Text struct {
	raw []rune
}

func NewText(raw []rune) *Text {
	t := &Text{rr: raw}
	return t
}

func (t *Text) Raw() []rune {
	return t.rr
}

package text

type Text struct {
	rr []rune
}

func New(s string) *Text {
	return &Text{
		rr: []rune(s),
	}
}

func (t *Text) Len() int {
	return len(t.rr)
}

func (t *Text) String() string {
	return string(t.rr)
}

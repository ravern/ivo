package nano

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
)

// Modes used in the editor.
const (
	EditorMode = ""
)

// EditorHandler provides actions related to the editor.
type EditorHandler interface {
	Next(ivo.Context, []ivo.Key)
	NextLine(ivo.Context, []ivo.Key)
	Prev(ivo.Context, []ivo.Key)
	PrevLine(ivo.Context, []ivo.Key)
	Cut(ivo.Context, []ivo.Key)
	Paste(ivo.Context, []ivo.Key)
	Raw(ivo.Context, []ivo.Key)
}

// NewEditorMapper creates a new key.Mapper for the editor.
func NewEditorMapper(h EditorHandler) *key.Mapper {
	m := key.NewMap()

	m.Set(EditorMode, []ivo.Key{{Code: ivo.KeyCodeArrowLeft}}, key.HandlerFunc(h.Prev))
	m.Set(EditorMode, []ivo.Key{{Code: ivo.KeyCodeArrowRight}}, key.HandlerFunc(h.Next))
	m.Set(EditorMode, []ivo.Key{{Code: ivo.KeyCodeArrowUp}}, key.HandlerFunc(h.PrevLine))
	m.Set(EditorMode, []ivo.Key{{Code: ivo.KeyCodeArrowDown}}, key.HandlerFunc(h.NextLine))
	m.Set(EditorMode, []ivo.Key{{Rune: 'k', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Cut))
	m.Set(EditorMode, []ivo.Key{{Rune: 'u', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Paste))
	m.SetFallback(EditorMode, key.HandlerFunc(h.Raw))

	return key.NewMapper(m)
}

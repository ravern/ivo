package nano

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
	"ivoeditor.com/ivo/key/handler"
)

// Modes used in the editor.
const (
	EditorMode = ""
)

// EditorHandler provides actions related to the editor.
type EditorHandler interface {
	handler.Cursor
	handler.Text
}

// NewEditorMapper creates a new key.Mapper for the editor.
func NewEditorMapper(h EditorHandler) *key.Mapper {
	m := key.NewMap()

	// Root mode
	m.SetFallback(EditorMode, key.HandlerFunc(h.Raw))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, key.HandlerFunc(h.Prev))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, key.HandlerFunc(h.Next))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, key.HandlerFunc(h.PrevLine))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, key.HandlerFunc(h.NextLine))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'k', Mod: ivo.KeyModCtrl},
	}, key.HandlerFunc(h.Cut))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'u', Mod: ivo.KeyModCtrl},
	}, key.HandlerFunc(h.Paste))

	return key.NewMapper(m)
}

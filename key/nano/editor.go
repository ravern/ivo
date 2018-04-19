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
	m.SetFallback(EditorMode, key.Handler(h.Raw))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, key.Handler(h.Prev))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, key.Handler(h.Next))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, key.Handler(h.PrevLine))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, key.Handler(h.NextLine))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'k', Mod: ivo.KeyModCtrl},
	}, key.Handler(h.Cut))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'u', Mod: ivo.KeyModCtrl},
	}, key.Handler(h.Paste))

	return key.NewMapper(m)
}

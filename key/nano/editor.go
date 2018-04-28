package nano

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/handler"
	"ivoeditor.com/ivo/key"
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

// NewEditorMap creates a new key.Mapper for the editor.
func NewEditorMap(h EditorHandler) *key.Map {
	m := key.NewMap()

	// Root mode
	m.SetFallback(EditorMode, h.Raw)

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, handler.KeyFunc(h.MovePrev))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, handler.KeyFunc(h.MoveNext))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, handler.KeyFunc(h.MovePrevLine))

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, handler.KeyFunc(h.MoveNextLine))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'k', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Cut))

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'u', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Paste))

	return m
}

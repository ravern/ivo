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

// NewEditorMap creates a new key.Mapper for the editor.
func NewEditorMap(h EditorHandler) *key.Map {
	m := key.NewMap()

	// Root mode
	m.SetFallback(EditorMode, h.Raw)
	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, h.MovePrev)
	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, h.MoveNext)
	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, h.MovePrevLine)
	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, h.MoveNextLine)
	m.Set(EditorMode, []ivo.Key{
		{Rune: 'k', Mod: ivo.KeyModCtrl},
	}, h.Cut)
	m.Set(EditorMode, []ivo.Key{
		{Rune: 'u', Mod: ivo.KeyModCtrl},
	}, h.Paste)

	return m
}

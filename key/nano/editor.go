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
	key.CursorHandler
	key.TextHandler
}

// NewEditorMapper creates a new key.Mapper for the editor.
func NewEditorMapper(h EditorHandler) *key.Mapper {
	m := key.NewMap()

	// Root mode
	m.SetFallback(EditorMode, h.Raw)

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, h.Prev)

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, h.Next)

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, h.PrevLine)

	m.Set(EditorMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, h.NextLine)

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'k', Mod: ivo.KeyModCtrl},
	}, h.Cut)

	m.Set(EditorMode, []ivo.Key{
		{Rune: 'u', Mod: ivo.KeyModCtrl},
	}, h.Paste)

	return key.NewMapper(m)
}

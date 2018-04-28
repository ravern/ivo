package vim

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/handler"
	"ivoeditor.com/ivo/key"
)

// Modes used in the editor.
const (
	EditorNormalMode = "normal"
	EditorInsertMode = "insert"
)

// EditorHandler provides actions related to the editor.
type EditorHandler interface {
	handler.Cursor
	handler.Text

	Normal(ivo.Context, []ivo.Key)
	Insert(ivo.Context, []ivo.Key)
}

// NewEditorMap creates a new key.Map for the editor.
func NewEditorMap(h EditorHandler) *key.Map {
	m := key.NewMap()

	// Normal mode
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'i'},
	}, h.Insert)

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, handler.KeyFunc(h.MovePrev))
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'h'},
	}, handler.KeyFunc(h.MovePrev))

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, handler.KeyFunc(h.MoveNext))
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'l'},
	}, handler.KeyFunc(h.MoveNext))

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, handler.KeyFunc(h.MovePrevLine))
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'k'},
	}, handler.KeyFunc(h.MovePrevLine))

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, handler.KeyFunc(h.MoveNextLine))
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'j'},
	}, handler.KeyFunc(h.MoveNextLine))

	// Insert mode
	m.SetFallback(EditorInsertMode, h.Raw)

	m.Set(EditorInsertMode, []ivo.Key{
		{Code: ivo.KeyCodeEsc},
	}, h.Normal)

	return m
}

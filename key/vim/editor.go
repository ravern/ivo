package vim

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
	"ivoeditor.com/ivo/key/handler"
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

// NewEditorMapper creates a new key.Mapper for the editor.
func NewEditorMapper(h EditorHandler) *key.Mapper {
	m := key.NewMap()

	// Normal mode
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'i'},
	}, h.Insert)

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowLeft},
	}, h.MovePrev)
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'h'},
	}, h.MovePrev)

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowRight},
	}, h.MoveNext)
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'l'},
	}, h.MoveNext)

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowUp},
	}, h.MovePrevLine)
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'k'},
	}, h.MovePrevLine)

	m.Set(EditorNormalMode, []ivo.Key{
		{Code: ivo.KeyCodeArrowDown},
	}, h.MoveNextLine)
	m.Set(EditorNormalMode, []ivo.Key{
		{Rune: 'j'},
	}, h.MoveNextLine)

	// Insert mode
	m.SetFallback(EditorInsertMode, h.Raw)

	m.Set(EditorInsertMode, []ivo.Key{
		{Code: ivo.KeyCodeEsc},
	}, h.Normal)

	return key.NewMapper(m)
}

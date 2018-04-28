package nano

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/handler"
	"ivoeditor.com/ivo/key"
)

// Modes used in the container.
const (
	ContainerForwardMode = "forward"
	ContainerCommandMode = "command"
)

// ContainerHandler provides actions related to the container.
type ContainerHandler interface {
	handler.Proxy
	handler.Prompt

	Help(ivo.Context)
	Search(ivo.Context)
	Quit(ivo.Context)
	Write(ivo.Context)
}

// NewContainerMap creates a new key.Map for the container.
func NewContainerMap(h ContainerHandler) *key.Map {
	m := key.NewMap()

	// Forward mode
	m.SetFallback(ContainerForwardMode, h.Forward)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'g', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Help))

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'w', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Search))

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'x', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Quit))

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'o', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Write))

	// Command mode
	m.SetFallback(ContainerCommandMode, h.Raw)

	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEnter},
	}, handler.KeyFunc(h.Confirm))

	m.Set(ContainerCommandMode, []ivo.Key{
		{Rune: 'c', Mod: ivo.KeyModCtrl},
	}, handler.KeyFunc(h.Cancel))

	return m
}

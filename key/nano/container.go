package nano

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
)

// Modes used in the container.
const (
	ContainerForwardMode = "forward"
	ContainerCommandMode = "command"
)

// ContainerHandler provides actions related to the container.
type ContainerHandler interface {
	key.ProxyHandler
	key.PromptHandler

	Help(ivo.Context, []ivo.Key)
	Search(ivo.Context, []ivo.Key)
	Quit(ivo.Context, []ivo.Key)
	Write(ivo.Context, []ivo.Key)
}

// NewContainerMapper creates a new key.Mapper for the container.
func NewContainerMapper(h ContainerHandler) *key.Mapper {
	m := key.NewMap()

	// Forward mode
	m.SetFallback(ContainerForwardMode, h.Forward)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'g', Mod: ivo.KeyModCtrl},
	}, h.Help)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'w', Mod: ivo.KeyModCtrl},
	}, h.Search)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'x', Mod: ivo.KeyModCtrl},
	}, h.Quit)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: 'o', Mod: ivo.KeyModCtrl},
	}, h.Write)

	// Command mode
	m.SetFallback(ContainerCommandMode, h.Raw)

	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEnter},
	}, h.Confirm)

	m.Set(ContainerCommandMode, []ivo.Key{
		{Rune: 'c', Mod: ivo.KeyModCtrl},
	}, h.Cancel)

	return key.NewMapper(m)
}

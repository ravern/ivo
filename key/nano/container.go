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
	Help(ivo.Context, []ivo.Key)
	Search(ivo.Context, []ivo.Key)
	Quit(ivo.Context, []ivo.Key)
	Write(ivo.Context, []ivo.Key)
	Forward(ivo.Context, []ivo.Key)

	Confirm(ivo.Context, []ivo.Key)
	Cancel(ivo.Context, []ivo.Key)
	Raw(ivo.Context, []ivo.Key)
}

// NewContainerMapper creates a new key.Mapper for the container.
func NewContainerMapper(h ContainerHandler) *key.Mapper {
	m := key.NewMap()

	m.Set(ContainerForwardMode, []ivo.Key{{Rune: 'g', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Help))
	m.Set(ContainerForwardMode, []ivo.Key{{Rune: 'w', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Search))
	m.Set(ContainerForwardMode, []ivo.Key{{Rune: 'x', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Quit))
	m.Set(ContainerForwardMode, []ivo.Key{{Rune: 'o', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Write))
	m.SetFallback(ContainerForwardMode, key.HandlerFunc(h.Forward))

	m.Set(ContainerCommandMode, []ivo.Key{{Code: ivo.KeyCodeEnter}}, key.HandlerFunc(h.Confirm))
	m.Set(ContainerCommandMode, []ivo.Key{{Rune: 'c', Mod: ivo.KeyModCtrl}}, key.HandlerFunc(h.Cancel))
	m.SetFallback(ContainerCommandMode, key.HandlerFunc(h.Raw))

	return key.NewMapper(m)
}

package vim

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

	Command(ivo.Context)
}

// NewContainerMap creates a new key.Mapper for the container.
func NewContainerMap(h ContainerHandler) *key.Map {
	m := key.NewMap()

	// Forward mode
	m.SetFallback(ContainerForwardMode, h.Forward)

	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: ':'},
	}, handler.KeyFunc(h.Command))

	// Command mode
	m.SetFallback(ContainerCommandMode, h.Raw)

	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEnter},
	}, handler.KeyFunc(h.Confirm))

	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEsc},
	}, handler.KeyFunc(h.Cancel))

	return m
}

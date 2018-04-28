package vim

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
	"ivoeditor.com/ivo/key/handler"
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

	Command(ivo.Context, []ivo.Key)
}

// NewContainerMap creates a new key.Mapper for the container.
func NewContainerMap(h ContainerHandler) *key.Map {
	m := key.NewMap()

	// Forward mode
	m.SetFallback(ContainerForwardMode, h.Forward)
	m.Set(ContainerForwardMode, []ivo.Key{
		{Rune: ':'},
	}, h.Command)

	// Command mode
	m.SetFallback(ContainerCommandMode, h.Raw)
	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEnter},
	}, h.Confirm)
	m.Set(ContainerCommandMode, []ivo.Key{
		{Code: ivo.KeyCodeEsc},
	}, h.Cancel)

	return m
}

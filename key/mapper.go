package key

import (
	"sync"
	"time"

	"ivoeditor.com/ivo"
)

type Mapper struct {
	Timeout time.Duration

	m    *Map
	mode string

	init sync.Once
	ctxs chan ivo.Context
	keys chan ivo.Key
}

func NewMapper(m *Map) *Mapper {
	return &Mapper{
		Timeout: 2 * time.Second,
		m:       m,
		ctxs:    make(chan ivo.Context),
		keys:    make(chan ivo.Key),
	}
}

func (mr *Mapper) SetMode(mode string) {
	mr.mode = mode
}

func (mr *Mapper) Process(ctx ivo.Context, k ivo.Key) {
	mr.init.Do(func() {
		go mr.process()
	})
	mr.ctxs <- ctx
	mr.keys <- k
}

func (mr *Mapper) process() {
	var (
		kk     []ivo.Key
		ctx    ivo.Context
		action func(ivo.Context)
	)
	for {
		var k ivo.Key

		if len(kk) > 0 {
			// More keys awaiting
			select {
			case ctx = <-mr.ctxs:
				k = <-mr.keys
			case <-time.After(mr.Timeout):
				if action != nil {
					action(ctx)
				}
			}
		} else {
			// New key
			ctx = <-mr.ctxs
			k = <-mr.keys
		}

		kk = append(kk, k)
		action, more, ok := mr.m.Get(mr.mode, kk)

		if !ok {
			ctx.Logger().Infof("key: failed to find mapping for %v", kk)
			kk = make([]ivo.Key, 0)
			ctx = nil
			action = nil
			continue
		}
		if more {
			continue
		}

		action(ctx)
	}
}

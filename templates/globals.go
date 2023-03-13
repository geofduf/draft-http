package templates

import (
	"sync"
)

type globals struct {
	data map[string]any
	mu   sync.RWMutex
}

func (g *globals) SetBool(k string, v bool) {
	g.set(k, v)
}

func (g *globals) SetFloat64(k string, v float64) {
	g.set(k, v)
}

func (g *globals) SetInt(k string, v int) {
	g.set(k, v)
}

func (g *globals) SetString(k string, v string) {
	g.set(k, v)
}

func (g *globals) Get(k string) (any, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	v, ok := g.data[k]
	return v, ok
}

func (g *globals) Delete(k string) {
	g.mu.Lock()
	delete(g.data, k)
	g.mu.Unlock()
}

func (g *globals) set(k string, v any) {
	g.mu.Lock()
	g.data[k] = v
	g.mu.Unlock()
}

package templates

import (
	"testing"
)

func TestSetBool(t *testing.T) {
	tests := []struct {
		id string
		k  string
		v  bool
	}{
		{"update", "k1", false},
		{"add", "k2", true},
	}
	g := globals{data: map[string]any{"k1": true}}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			g.SetBool(tt.k, tt.v)
			g.mu.RLock()
			defer g.mu.RUnlock()
			v, ok := g.data[tt.k]
			if !ok {
				t.Fatal("key not found")
			}
			got, ok := v.(bool)
			if !ok {
				t.Fatal("type is not bool")
			}
			if got != tt.v {
				t.Errorf("got %t, want %t", got, tt.v)
			}
		})
	}
}

func TestSetFloat64(t *testing.T) {
	tests := []struct {
		id string
		k  string
		v  float64
	}{
		{"update", "k1", 0.3},
		{"add", "k2", 0.2},
	}
	g := globals{data: map[string]any{"k1": 0.1}}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			g.SetFloat64(tt.k, tt.v)
			g.mu.RLock()
			defer g.mu.RUnlock()
			v, ok := g.data[tt.k]
			if !ok {
				t.Fatal("key not found")
			}
			got, ok := v.(float64)
			if !ok {
				t.Fatal("type is not float64")
			}
			if got != tt.v {
				t.Errorf("got %f, want %f", got, tt.v)
			}
		})
	}
}

func TestSetInt(t *testing.T) {
	tests := []struct {
		id string
		k  string
		v  int
	}{
		{"update", "k1", 3},
		{"add", "k2", 2},
	}
	g := globals{data: map[string]any{"k1": 1}}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			g.SetInt(tt.k, tt.v)
			g.mu.RLock()
			defer g.mu.RUnlock()
			v, ok := g.data[tt.k]
			if !ok {
				t.Fatal("key not found")
			}
			got, ok := v.(int)
			if !ok {
				t.Fatal("type is not int")
			}
			if got != tt.v {
				t.Errorf("got %d, want %d", got, tt.v)
			}
		})
	}
}

func TestSetString(t *testing.T) {
	tests := []struct {
		id string
		k  string
		v  string
	}{
		{"update", "k1", "updated"},
		{"add", "k2", "added"},
	}
	g := globals{data: map[string]any{"k1": "initial"}}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			g.SetString(tt.k, tt.v)
			g.mu.RLock()
			defer g.mu.RUnlock()
			v, ok := g.data[tt.k]
			if !ok {
				t.Fatal("key not found")
			}
			got, ok := v.(string)
			if !ok {
				t.Fatal("type is not string")
			}
			if got != tt.v {
				t.Errorf("got %s, want %s", got, tt.v)
			}
		})
	}
}

func TestGet(t *testing.T) {
	want := 10000
	g := globals{data: map[string]any{"k1": want}}
	v, ok := g.Get("k1")
	if !ok {
		t.Fatal("key not found")
	}
	got, ok := v.(int)
	if !ok {
		t.Fatal("type is not int")
	}
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
	if _, ok := g.Get("k2"); ok {
		t.Error("key should not be found")
	}
}

func TestDelete(t *testing.T) {
	g := globals{data: map[string]any{"k1": true}}
	g.Delete("k1")
	g.mu.RLock()
	defer g.mu.RUnlock()
	if _, ok := g.data["k1"]; ok {
		t.Error("key should not be found")
	}
}

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

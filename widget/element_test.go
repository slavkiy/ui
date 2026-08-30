package widget

import "testing"

func TestElementResolveSizeUsesViewportPercent(t *testing.T) {
	e := NewElement(kindColumn, 0)
	e.Size(50, 25)

	w, h := e.ResolveSize(1000, 600)
	if w != 500 {
		t.Fatalf("width = %v, want 500", w)
	}
	if h != 150 {
		t.Fatalf("height = %v, want 150", h)
	}
}

func TestResolvePercentClampsToViewport(t *testing.T) {
	if got := ResolvePercent(120, 1000); got != 1000 {
		t.Fatalf("ResolvePercent(120,1000) = %v, want 1000", got)
	}
	if got := ResolvePercent(-10, 1000); got != 0 {
		t.Fatalf("ResolvePercent(-10,1000) = %v, want 0", got)
	}
}

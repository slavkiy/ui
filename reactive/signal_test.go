package reactive

import (
	"sync"
	"testing"
)

func TestGetSignalRegistersByName(t *testing.T) {
	signal := NewSignal("demo", 42)
	got := GetSignal[int]("demo")
	if got == nil {
		t.Fatal("expected signal registered under name")
	}
	if got != signal {
		t.Fatal("expected same signal instance")
	}
}

func TestSignalSameNameDifferentTypeAllowed(t *testing.T) {
	strSig := NewSignal("shared", "value")
	intSig := NewSignal("shared", 42)

	if got := GetSignal[string]("shared"); got == nil || got != strSig {
		t.Fatal("expected string signal for shared name")
	}
	if got := GetSignal[int]("shared"); got == nil || got != intSig {
		t.Fatal("expected int signal for shared name")
	}
}

func TestSignalSameNameSameTypePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for duplicate registration of same package/name/type")
		}
	}()

	NewSignal("dup", 1)
	NewSignal("dup", 2)
}

func TestSignalConcurrentUse(t *testing.T) {
	s := NewSignal("concurrent", 0)
	const goroutines = 32
	const iterations = 200

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				value := offset + j
				s.Set(value)
				_ = s.Get()
				s.Update(func(current int) int { return current + 1 })
			}
		}(i)
	}
	wg.Wait()
}

func TestSignalNilableValueDoesNotPanic(t *testing.T) {
	var value any = nil
	if value != nil {
		t.Fatal("unexpected non-nil value")
	}

	// The key should be generated without dereferencing a nil type.
	if got := signalTypeKey("", "", value); got != "" {
		t.Fatalf("expected empty key when both scope and name are blank, got %q", got)
	}

	if got := signalTypeKey("app", "demo", value); got != "app::demo:<nil>" {
		t.Fatalf("expected nil-safe key, got %q", got)
	}

	if got := NewSignal("demo-nil", (*int)(nil)); got == nil {
		t.Fatal("expected signal to be created")
	}
}

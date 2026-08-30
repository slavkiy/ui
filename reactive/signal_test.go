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

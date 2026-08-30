package event

import "testing"

func TestKeyString(t *testing.T) {
	cases := map[Key]string{
		KeyA:          "A",
		KeyZ:          "Z",
		Key0:          "0",
		Key9:          "9",
		KeyF1:         "F1",
		KeyF24:        "F24",
		KeyEnter:      "Enter",
		KeyEscape:     "Escape",
		KeyArrowUp:    "ArrowUp",
		KeyArrowDown:  "ArrowDown",
		KeyArrowLeft:  "ArrowLeft",
		KeyArrowRight: "ArrowRight",
		KeyNumpad0:    "Numpad0",
		KeyNumpad9:    "Numpad9",
		KeyUnknown:    "Unknown",
	}

	for key, want := range cases {
		if got := key.String(); got != want {
			t.Fatalf("Key(%d).String() = %q, want %q", key, got, want)
		}
	}
}

func TestModifierMask(t *testing.T) {
	mask := ModControl | ModShift
	if !mask.Has(ModControl) {
		t.Fatal("expected ModControl to be present")
	}
	if !mask.Has(ModShift) {
		t.Fatal("expected ModShift to be present")
	}
	if mask.Has(ModAlt) {
		t.Fatal("unexpected ModAlt present")
	}
	if (ModCommand | ModOption).Empty() {
		t.Fatal("expected composite modifiers to be non-empty")
	}
}

func TestModifierString(t *testing.T) {
	if got := (ModControl | ModShift).String(); got == "" {
		t.Fatal("expected non-empty modifier string")
	}
	if got := ModNone.String(); got != "None" {
		t.Fatalf("ModNone.String() = %q, want None", got)
	}
}

func TestKeyEventAndKeyStroke(t *testing.T) {
	ke := KeyEvent{Key: KeyA, Modifiers: ModControl | ModShift, Repeat: true}
	if ke.Key != KeyA {
		t.Fatal("expected KeyA")
	}
	if !ke.Modifiers.Has(ModControl) || !ke.Modifiers.Has(ModShift) {
		t.Fatal("expected modifier combination to work")
	}

	ks := KeyStroke{Key: KeyEnter, Modifiers: ModCommand | ModOption}
	if ks.Key != KeyEnter {
		t.Fatal("expected KeyEnter")
	}
	if !ks.Modifiers.Has(ModCommand) {
		t.Fatal("expected ModCommand to be present")
	}
}

func TestUnknownKey(t *testing.T) {
	if got := Key(65535).String(); got != "Unknown" {
		t.Fatalf("unknown key string = %q, want Unknown", got)
	}
}

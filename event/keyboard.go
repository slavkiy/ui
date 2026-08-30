package event

import (
	"strings"
	"time"
)

// Key identifies a physical or logical keyboard key in a platform-neutral way.
// It does not represent text input; text input is provided by TextInput.
type Key uint16

const (
	KeyUnknown Key = iota

	// Letters.
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	// Digits.
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9

	// Control keys.
	KeyEnter
	KeyEscape
	KeyTab
	KeyBackspace
	KeyDelete
	KeyInsert
	KeySpace
	KeyPause
	KeyPrintScreen

	// Navigation.
	KeyHome
	KeyEnd
	KeyPageUp
	KeyPageDown
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight

	// Function keys.
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24

	// Modifier keys.
	KeyShift
	KeyControl
	KeyAlt
	KeySuper
	KeyCommand
	KeyOption
	KeyWindows
	KeyCapsLock
	KeyNumLock
	KeyScrollLock
	KeyFn

	// Numpad.
	KeyNumpad0
	KeyNumpad1
	KeyNumpad2
	KeyNumpad3
	KeyNumpad4
	KeyNumpad5
	KeyNumpad6
	KeyNumpad7
	KeyNumpad8
	KeyNumpad9
	KeyNumpadDecimal
	KeyNumpadAdd
	KeyNumpadSubtract
	KeyNumpadMultiply
	KeyNumpadDivide
	KeyNumpadEnter
	KeyNumpadEqual

	// Punctuation.
	KeyMinus
	KeyEqual
	KeyLeftBracket
	KeyRightBracket
	KeyBackslash
	KeySemicolon
	KeyQuote
	KeyComma
	KeyPeriod
	KeySlash
	KeyBacktick

	// International and IME.
	KeyKana
	KeyKanji
	KeyHiragana
	KeyKatakana
	KeyHangul
	KeyHanja
	KeyConvert
	KeyNonConvert
	KeyMuhenkan
	KeyHenkan
	KeyCompose
	KeyDeadGrave
	KeyDeadAcute
	KeyDeadCircumflex
	KeyDeadTilde
	KeyDeadDiaeresis
	KeyDeadMacron
	KeyDeadBreve
	KeyDeadRing
	KeyDeadCedilla
	KeyDeadOgonek
	KeyDeadCaron
	KeyIMEOff
	KeyIMEOn

	// Editing.
	KeyCut
	KeyCopy
	KeyPaste
	KeyUndo
	KeyRedo
	KeyFind
	KeySelectAll
	KeyHelp
	KeyMenu

	// Media and system keys.
	KeyVolumeUp
	KeyVolumeDown
	KeyVolumeMute
	KeyMediaPlay
	KeyMediaPause
	KeyMediaPlayPause
	KeyMediaStop
	KeyMediaNext
	KeyMediaPrevious
	KeyBrightnessUp
	KeyBrightnessDown
	KeyEject
	KeyPower
	KeySleep
	KeyWake

	// Browser and application keys.
	KeyBrowserBack
	KeyBrowserForward
	KeyBrowserRefresh
	KeyBrowserStop
	KeyBrowserSearch
	KeyBrowserFavorites
	KeyBrowserHome
	KeyMail
	KeyCalculator
	KeyExplorer
)

// Modifiers is a platform-agnostic bitmask for keyboard modifiers.
// It supports combinations via bitwise OR, for example ModControl | ModShift.
type Modifiers uint16

const (
	ModNone Modifiers = 0

	ModShift Modifiers = 1 << iota
	ModControl
	ModAlt
	ModSuper
	ModCommand
	ModOption
	ModWindows
	ModCapsLock
	ModNumLock
	ModScrollLock
	ModFn
)

// ModMeta is a generic alias for the platform's "super/meta" key.
const ModMeta = ModSuper

// ModAltGr is a common alias for AltGr/Option-like combinations on some layouts.
const ModAltGr = ModAlt

// Has reports whether the given modifier bit is present in the mask.
func (m Modifiers) Has(flag Modifiers) bool {
	return m&flag == flag
}

// Empty reports whether the modifier mask contains any set bits.
func (m Modifiers) Empty() bool {
	return m == ModNone
}

// KeyEvent represents a keyboard event with a normalized key and modifier state.
type KeyEvent struct {
	Key       Key
	Modifiers Modifiers
	Repeat    bool
	Timestamp time.Duration
}

// KeyStroke is a normalized key press event without repetition metadata.
type KeyStroke struct {
	Key       Key
	Modifiers Modifiers
}

// KeyDown is emitted when a key is pressed.
type KeyDown struct {
	KeyEvent
}

// KeyUp is emitted when a key is released.
type KeyUp struct {
	KeyEvent
}

// TextInput represents text that was produced by the platform input system.
// This is intentionally separate from Key because keyboard layout, IME, dead keys,
// and Unicode input all make a direct Key-to-text mapping incorrect.
type TextInput struct {
	Text      string
	Modifiers Modifiers
	Timestamp time.Duration
}

// String returns a human-readable name for a normalized key.
func (k Key) String() string {
	switch k {
	case KeyUnknown:
		return "Unknown"
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case Key0:
		return "0"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case KeyEnter:
		return "Enter"
	case KeyEscape:
		return "Escape"
	case KeyTab:
		return "Tab"
	case KeyBackspace:
		return "Backspace"
	case KeyDelete:
		return "Delete"
	case KeyInsert:
		return "Insert"
	case KeySpace:
		return "Space"
	case KeyPause:
		return "Pause"
	case KeyPrintScreen:
		return "PrintScreen"
	case KeyHome:
		return "Home"
	case KeyEnd:
		return "End"
	case KeyPageUp:
		return "PageUp"
	case KeyPageDown:
		return "PageDown"
	case KeyArrowUp:
		return "ArrowUp"
	case KeyArrowDown:
		return "ArrowDown"
	case KeyArrowLeft:
		return "ArrowLeft"
	case KeyArrowRight:
		return "ArrowRight"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyF13:
		return "F13"
	case KeyF14:
		return "F14"
	case KeyF15:
		return "F15"
	case KeyF16:
		return "F16"
	case KeyF17:
		return "F17"
	case KeyF18:
		return "F18"
	case KeyF19:
		return "F19"
	case KeyF20:
		return "F20"
	case KeyF21:
		return "F21"
	case KeyF22:
		return "F22"
	case KeyF23:
		return "F23"
	case KeyF24:
		return "F24"
	case KeyShift:
		return "Shift"
	case KeyControl:
		return "Control"
	case KeyAlt:
		return "Alt"
	case KeySuper:
		return "Super"
	case KeyCommand:
		return "Command"
	case KeyOption:
		return "Option"
	case KeyWindows:
		return "Windows"
	case KeyCapsLock:
		return "CapsLock"
	case KeyNumLock:
		return "NumLock"
	case KeyScrollLock:
		return "ScrollLock"
	case KeyFn:
		return "Fn"
	case KeyNumpad0:
		return "Numpad0"
	case KeyNumpad1:
		return "Numpad1"
	case KeyNumpad2:
		return "Numpad2"
	case KeyNumpad3:
		return "Numpad3"
	case KeyNumpad4:
		return "Numpad4"
	case KeyNumpad5:
		return "Numpad5"
	case KeyNumpad6:
		return "Numpad6"
	case KeyNumpad7:
		return "Numpad7"
	case KeyNumpad8:
		return "Numpad8"
	case KeyNumpad9:
		return "Numpad9"
	case KeyNumpadDecimal:
		return "NumpadDecimal"
	case KeyNumpadAdd:
		return "NumpadAdd"
	case KeyNumpadSubtract:
		return "NumpadSubtract"
	case KeyNumpadMultiply:
		return "NumpadMultiply"
	case KeyNumpadDivide:
		return "NumpadDivide"
	case KeyNumpadEnter:
		return "NumpadEnter"
	case KeyNumpadEqual:
		return "NumpadEqual"
	case KeyMinus:
		return "Minus"
	case KeyEqual:
		return "Equal"
	case KeyLeftBracket:
		return "LeftBracket"
	case KeyRightBracket:
		return "RightBracket"
	case KeyBackslash:
		return "Backslash"
	case KeySemicolon:
		return "Semicolon"
	case KeyQuote:
		return "Quote"
	case KeyComma:
		return "Comma"
	case KeyPeriod:
		return "Period"
	case KeySlash:
		return "Slash"
	case KeyBacktick:
		return "Backtick"
	case KeyKana:
		return "Kana"
	case KeyKanji:
		return "Kanji"
	case KeyHiragana:
		return "Hiragana"
	case KeyKatakana:
		return "Katakana"
	case KeyHangul:
		return "Hangul"
	case KeyHanja:
		return "Hanja"
	case KeyConvert:
		return "Convert"
	case KeyNonConvert:
		return "NonConvert"
	case KeyMuhenkan:
		return "Muhenkan"
	case KeyHenkan:
		return "Henkan"
	case KeyCompose:
		return "Compose"
	case KeyDeadGrave:
		return "DeadGrave"
	case KeyDeadAcute:
		return "DeadAcute"
	case KeyDeadCircumflex:
		return "DeadCircumflex"
	case KeyDeadTilde:
		return "DeadTilde"
	case KeyDeadDiaeresis:
		return "DeadDiaeresis"
	case KeyDeadMacron:
		return "DeadMacron"
	case KeyDeadBreve:
		return "DeadBreve"
	case KeyDeadRing:
		return "DeadRing"
	case KeyDeadCedilla:
		return "DeadCedilla"
	case KeyDeadOgonek:
		return "DeadOgonek"
	case KeyDeadCaron:
		return "DeadCaron"
	case KeyIMEOff:
		return "IMEOff"
	case KeyIMEOn:
		return "IMEOn"
	case KeyCut:
		return "Cut"
	case KeyCopy:
		return "Copy"
	case KeyPaste:
		return "Paste"
	case KeyUndo:
		return "Undo"
	case KeyRedo:
		return "Redo"
	case KeyFind:
		return "Find"
	case KeySelectAll:
		return "SelectAll"
	case KeyHelp:
		return "Help"
	case KeyMenu:
		return "Menu"
	case KeyVolumeUp:
		return "VolumeUp"
	case KeyVolumeDown:
		return "VolumeDown"
	case KeyVolumeMute:
		return "VolumeMute"
	case KeyMediaPlay:
		return "MediaPlay"
	case KeyMediaPause:
		return "MediaPause"
	case KeyMediaPlayPause:
		return "MediaPlayPause"
	case KeyMediaStop:
		return "MediaStop"
	case KeyMediaNext:
		return "MediaNext"
	case KeyMediaPrevious:
		return "MediaPrevious"
	case KeyBrightnessUp:
		return "BrightnessUp"
	case KeyBrightnessDown:
		return "BrightnessDown"
	case KeyEject:
		return "Eject"
	case KeyPower:
		return "Power"
	case KeySleep:
		return "Sleep"
	case KeyWake:
		return "Wake"
	case KeyBrowserBack:
		return "BrowserBack"
	case KeyBrowserForward:
		return "BrowserForward"
	case KeyBrowserRefresh:
		return "BrowserRefresh"
	case KeyBrowserStop:
		return "BrowserStop"
	case KeyBrowserSearch:
		return "BrowserSearch"
	case KeyBrowserFavorites:
		return "BrowserFavorites"
	case KeyBrowserHome:
		return "BrowserHome"
	case KeyMail:
		return "Mail"
	case KeyCalculator:
		return "Calculator"
	case KeyExplorer:
		return "Explorer"
	default:
		return "Unknown"
	}
}

func (k Key) Normalized() Key {
	if k == KeyUnknown {
		return KeyUnknown
	}
	return k
}

func (m Modifiers) String() string {
	if m == ModNone {
		return "None"
	}

	parts := make([]string, 0, 6)
	if m.Has(ModShift) {
		parts = append(parts, "Shift")
	}
	if m.Has(ModControl) {
		parts = append(parts, "Control")
	}
	if m.Has(ModAlt) {
		parts = append(parts, "Alt")
	}
	if m.Has(ModSuper) {
		parts = append(parts, "Super")
	}
	if m.Has(ModCommand) {
		parts = append(parts, "Command")
	}
	if m.Has(ModOption) {
		parts = append(parts, "Option")
	}
	if m.Has(ModWindows) {
		parts = append(parts, "Windows")
	}
	if m.Has(ModCapsLock) {
		parts = append(parts, "CapsLock")
	}
	if m.Has(ModNumLock) {
		parts = append(parts, "NumLock")
	}
	if m.Has(ModScrollLock) {
		parts = append(parts, "ScrollLock")
	}
	if m.Has(ModFn) {
		parts = append(parts, "Fn")
	}
	return strings.Join(parts, "|")
}

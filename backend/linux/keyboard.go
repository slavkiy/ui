package linux

import "github.com/slavkiy/ui/event"

// KeyCodeToEvent converts Linux evdev key codes to a normalized ui event.Key.
func KeyCodeToEvent(code uint16) event.Key {
	switch code {
	case 1:
		return event.KeyEscape
	case 2:
		return event.Key1
	case 3:
		return event.Key2
	case 4:
		return event.Key3
	case 5:
		return event.Key4
	case 6:
		return event.Key5
	case 7:
		return event.Key6
	case 8:
		return event.Key7
	case 9:
		return event.Key8
	case 10:
		return event.Key9
	case 11:
		return event.Key0
	case 12:
		return event.KeyMinus
	case 13:
		return event.KeyEqual
	case 14:
		return event.KeyBackspace
	case 15:
		return event.KeyTab
	case 16:
		return event.KeyQ
	case 17:
		return event.KeyW
	case 18:
		return event.KeyE
	case 19:
		return event.KeyR
	case 20:
		return event.KeyT
	case 21:
		return event.KeyY
	case 22:
		return event.KeyU
	case 23:
		return event.KeyI
	case 24:
		return event.KeyO
	case 25:
		return event.KeyP
	case 26:
		return event.KeyLeftBracket
	case 27:
		return event.KeyRightBracket
	case 28:
		return event.KeyEnter
	case 29:
		return event.KeyControl
	case 30:
		return event.KeyA
	case 31:
		return event.KeyS
	case 32:
		return event.KeyD
	case 33:
		return event.KeyF
	case 34:
		return event.KeyG
	case 35:
		return event.KeyH
	case 36:
		return event.KeyJ
	case 37:
		return event.KeyK
	case 38:
		return event.KeyL
	case 39:
		return event.KeySemicolon
	case 40:
		return event.KeyQuote
	case 41:
		return event.KeyBacktick
	case 42:
		return event.KeyShift
	case 43:
		return event.KeyBackslash
	case 44:
		return event.KeyZ
	case 45:
		return event.KeyX
	case 46:
		return event.KeyC
	case 47:
		return event.KeyV
	case 48:
		return event.KeyB
	case 49:
		return event.KeyN
	case 50:
		return event.KeyM
	case 51:
		return event.KeyComma
	case 52:
		return event.KeyPeriod
	case 53:
		return event.KeySlash
	case 54:
		return event.KeyShift
	case 55:
		return event.KeyNumpadMultiply
	case 56:
		return event.KeyAlt
	case 57:
		return event.KeySpace
	case 58:
		return event.KeyCapsLock
	case 59:
		return event.KeyF1
	case 60:
		return event.KeyF2
	case 61:
		return event.KeyF3
	case 62:
		return event.KeyF4
	case 63:
		return event.KeyF5
	case 64:
		return event.KeyF6
	case 65:
		return event.KeyF7
	case 66:
		return event.KeyF8
	case 67:
		return event.KeyF9
	case 68:
		return event.KeyF10
	case 69:
		return event.KeyNumLock
	case 70:
		return event.KeyScrollLock
	case 71:
		return event.KeyNumpad7
	case 72:
		return event.KeyNumpad8
	case 73:
		return event.KeyNumpad9
	case 74:
		return event.KeyNumpadSubtract
	case 75:
		return event.KeyNumpad4
	case 76:
		return event.KeyNumpad5
	case 77:
		return event.KeyNumpad6
	case 78:
		return event.KeyNumpadAdd
	case 79:
		return event.KeyNumpad1
	case 80:
		return event.KeyNumpad2
	case 81:
		return event.KeyNumpad3
	case 82:
		return event.KeyNumpad0
	case 83:
		return event.KeyNumpadDecimal
	case 86:
		return event.KeyBackslash
	case 87:
		return event.KeyF11
	case 88:
		return event.KeyF12
	case 89:
		return event.KeyNumpadEqual
	case 90:
		return event.KeyNumpadEnter
	case 91:
		return event.KeySuper
	case 92:
		return event.KeySuper
	case 94:
		return event.KeyControl
	case 95:
		return event.KeyF11
	case 96:
		return event.KeyF12
	case 103:
		return event.KeyArrowUp
	case 105:
		return event.KeyArrowLeft
	case 106:
		return event.KeyArrowRight
	case 108:
		return event.KeyArrowDown
	case 110:
		return event.KeyHome
	case 111:
		return event.KeyPageUp
	case 112:
		return event.KeyPageDown
	case 113:
		return event.KeyEnd
	default:
		return event.KeyUnknown
	}
}

// ModifiersFromNative converts Linux evdev flags to ui event.Modifiers.
func ModifiersFromNative(bits uint16) event.Modifiers {
	var m event.Modifiers
	if bits&0x01 != 0 {
		m |= event.ModShift
	}
	if bits&0x02 != 0 {
		m |= event.ModControl
	}
	if bits&0x04 != 0 {
		m |= event.ModAlt
	}
	if bits&0x08 != 0 {
		m |= event.ModSuper
	}
	return m
}

// EventFromNative converts a native Linux key event to a normalized event.KeyEvent.
func EventFromNative(code uint16, bits uint16, repeat bool) event.KeyEvent {
	return event.KeyEvent{
		Key:       KeyCodeToEvent(code),
		Modifiers: ModifiersFromNative(bits),
		Repeat:    repeat,
	}
}

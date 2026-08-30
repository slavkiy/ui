package macos

import (
	"github.com/slavkiy/ui/event"
)

// KeyCodeToEvent converts a macOS virtual key code into a normalized ui event.Key.
// This is intentionally a table-driven conversion instead of plain integer cast.
func KeyCodeToEvent(code uint16) event.Key {
	switch code {
	case 0x00:
		return event.KeyA
	case 0x01:
		return event.KeyS
	case 0x02:
		return event.KeyD
	case 0x03:
		return event.KeyF
	case 0x04:
		return event.KeyH
	case 0x05:
		return event.KeyG
	case 0x06:
		return event.KeyZ
	case 0x07:
		return event.KeyX
	case 0x08:
		return event.KeyC
	case 0x09:
		return event.KeyV
	case 0x0B:
		return event.KeyB
	case 0x0C:
		return event.KeyQ
	case 0x0D:
		return event.KeyW
	case 0x0E:
		return event.KeyE
	case 0x0F:
		return event.KeyR
	case 0x10:
		return event.KeyY
	case 0x11:
		return event.KeyT
	case 0x12:
		return event.Key1
	case 0x13:
		return event.Key2
	case 0x14:
		return event.Key3
	case 0x15:
		return event.Key4
	case 0x16:
		return event.Key6
	case 0x17:
		return event.Key5
	case 0x18:
		return event.KeyEqual
	case 0x19:
		return event.Key9
	case 0x1A:
		return event.Key7
	case 0x1B:
		return event.KeyMinus
	case 0x1C:
		return event.Key8
	case 0x1D:
		return event.Key0
	case 0x1E:
		return event.KeyRightBracket
	case 0x1F:
		return event.KeyO
	case 0x20:
		return event.KeyU
	case 0x21:
		return event.KeyLeftBracket
	case 0x22:
		return event.KeyI
	case 0x23:
		return event.KeyP
	case 0x24:
		return event.KeyEnter
	case 0x25:
		return event.KeyArrowLeft
	case 0x26:
		return event.KeyArrowRight
	case 0x27:
		return event.KeyArrowDown
	case 0x28:
		return event.KeyArrowUp
	case 0x29:
		return event.KeyEscape
	case 0x2A:
		return event.KeyDelete
	case 0x2B:
		return event.KeyTab
	case 0x2C:
		return event.KeySpace
	case 0x2D:
		return event.KeyMinus
	case 0x2E:
		return event.KeyBackspace
	case 0x2F:
		return event.KeyEnter
	case 0x30:
		return event.KeyNumpad0
	case 0x31:
		return event.KeyNumpad1
	case 0x32:
		return event.KeyNumpad2
	case 0x33:
		return event.KeyNumpad3
	case 0x34:
		return event.KeyNumpad4
	case 0x35:
		return event.KeyNumpad5
	case 0x36:
		return event.KeyNumpad6
	case 0x37:
		return event.KeyNumpad7
	case 0x38:
		return event.KeyNumpad8
	case 0x39:
		return event.KeyNumpad9
	case 0x3A:
		return event.KeyCapsLock
	case 0x3B:
		return event.KeyF1
	case 0x3C:
		return event.KeyF2
	case 0x3D:
		return event.KeyF3
	case 0x3E:
		return event.KeyF4
	case 0x3F:
		return event.KeyF5
	case 0x40:
		return event.KeyF6
	case 0x41:
		return event.KeyF7
	case 0x42:
		return event.KeyF8
	case 0x43:
		return event.KeyF9
	case 0x44:
		return event.KeyF10
	case 0x45:
		return event.KeyF11
	case 0x46:
		return event.KeyF12
	default:
		return event.KeyUnknown
	}
}

// ModifiersFromNative converts native macOS modifier flags to ui event.Modifiers.
func ModifiersFromNative(flags uint) event.Modifiers {
	var m event.Modifiers
	if flags&0x20000 != 0 {
		m |= event.ModControl
	}
	if flags&0x40000 != 0 {
		m |= event.ModShift
	}
	if flags&0x80000 != 0 {
		m |= event.ModOption
	}
	if flags&0x100000 != 0 {
		m |= event.ModCommand
	}
	if flags&0x200000 != 0 {
		m |= event.ModControl
	}
	return m
}

// EventFromNative converts a native macOS key event to a normalized event.KeyEvent.
func EventFromNative(code uint16, flags uint, repeat bool) event.KeyEvent {
	return event.KeyEvent{
		Key:       KeyCodeToEvent(code),
		Modifiers: ModifiersFromNative(flags),
		Repeat:    repeat,
	}
}

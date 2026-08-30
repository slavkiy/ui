package windows

import "github.com/slavkiy/ui/event"

// VKeyToEvent converts Windows VK codes to a normalized ui event.Key.
func VKeyToEvent(vk uint16) event.Key {
	switch vk {
	case 0x08:
		return event.KeyBackspace
	case 0x09:
		return event.KeyTab
	case 0x0D:
		return event.KeyEnter
	case 0x10:
		return event.KeyShift
	case 0x11:
		return event.KeyControl
	case 0x12:
		return event.KeyAlt
	case 0x13:
		return event.KeyPause
	case 0x14:
		return event.KeyCapsLock
	case 0x1B:
		return event.KeyEscape
	case 0x20:
		return event.KeySpace
	case 0x21:
		return event.KeyPageUp
	case 0x22:
		return event.KeyPageDown
	case 0x23:
		return event.KeyEnd
	case 0x24:
		return event.KeyHome
	case 0x25:
		return event.KeyArrowLeft
	case 0x26:
		return event.KeyArrowUp
	case 0x27:
		return event.KeyArrowRight
	case 0x28:
		return event.KeyArrowDown
	case 0x2C:
		return event.KeyPrintScreen
	case 0x2D:
		return event.KeyInsert
	case 0x2E:
		return event.KeyDelete
	case 0x30:
		return event.Key0
	case 0x31:
		return event.Key1
	case 0x32:
		return event.Key2
	case 0x33:
		return event.Key3
	case 0x34:
		return event.Key4
	case 0x35:
		return event.Key5
	case 0x36:
		return event.Key6
	case 0x37:
		return event.Key7
	case 0x38:
		return event.Key8
	case 0x39:
		return event.Key9
	case 0x41:
		return event.KeyA
	case 0x42:
		return event.KeyB
	case 0x43:
		return event.KeyC
	case 0x44:
		return event.KeyD
	case 0x45:
		return event.KeyE
	case 0x46:
		return event.KeyF
	case 0x47:
		return event.KeyG
	case 0x48:
		return event.KeyH
	case 0x49:
		return event.KeyI
	case 0x4A:
		return event.KeyJ
	case 0x4B:
		return event.KeyK
	case 0x4C:
		return event.KeyL
	case 0x4D:
		return event.KeyM
	case 0x4E:
		return event.KeyN
	case 0x4F:
		return event.KeyO
	case 0x50:
		return event.KeyP
	case 0x51:
		return event.KeyQ
	case 0x52:
		return event.KeyR
	case 0x53:
		return event.KeyS
	case 0x54:
		return event.KeyT
	case 0x55:
		return event.KeyU
	case 0x56:
		return event.KeyV
	case 0x57:
		return event.KeyW
	case 0x58:
		return event.KeyX
	case 0x59:
		return event.KeyY
	case 0x5A:
		return event.KeyZ
	case 0x5B:
		return event.KeyWindows
	case 0x5C:
		return event.KeyWindows
	case 0x60:
		return event.KeyNumpad0
	case 0x61:
		return event.KeyNumpad1
	case 0x62:
		return event.KeyNumpad2
	case 0x63:
		return event.KeyNumpad3
	case 0x64:
		return event.KeyNumpad4
	case 0x65:
		return event.KeyNumpad5
	case 0x66:
		return event.KeyNumpad6
	case 0x67:
		return event.KeyNumpad7
	case 0x68:
		return event.KeyNumpad8
	case 0x69:
		return event.KeyNumpad9
	case 0x6A:
		return event.KeyNumpadMultiply
	case 0x6B:
		return event.KeyNumpadAdd
	case 0x6C:
		return event.KeyNumpadDecimal
	case 0x6D:
		return event.KeyNumpadSubtract
	case 0x6E:
		return event.KeyNumpadDecimal
	case 0x6F:
		return event.KeyNumpadDivide
	case 0x70:
		return event.KeyF1
	case 0x71:
		return event.KeyF2
	case 0x72:
		return event.KeyF3
	case 0x73:
		return event.KeyF4
	case 0x74:
		return event.KeyF5
	case 0x75:
		return event.KeyF6
	case 0x76:
		return event.KeyF7
	case 0x77:
		return event.KeyF8
	case 0x78:
		return event.KeyF9
	case 0x79:
		return event.KeyF10
	case 0x7A:
		return event.KeyF11
	case 0x7B:
		return event.KeyF12
	case 0x7C:
		return event.KeyF13
	case 0x7D:
		return event.KeyF14
	case 0x7E:
		return event.KeyF15
	case 0x7F:
		return event.KeyF16
	case 0x80:
		return event.KeyF17
	case 0x81:
		return event.KeyF18
	case 0x82:
		return event.KeyF19
	case 0x83:
		return event.KeyF20
	case 0x84:
		return event.KeyF21
	case 0x85:
		return event.KeyF22
	case 0x86:
		return event.KeyF23
	case 0x87:
		return event.KeyF24
	case 0x90:
		return event.KeyNumLock
	case 0x91:
		return event.KeyScrollLock
	case 0xA0:
		return event.KeyShift
	case 0xA1:
		return event.KeyShift
	case 0xA2:
		return event.KeyControl
	case 0xA3:
		return event.KeyControl
	case 0xA4:
		return event.KeyAlt
	case 0xA5:
		return event.KeyAlt
	default:
		return event.KeyUnknown
	}
}

// ModifiersFromNative converts native Windows modifier flags to ui event.Modifiers.
func ModifiersFromNative(flags uint32) event.Modifiers {
	var m event.Modifiers
	if flags&0x0001 != 0 {
		m |= event.ModShift
	}
	if flags&0x0002 != 0 {
		m |= event.ModControl
	}
	if flags&0x0004 != 0 {
		m |= event.ModAlt
	}
	if flags&0x0008 != 0 {
		m |= event.ModWindows
	}
	return m
}

// EventFromNative converts a native Windows key event to a normalized event.KeyEvent.
func EventFromNative(vk uint16, flags uint32, repeat bool) event.KeyEvent {
	return event.KeyEvent{
		Key:       VKeyToEvent(vk),
		Modifiers: ModifiersFromNative(flags),
		Repeat:    repeat,
	}
}

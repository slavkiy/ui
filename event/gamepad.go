package event

// GamepadButton identifies a button on a gamepad.
type GamepadButton uint8

const (
	GamepadButtonUnknown GamepadButton = iota
	GamepadButtonA
	GamepadButtonB
	GamepadButtonX
	GamepadButtonY
	GamepadButtonLeftTrigger
	GamepadButtonRightTrigger
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonBack
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonDPadUp
	GamepadButtonDPadDown
	GamepadButtonDPadLeft
	GamepadButtonDPadRight
	GamepadButtonGuide
)

// GamepadAxis identifies a gamepad analog axis.
type GamepadAxis uint8

const (
	GamepadAxisUnknown GamepadAxis = iota
	GamepadAxisLeftX
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
)

// GamepadState is a normalized snapshot of a connected controller.
type GamepadState struct {
	ID        uint64
	Name      string
	Connected bool
	Buttons   map[GamepadButton]bool
	Axes      map[GamepadAxis]float32
}

// GamepadConnected is emitted when a controller is connected.
type GamepadConnected struct {
	Event
	Gamepad GamepadState
}

// GamepadDisconnected is emitted when a controller is disconnected.
type GamepadDisconnected struct {
	Event
	Gamepad GamepadState
}

// GamepadButtonDown is emitted when a button is pressed.
type GamepadButtonDown struct {
	Event
	Gamepad GamepadState
	Button  GamepadButton
}

// GamepadButtonUp is emitted when a button is released.
type GamepadButtonUp struct {
	Event
	Gamepad GamepadState
	Button  GamepadButton
}

// GamepadAxis is emitted with analog stick or trigger changes.
type GamepadAxisEvent struct {
	Event
	Gamepad GamepadState
	Axis    GamepadAxis
	Value   float32
}

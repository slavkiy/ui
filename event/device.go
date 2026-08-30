package event

// DeviceKind identifies the origin class of an input device.
type DeviceKind uint8

const (
	DeviceKindUnknown DeviceKind = iota
	DeviceKindKeyboard
	DeviceKindPointer
	DeviceKindTouch
	DeviceKindPen
	DeviceKindGamepad
	DeviceKindSensor
	DeviceKindAudio
	DeviceKindDisplay
	DeviceKindNetwork
)

// DeviceInfo describes a connected or active device.
type DeviceInfo struct {
	ID           uint64
	Name         string
	Kind         DeviceKind
	VendorID     uint32
	ProductID    uint32
	Version      uint32
	Connected    bool
	IsPrimary    bool
	Capabilities []string
}

// DeviceEvent is emitted for device lifecycle changes.
type DeviceEvent struct {
	Event
	Device DeviceInfo
}

// DeviceConnected is emitted when a new device becomes available.
type DeviceConnected struct{ DeviceEvent }

// DeviceDisconnected is emitted when a device is removed.
type DeviceDisconnected struct{ DeviceEvent }

// DeviceChanged is emitted when device capabilities change.
type DeviceChanged struct{ DeviceEvent }

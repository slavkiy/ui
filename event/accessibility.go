package event

// AccessibilityRole describes the semantic role of a UI element.
type AccessibilityRole string

const (
	AccessibilityRoleUnknown AccessibilityRole = "unknown"
	AccessibilityRoleButton  AccessibilityRole = "button"
	AccessibilityRoleText    AccessibilityRole = "text"
	AccessibilityRoleImage   AccessibilityRole = "image"
	AccessibilityRoleInput   AccessibilityRole = "input"
)

// AccessibilityEvent is emitted for assistive-tech metadata or state changes.
type AccessibilityEvent struct {
	Event
	TargetID uint64
	Role     AccessibilityRole
	Label    string
	Value    string
	Enabled  bool
}

// AccessibilityFocus is emitted when a screen reader focus changes.
type AccessibilityFocus struct{ AccessibilityEvent }

// AccessibilityActivated is emitted when an accessible element is activated.
type AccessibilityActivated struct{ AccessibilityEvent }

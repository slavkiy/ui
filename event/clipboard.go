package event

// ClipboardEvent describes content transfer between app and system clipboard.
type ClipboardEvent struct {
	Event
	Text      string
	MimeTypes []string
}

// ClipboardChanged indicates clipboard content changed.
type ClipboardChanged struct{ ClipboardEvent }

// Copy indicates a copy operation was requested.
type Copy struct{ ClipboardEvent }

// Cut indicates a cut operation was requested.
type Cut struct{ ClipboardEvent }

// Paste indicates a paste operation was requested.
type Paste struct{ ClipboardEvent }

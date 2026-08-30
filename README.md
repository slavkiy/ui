# ui

`ui` is a cross-platform UI toolkit with a platform-neutral event model and a backend abstraction.

## Architecture

The project is split into a few simple layers:

- `event` — shared, platform-independent event model
- `reactive` — state and signal updates
- `render` — drawing primitives and canvas interface
- `backend` — platform adapter layer
- `widget` — UI composition layer

The main idea is simple:

- native OS code belongs in platform backends
- backend adapters normalize native input into shared `event.Event` values
- the runtime dispatches those values into shared signals
- client code subscribes to client-only signals and reads updates

## Backend contract

The common backend contract is defined in `backend/backend.go`.

```go
package backend

type Backend interface {
	Name() string
	Platform() Platform
	Canvas() render.Canvas
	Emit(evt *event.Event)
	Close() error
}
```

This makes it easy to accept any backend in application code without coupling the host to a single platform.

## Available platforms

Supported platform identifiers:

- `macos`
- `windows`
- `linux`
- `android`
- `web`
- `ios`

Factory helpers:

```go
b, err := backend.NewPlatform(backend.PlatformMacOS)
if err != nil {
	panic(err)
}

_ = b
```

You can also instantiate concrete adapters directly:

```go
b := backend.NewMacOS()
_ = b
```

## Event flow

The general event flow is:

1. platform native event arrives
2. platform backend converts it to a normalized shared `event.Event`
3. backend pushes it to the runtime via `Emit`
4. the runtime updates current event signals
5. client code listens via `ClientSignal.Subscribe` or `SubscribeChan`

## Client-only public event API

Public event consumers work with client-side signals, not writeable server-side state.

```go
client := event.KeyDownSignal
ch, sub := client.SubscribeChan(16)
defer sub.Unsubscribe()

select {
case evt := <-ch:
	println(evt.Type)
default:
}
```

In the same style, other families are available:

- `CurrentEventClient`
- `KeyDownSignal`
- `KeyUpSignal`
- `PointerMoveSignal`
- `PointerDownSignal`
- `PointerUpSignal`
- `TouchMoveSignal`
- `TouchStartSignal`
- `TouchEndSignal`
- `ScrollSignal`
- `WindowSignal`
- `FocusSignal`
- `TextInputSignal`

## Canvas abstraction

Rendering is abstracted via `render.Canvas`, so all platforms can plug into the same drawing API.

The shared canvas contract includes:

- geometry drawing
- filled/stroked primitives
- gradient support
- images and text
- transforms, clipping, and compositing

A platform backend may provide a native canvas implementation, but the API seen by the app remains the same.

## Platform backends

The project currently contains the common adapter layer and platform entry points:

- `backend/macos`
- `backend/windows`
- `backend/linux`

Additional concrete platform implementations can be added in the same style for:

- Android
- Web
- iOS

The backend abstraction is intentionally easy to extend.

## Build per platform

The project is structured so each backend can be selected independently by platform name.

Example:

```bash
# build all packages
 go test ./...

# platform-specific build is a matter of selecting the target backend in host code
# for example:
# - macos backend for macOS
# - windows backend for Windows
# - linux backend for Linux
# - android backend for Android
# - web backend for browser
# - ios backend for iOS
```

In practice, the application chooses the target backend at startup and injects the corresponding canvas.

## Example

```go
package main

import (
	"github.com/slavkiy/ui/backend"
	"github.com/slavkiy/ui/event"
)

func main() {
	b, err := backend.NewPlatform(backend.PlatformMacOS)
	if err != nil {
		panic(err)
	}

	ch, sub := event.KeyDownSignal.SubscribeChan(32)
	defer sub.Unsubscribe()

	_ = b
	_ = ch
}
```

## Notes

- This is a framework-level abstraction, so the real native window and canvas implementation should be bound in the host application.
- The library exposes the portable model; the OS-specific implementation lives in each backend package.
- The backend abstraction is intentionally low-level and easy to adapt to new platforms without changing the app code.

package render

import (
	"fmt"
	"sync"
)

// CanvasCommand is a recorded command for a platform-neutral in-memory canvas.
// It is useful as a concrete backend for tests, headless rendering, and validation
// of command ordering in the core rendering pipeline.
type CanvasCommand struct {
	Kind string
	Data any
}

// MemoryCanvas is a concrete canvas implementation that records drawing commands
// in a thread-safe list. It is fully platform-agnostic and useful as a default
// rendering backend when no native GPU/backend is bound.
type MemoryCanvas struct {
	mu       sync.Mutex
	commands []CanvasCommand
	stack    []canvasState
	state    canvasState
}

type canvasState struct {
	transform Matrix
	opacity   float32
	blend     BlendMode
	shadow    Shadow
	blur      float32
}

func NewMemoryCanvas() *MemoryCanvas {
	c := &MemoryCanvas{state: canvasState{transform: IdentityMatrix(), opacity: 1, blend: BlendModeSrcOver}}
	c.stack = append(c.stack, c.state)
	return c
}

func (c *MemoryCanvas) record(kind string, data any) {
	if c == nil {
		return
	}
	c.mu.Lock()
	c.commands = append(c.commands, CanvasCommand{Kind: kind, Data: data})
	c.mu.Unlock()
}

func (c *MemoryCanvas) Commands() []CanvasCommand {
	if c == nil {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]CanvasCommand, len(c.commands))
	copy(out, c.commands)
	return out
}

func (c *MemoryCanvas) Reset() {
	if c == nil {
		return
	}
	c.mu.Lock()
	c.commands = nil
	c.stack = nil
	c.state = canvasState{transform: IdentityMatrix(), opacity: 1, blend: BlendModeSrcOver}
	c.stack = append(c.stack, c.state)
	c.mu.Unlock()
}

func (c *MemoryCanvas) Save() {
	if c == nil {
		return
	}
	c.mu.Lock()
	c.stack = append(c.stack, c.state)
	c.mu.Unlock()
	c.record("save", nil)
}

func (c *MemoryCanvas) Restore() {
	if c == nil {
		return
	}
	c.mu.Lock()
	if len(c.stack) > 0 {
		c.state = c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
	}
	c.mu.Unlock()
	c.record("restore", nil)
}

func (c *MemoryCanvas) FillRect(rect Rect, color Color) {
	c.record("fill_rect", struct {
		Rect
		Color
	}{Rect: rect, Color: color})
}
func (c *MemoryCanvas) StrokeRect(rect Rect, color Color, width float32) {
	c.record("stroke_rect", struct {
		Rect
		Color
		Width float32
	}{Rect: rect, Color: color, Width: width})
}
func (c *MemoryCanvas) FillRoundedRect(rect Rect, radius float32, color Color) {
	c.record("fill_rounded_rect", struct {
		Rect
		Radius float32
		Color
	}{Rect: rect, Radius: radius, Color: color})
}
func (c *MemoryCanvas) StrokeRoundedRect(rect Rect, radius float32, color Color, width float32) {
	c.record("stroke_rounded_rect", struct {
		Rect
		Radius float32
		Color
		Width float32
	}{Rect: rect, Radius: radius, Color: color, Width: width})
}
func (c *MemoryCanvas) FillCircle(center Point, radius float32, color Color) {
	c.record("fill_circle", struct {
		Point
		Radius float32
		Color
	}{Point: center, Radius: radius, Color: color})
}
func (c *MemoryCanvas) StrokeCircle(center Point, radius float32, color Color, width float32) {
	c.record("stroke_circle", struct {
		Point
		Radius float32
		Color
		Width float32
	}{Point: center, Radius: radius, Color: color, Width: width})
}
func (c *MemoryCanvas) FillEllipse(rect Rect, color Color) {
	c.record("fill_ellipse", struct {
		Rect
		Color
	}{Rect: rect, Color: color})
}
func (c *MemoryCanvas) StrokeEllipse(rect Rect, color Color, width float32) {
	c.record("stroke_ellipse", struct {
		Rect
		Color
		Width float32
	}{Rect: rect, Color: color, Width: width})
}
func (c *MemoryCanvas) DrawLine(from, to Point, color Color, width float32) {
	c.record("draw_line", struct {
		From, To Point
		Color    Color
		Width    float32
	}{From: from, To: to, Color: color, Width: width})
}
func (c *MemoryCanvas) FillPath(path Path, color Color) {
	c.record("fill_path", struct {
		Path  Path
		Color Color
	}{Path: path, Color: color})
}
func (c *MemoryCanvas) StrokePath(path Path, color Color, width float32) {
	c.record("stroke_path", struct {
		Path  Path
		Color Color
		Width float32
	}{Path: path, Color: color, Width: width})
}
func (c *MemoryCanvas) DrawText(text string, position Point, style TextStyle) {
	c.record("draw_text", struct {
		Text     string
		Position Point
		Style    TextStyle
	}{Text: text, Position: position, Style: style})
}
func (c *MemoryCanvas) DrawImage(image *Image, rect Rect) {
	c.record("draw_image", struct {
		Image *Image
		Rect  Rect
	}{Image: image, Rect: rect})
}
func (c *MemoryCanvas) DrawImageRegion(image *Image, source, destination Rect) {
	c.record("draw_image_region", struct {
		Image               *Image
		Source, Destination Rect
	}{Image: image, Source: source, Destination: destination})
}
func (c *MemoryCanvas) FillLinearGradient(rect Rect, gradient LinearGradient) {
	c.record("fill_linear_gradient", struct {
		Rect     Rect
		Gradient LinearGradient
	}{Rect: rect, Gradient: gradient})
}
func (c *MemoryCanvas) FillRadialGradient(rect Rect, gradient RadialGradient) {
	c.record("fill_radial_gradient", struct {
		Rect     Rect
		Gradient RadialGradient
	}{Rect: rect, Gradient: gradient})
}

func (c *MemoryCanvas) Translate(x, y float32) {
	c.mu.Lock()
	c.state.transform = c.state.transform.Multiply(TranslationMatrix(x, y))
	c.mu.Unlock()
	c.record("translate", struct{ X, Y float32 }{X: x, Y: y})
}

func (c *MemoryCanvas) Scale(x, y float32) {
	c.mu.Lock()
	c.state.transform = c.state.transform.Multiply(ScaleMatrix(x, y))
	c.mu.Unlock()
	c.record("scale", struct{ X, Y float32 }{X: x, Y: y})
}

func (c *MemoryCanvas) Rotate(angle float32) {
	c.mu.Lock()
	c.state.transform = c.state.transform.Multiply(RotationMatrix(angle))
	c.mu.Unlock()
	c.record("rotate", struct{ Angle float32 }{Angle: angle})
}

func (c *MemoryCanvas) Transform(matrix Matrix) {
	c.mu.Lock()
	c.state.transform = c.state.transform.Multiply(matrix)
	c.mu.Unlock()
	c.record("transform", matrix)
}

func (c *MemoryCanvas) ClipRect(rect Rect) { c.record("clip_rect", rect) }
func (c *MemoryCanvas) ClipPath(path Path) { c.record("clip_path", path) }
func (c *MemoryCanvas) SetOpacity(opacity float32) {
	c.mu.Lock()
	c.state.opacity = opacity
	c.mu.Unlock()
	c.record("set_opacity", opacity)
}
func (c *MemoryCanvas) SetBlendMode(mode BlendMode) {
	c.mu.Lock()
	c.state.blend = mode
	c.mu.Unlock()
	c.record("set_blend_mode", mode)
}
func (c *MemoryCanvas) SetShadow(shadow Shadow) {
	c.mu.Lock()
	c.state.shadow = shadow
	c.mu.Unlock()
	c.record("set_shadow", shadow)
}
func (c *MemoryCanvas) SetBlur(radius float32) { c.record("set_blur", radius) }

func (c *MemoryCanvas) String() string {
	if c == nil {
		return "<nil>"
	}
	commands := c.Commands()
	if len(commands) == 0 {
		return "MemoryCanvas{}"
	}
	return fmt.Sprintf("MemoryCanvas{ops=%d}", len(commands))
}

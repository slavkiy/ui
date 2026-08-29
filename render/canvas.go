package render

type Canvas interface {
	// Geometry
	FillRect(rect Rect, color Color)
	StrokeRect(rect Rect, color Color, width float32)

	FillRoundedRect(rect Rect, radius float32, color Color)
	StrokeRoundedRect(rect Rect, radius float32, color Color, width float32)

	FillCircle(center Point, radius float32, color Color)
	StrokeCircle(center Point, radius float32, color Color, width float32)

	FillEllipse(rect Rect, color Color)
	StrokeEllipse(rect Rect, color Color, width float32)

	// Lines
	DrawLine(from, to Point, color Color, width float32)

	// Paths
	FillPath(path Path, color Color)
	StrokePath(path Path, color Color, width float32)

	// Text
	DrawText(text string, position Point, style TextStyle)

	// Images
	DrawImage(image *Image, rect Rect)
	DrawImageRegion(image *Image, source, destination Rect)

	// Gradients
	FillLinearGradient(rect Rect, gradient LinearGradient)
	FillRadialGradient(rect Rect, gradient RadialGradient)

	// State
	Save()
	Restore()

	// Transform
	Translate(x, y float32)
	Scale(x, y float32)
	Rotate(angle float32)
	Transform(matrix Matrix)

	// Clipping
	ClipRect(rect Rect)
	ClipPath(path Path)

	// Compositing
	SetOpacity(opacity float32)
	SetBlendMode(mode BlendMode)

	// Effects
	SetShadow(shadow Shadow)
	SetBlur(radius float32)
}

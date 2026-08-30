package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Point struct {
	X float32
	Y float32
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p Point) Scale(s float32) Point {
	return Point{X: p.X * s, Y: p.Y * s}
}

func (p Point) Length() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

func (p Point) Distance(other Point) float32 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (p Point) Normalize() Point {
	len := p.Length()
	if len == 0 {
		return Point{}
	}
	return Point{X: p.X / len, Y: p.Y / len}
}

func (r Rect) Center() Point {
	return Point{X: r.X + r.Width/2, Y: r.Y + r.Height/2}
}

func (r Rect) Contains(p Point) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}

func (r Rect) Intersects(other Rect) bool {
	return r.X < other.X+other.Width && r.X+r.Width > other.X && r.Y < other.Y+other.Height && r.Y+r.Height > other.Y
}

func (r Rect) Inset(x, y float32) Rect {
	return Rect{X: r.X + x, Y: r.Y + y, Width: r.Width - 2*x, Height: r.Height - 2*y}
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Point3 = Vec3

type Matrix struct {
	M11 float32
	M12 float32
	M13 float32
	M21 float32
	M22 float32
	M23 float32
	M31 float32
	M32 float32
	M33 float32
}

func IdentityMatrix() Matrix {
	return Matrix{
		M11: 1, M12: 0, M13: 0,
		M21: 0, M22: 1, M23: 0,
		M31: 0, M32: 0, M33: 1,
	}
}

func TranslationMatrix(x, y float32) Matrix {
	return Matrix{
		M11: 1, M12: 0, M13: x,
		M21: 0, M22: 1, M23: y,
		M31: 0, M32: 0, M33: 1,
	}
}

func ScaleMatrix(x, y float32) Matrix {
	return Matrix{
		M11: x, M12: 0, M13: 0,
		M21: 0, M22: y, M23: 0,
		M31: 0, M32: 0, M33: 1,
	}
}

func RotationMatrix(angle float32) Matrix {
	s, c := math.Sincos(float64(angle))
	return Matrix{
		M11: float32(c), M12: float32(-s), M13: 0,
		M21: float32(s), M22: float32(c), M23: 0,
		M31: 0, M32: 0, M33: 1,
	}
}

func RotationXMatrix(angle float32) Matrix {
	s, c := math.Sincos(float64(angle))
	return Matrix{
		M11: 1, M12: 0, M13: 0,
		M21: 0, M22: float32(c), M23: float32(-s),
		M31: 0, M32: float32(s), M33: float32(c),
	}
}

func RotationYMatrix(angle float32) Matrix {
	s, c := math.Sincos(float64(angle))
	return Matrix{
		M11: float32(c), M12: 0, M13: float32(s),
		M21: 0, M22: 1, M23: 0,
		M31: float32(-s), M32: 0, M33: float32(c),
	}
}

func RotationZMatrix(angle float32) Matrix {
	s, c := math.Sincos(float64(angle))
	return Matrix{
		M11: float32(c), M12: float32(-s), M13: 0,
		M21: float32(s), M22: float32(c), M23: 0,
		M31: 0, M32: 0, M33: 1,
	}
}

func (m Matrix) Multiply(other Matrix) Matrix {
	return Matrix{
		M11: m.M11*other.M11 + m.M12*other.M21 + m.M13*other.M31,
		M12: m.M11*other.M12 + m.M12*other.M22 + m.M13*other.M32,
		M13: m.M11*other.M13 + m.M12*other.M23 + m.M13*other.M33,
		M21: m.M21*other.M11 + m.M22*other.M21 + m.M23*other.M31,
		M22: m.M21*other.M12 + m.M22*other.M22 + m.M23*other.M32,
		M23: m.M21*other.M13 + m.M22*other.M23 + m.M23*other.M33,
		M31: m.M31*other.M11 + m.M32*other.M21 + m.M33*other.M31,
		M32: m.M31*other.M12 + m.M32*other.M22 + m.M33*other.M32,
		M33: m.M31*other.M13 + m.M32*other.M23 + m.M33*other.M33,
	}
}

func (m Matrix) TransformPoint(p Point) Point {
	return Point{
		X: p.X*m.M11 + p.Y*m.M12 + m.M13,
		Y: p.X*m.M21 + p.Y*m.M22 + m.M23,
	}
}

func (m Matrix) TransformPoint3(p Point3) Point3 {
	return Point3{
		X: p.X*m.M11 + p.Y*m.M12 + p.Z*m.M13,
		Y: p.X*m.M21 + p.Y*m.M22 + p.Z*m.M23,
		Z: p.X*m.M31 + p.Y*m.M32 + p.Z*m.M33,
	}
}

func Translate3D(x, y, z float32) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v Vec3) Scale(s float32) Vec3 {
	return Vec3{X: v.X * s, Y: v.Y * s, Z: v.Z * s}
}

func (v Vec3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3) Normalize() Vec3 {
	len := v.Length()
	if len == 0 {
		return Vec3{}
	}
	return Vec3{X: v.X / len, Y: v.Y / len, Z: v.Z / len}
}

func Cross(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Dot(a, b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

type GradientStop struct {
	Offset float32
	Color  Color
}

type LinearGradient struct {
	Start Point
	End   Point
	Stops []GradientStop
}

type RadialGradient struct {
	Center Point
	Radius float32
	Focus  Point
	Stops  []GradientStop
}

type BlendMode uint8

const (
	BlendModeSrcOver BlendMode = iota
	BlendModeSrcIn
	BlendModeSrcOut
	BlendModeSrcAtop
	BlendModeScreen
	BlendModeMultiply
	BlendModeOverlay
	BlendModeDarken
	BlendModeLighten
)

type Shadow struct {
	Color  Color
	X      float32
	Y      float32
	Blur   float32
	Spread float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func Hex(s string) Color {
	color, err := parseHex(s)
	if err != nil {
		return Hex("#e100ff")
	}
	return color
}

func parseHex(s string) (Color, error) {
	s = strings.TrimPrefix(s, "#")
	if len(s) != 6 && len(s) != 8 {
		return Color{}, fmt.Errorf("invalid color format: %q", s)
	}

	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	g, _ := strconv.ParseUint(s[2:4], 16, 8)
	b, _ := strconv.ParseUint(s[4:6], 16, 8)

	a := uint64(255)
	if len(s) == 8 {
		a, _ = strconv.ParseUint(s[6:8], 16, 8)
	}

	return Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}, nil
}

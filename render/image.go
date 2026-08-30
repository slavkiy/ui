package render

type Image struct {
	Width  int
	Height int
	Data   []byte

	texture *Texture
}

func NewImageRGBA(width, height int, data []byte) *Image {
	if width <= 0 || height <= 0 {
		return &Image{Width: 0, Height: 0, Data: nil}
	}
	needed := width * height * 4
	pixels := make([]byte, needed)
	if len(data) > 0 {
		copy(pixels, data[:min(len(data), needed)])
	}
	return &Image{Width: width, Height: height, Data: pixels}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Texture struct {
	handle uint32
	width  float32
	height float32
}

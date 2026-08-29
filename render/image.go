package render

type Image struct {
	texture *Texture
}

type Texture struct {
	handle uint32
	width  float32
	height float32
}

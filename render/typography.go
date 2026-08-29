package render

type TextStyle struct {
	Font  *Font
	Size  float32
	Color Color

	Weight FontWeight

	LetterSpacing float32
	LineHeight    float32
}

type FontWeight uint16

const (
	FontThin       FontWeight = 100
	FontExtraLight FontWeight = 200
	FontLight      FontWeight = 300
	FontNormal     FontWeight = 400
	FontMedium     FontWeight = 500
	FontSemiBold   FontWeight = 600
	FontBold       FontWeight = 700
	FontExtraBold  FontWeight = 800
	FontBlack      FontWeight = 900
)

type Font struct {
	Name string
}

package test

import (
	"testing"

	. "github.com/slavkiy/ui/widget"
)

func TestTextWidget(t *testing.T) {
	Text("Hello").Align(Up, Left)
}

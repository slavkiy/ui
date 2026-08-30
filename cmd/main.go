package main

import (
	"github.com/slavkiy/ui"
	. "github.com/slavkiy/ui/widget"
)

func main() {
	ui.State(
		Column(
			Text("Hello").Align(Up, Left),
		).Align(Center),
	)
}

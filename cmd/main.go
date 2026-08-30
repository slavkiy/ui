package main

import (
	"github.com/slavkiy/ui"
	. "github.com/slavkiy/ui/widget"
)

func main() {
	app := ui.NewApp()

	app.State(
		Column(
			Text("Hello").Align(Up, Left),
		).Align(Center),
	)

	app.Run()
}

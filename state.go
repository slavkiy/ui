package ui

import (
	. "github.com/slavkiy/ui/reactive"
	. "github.com/slavkiy/ui/widget"
)

var signalState = NewSignal[Widget]("telefy/ui::State", Text("Home page slavkiy/ui").Align(Center))

func State(w Widget) { signalState.Set(w) }

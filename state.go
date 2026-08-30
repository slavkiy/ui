package ui

import (
	. "github.com/slavkiy/ui/reactive"
	. "github.com/slavkiy/ui/widget"
)

var signalState = NewSignal[Widget]("telefy/ui::State", Text("Home page slavkiy/ui").Align(Center))

func State(w Widget) {
	signalState.Set(w)
}

type app struct {
	sub *Subscription
	sch <-chan Widget
}

func NewApp() *app {
	ch, sub := signalState.SubscribeChan(10240)

	return &app{
		sub: sub,
		sch: ch,
	}
}

func (a *app) Run() {
	defer a.sub.Unsubscribe()
	for w := range a.sch {
		_ = w
	}
}

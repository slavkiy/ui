package ui

import (
	"fmt"
	"time"

	. "github.com/slavkiy/ui/reactive"
	. "github.com/slavkiy/ui/widget"
)

var signalState = NewSignal[Widget]("ui/ui::State", Text("Home page slavkiy/ui").Align(Center))

func (app) State(w Widget) {
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

func renderCurrentState(w Widget) {
	if w == nil {
		fmt.Println("ui: no widget state")
		return
	}
	fmt.Printf("ui: rendering %T\n", w)
}

func (a *app) Run() {
	if a == nil {
		return
	}
	defer a.sub.Unsubscribe()

	if w := signalState.Get(); w != nil {
		renderCurrentState(w)
		return
	}

	for {
		select {
		case w, ok := <-a.sch:
			if !ok {
				return
			}
			renderCurrentState(w)
			return
		case <-time.After(50 * time.Millisecond):
			if w := signalState.Get(); w != nil {
				renderCurrentState(w)
				return
			}
		}
	}
}

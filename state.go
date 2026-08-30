package ui

import (
	"sync"

	. "github.com/slavkiy/ui/reactive"
	. "github.com/slavkiy/ui/widget"
)

var once sync.Once

var signalState = NewSignal[Widget]("telefy/ui::State", Text("Home page slavkiy/ui").Align(Center))

func State(w Widget) {
	once.Do(func() {
		go start()
	})
	signalState.Set(w)
}

func start() {
	ch, sub := signalState.SubscribeChan(1024)
	defer sub.Unsubscribe()
	for w := range ch {
		_ = w
	}
}

package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/slavkiy/ui/reactive"
)

type Text struct {
	Text *Signal[string]
}

func TestReactive(t *testing.T) {
	name := NewSignal("name", "John")

	text := Text{
		Text: name,
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(text *Text, ctx context.Context) {
		defer wg.Done()

		sub := text.Text.Subscribe(func(value string) {
			fmt.Println("Text updated:", value)
		})
		defer sub.Unsubscribe()

		<-ctx.Done()
	}(&text, ctx)

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			wg.Wait()
			return

		case <-time.After(500 * time.Millisecond):
			name.Set(fmt.Sprintf("slavik %d", i))
		}
	}
}

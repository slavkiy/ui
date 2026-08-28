package widget

import (
	"context"

	"github.com/slavkiy/ui/reactive"
)

type Widget interface {
	Build(ctx *context.Context) Widget
}

type Element struct {
	Widget   Widget
	Parent   *Element
	Children []*Element

	Subscriptions []*reactive.Subscription
}

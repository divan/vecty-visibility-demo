package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// Page is our main page component.
type Page struct {
	vecty.Core
}

// NewPage creates and inits new app page.
func NewPage() *Page {
	return &Page{}
}

// Render implements the vecty.Component interface.
func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			event.VisibilityChange(p.VisibilityListener),
		),
	)
}

// VisibilityListener implements listener for visibilitychange events.
// We use it to pause animation when the page is hidden, because WebGL
// animation is pretty CPU expensive.
// See https://developer.mozilla.org/en-US/docs/Web/API/Page_Visibility_API for details.
func (p *Page) VisibilityListener(e *vecty.Event) {
	document := js.Global.Get("document")
	hidden := document.Get("hidden")
	fmt.Println("Page is hidden:", hidden)
}

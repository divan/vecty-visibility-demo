package main

import (
	"github.com/gopherjs/vecty"
)

func main() {
	page := NewPage()

	vecty.SetTitle("VisibilityChange Demo")
	vecty.RenderBody(page)
}

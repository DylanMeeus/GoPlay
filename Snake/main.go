package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	run()
}

func run() {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()

	fmt.Printf("%v %v\n", w, h)

	body := js.Global.Get("document").Get("body")
	canvas := js.Global.Get("document").Call("createElement", "canvas")
	canvasCtx := canvas.Call("getContext", "2d")

	// Set adjusts the element properties
	canvas.Set("width", w)
	canvas.Set("height", h)

	canvasCtx.Set("fillStyle", "#000")

	// render the background
	canvasCtx.Call("fillRect", 0, 0, w, h)

	// Call runs a function against the object
	body.Call("appendChild", canvas)
}

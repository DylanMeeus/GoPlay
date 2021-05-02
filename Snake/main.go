package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	run()
}

type Point struct {
	x, y int
}

type Snake struct {
	positions []Point
}

type Game struct {
	Score       int
	Player      Snake
	TileWidth   float64
	TileHeight  float64
	TileRows    int
	TileColumns int
	Canvas      *js.Object
}

func setupGame() *Game {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()
	rows, columns := 10, 10

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
	return &Game{
		Player:      Snake{},
		Canvas:      canvas,
		TileWidth:   w / float64(rows),
		TileHeight:  h / float64(columns),
		TileRows:    rows,
		TileColumns: columns,
		Score:       0,
	}
}

func run() {
	g := setupGame()

	fps := time.Tick(1 * time.Second / 60)
	gameloop := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-fps:
			render(g)
		case <-gameloop:
			loop(g)
		}
	}

}

// main game loop
func loop(g *Game) {
}

// main render loop
func render(g *Game) {
	ctx := g.Canvas.Call("getContext", "2d")
	blue := "#0000ff"
	green := "#00ff00"
	for r := 0; r < g.TileRows; r++ {
		for c := 0; c < g.TileColumns; c++ {
			if c%2 == 0 {
				ctx.Set("fillStyle", blue)
			} else {
				ctx.Set("fillStyle", green)
			}
			tileStartX := float64(c) * g.TileWidth
			tileStartY := float64(r) * g.TileHeight
			ctx.Call("fillRect", tileStartX, tileStartY, float64(g.TileWidth), float64(g.TileHeight))
		}
		blue, green = green, blue
	}
}

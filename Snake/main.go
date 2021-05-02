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

type Square struct {
	x, y, w, h float64
}

func (p Point) ToCanvasSquare(g *Game) Square {
	x := float64(p.x) * g.TileWidth
	y := float64(p.y) * g.TileHeight
	return Square{
		x: x,
		y: y,
		w: g.TileWidth,
		h: g.TileHeight,
	}
}

type Snake struct {
	positions []Point
	velocity  Point
}

func (s *Snake) move() {
	// take one from the tail, and stick it in the front?
	head := s.positions[0]
	//tail := s.positions[len(s.positions)-1]
	//fmt.Printf("%v\n", tail)
	newHead := Point{head.x + s.velocity.x, head.y + s.velocity.y}

	s.positions = append([]Point{newHead}, s.positions...)
	s.positions = s.positions[:len(s.positions)-1]
}

type Game struct {
	Score       int
	Player      Snake
	TileWidth   float64
	TileHeight  float64
	TileRows    int
	TileColumns int
	Width       float64
	Height      float64
	Canvas      *js.Object
}

func setupGame() *Game {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()
	rows, columns := 50, 50

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
		Player: Snake{
			positions: []Point{{5, 5}},
			velocity:  Point{1, 0},
		},
		Canvas:      canvas,
		TileWidth:   w / float64(rows),
		TileHeight:  h / float64(columns),
		TileRows:    rows,
		TileColumns: columns,
		Width:       w,
		Height:      h,
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
	g.Player.move()
}

// main render loop
func render(g *Game) {
	renderBackground(g)
	renderPlayer(g)
	/*
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
	*/
}

func renderBackground(g *Game) {
	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "#000")
	ctx.Call("fillRect", 0, 0, g.Width, g.Height)
}

func renderPlayer(g *Game) {

	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "white")

	for _, point := range g.Player.positions {
		sq := point.ToCanvasSquare(g)
		ctx.Call("fillRect", sq.x, sq.y, sq.w, sq.h)
	}

}

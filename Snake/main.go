package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

const (
	SNAKE_MOVE_INTERVAL = 200
	FOOD_SPAWN_INTERVAL = 1_000
)

type direction int

func (d direction) ToVelocity() Point {
	switch d {
	case UP:
		return Point{0, -1}
	case DOWN:
		return Point{0, 1}
	case LEFT:
		return Point{-1, 0}
	case RIGHT:
		return Point{1, 0}
	default:
		return Point{0, 1}

	}
}

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

var (
	DIRECTION = direction(RIGHT)
)

type Point struct {
	x, y int
}

type Food Point

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

func (s *Snake) move(d direction) {
	// take one from the tail, and stick it in the front?
	s.velocity = d.ToVelocity()
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
	Foods       []Food // yeah okay, foods is not the plural, but it makes it clear it's a slice :-)
	TileWidth   float64
	TileHeight  float64
	TileRows    int
	TileColumns int
	Width       float64
	Height      float64
	Canvas      *js.Object
}

// SpawnFood selects a random location to spawn food
func (g *Game) SpawnFood() {
	maxX := g.TileColumns
	maxY := g.TileRows
	// TODO: make sure the player is not on the food already

	x, y := rand.Intn(maxX), rand.Intn(maxY)
	g.Foods = append(g.Foods, Food{x, y})
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

	// attach event listener for keypresses
	js.Global.Get("document").Call("addEventListener", "keydown", keyPressEvent, true)

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

func keyPressEvent(e *js.Object) {
	fmt.Println(e.Get("keyCode"))
	switch e.Get("keyCode").String() {
	case "65": // left
		DIRECTION = LEFT
	case "68": // right
		DIRECTION = RIGHT
	case "87": // up
		DIRECTION = UP
	case "83": // down
		DIRECTION = DOWN
	}
}

func run() {
	g := setupGame()
	go gameLoop(g)

	fps := time.Tick(1 * time.Second / 60)
	for {
		select {
		case <-fps:
			render(g)
		}
	}

}

// main game loop
func gameLoop(g *Game) {
	moveLoop := time.Tick(SNAKE_MOVE_INTERVAL * time.Millisecond)
	foodLoop := time.Tick(FOOD_SPAWN_INTERVAL * time.Millisecond)
	for {
		select {
		case <-moveLoop:
			// todo: add snake state here (dead || alive)
			g.Player.move(DIRECTION)
		case <-foodLoop:
			g.SpawnFood()
		}

	}
}

// main render loop
func render(g *Game) {
	renderBackground(g)
	renderPlayer(g)
	renderFood(g)
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

func renderFood(g *Game) {

	ctx := g.Canvas.Call("getContext", "2d")
	foodColour := "#d13017" // kinda reddish, maybe like an apple
	ctx.Set("fillStyle", foodColour)

	for _, food := range g.Foods {
		sq := Point(food).ToCanvasSquare(g)
		ctx.Call("fillRect", sq.x, sq.y, sq.w, sq.h)
	}

}

func main() {
	run()
}

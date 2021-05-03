package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

const (
	SNAKE_MOVE_INTERVAL = 50
	FOOD_VALUE          = 1000
)

type GameState int

const (
	RUNNING GameState = iota
	GAME_OVER
	PAUSE
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
	length    int
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
	CurrentState GameState
	Score        int
	Player       Snake
	Food         Food
	TileWidth    float64
	TileHeight   float64
	TileRows     int
	TileColumns  int
	Width        float64
	Height       float64
	Canvas       *js.Object
}

// SpawnFood selects a random location to spawn food
func (g *Game) SpawnFood() {
	maxX := g.TileColumns
	maxY := g.TileRows
	// TODO: make sure the player is not on the food already

	x, y := rand.Intn(maxX), rand.Intn(maxY)
	g.Food = Food{x, y}
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
		Canvas:       canvas,
		TileWidth:    w / float64(rows),
		TileHeight:   h / float64(columns),
		TileRows:     rows,
		TileColumns:  columns,
		Width:        w,
		Height:       h,
		Score:        0,
		CurrentState: RUNNING,
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
	g.SpawnFood()
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
	for {
		switch g.CurrentState {
		case RUNNING:
			select {
			case <-moveLoop:
				// todo: add snake state here (dead || alive)
				g.Player.move(DIRECTION)
				if g.boundsCollisionDetection() || g.snakeEatsSnakeDetection() {
					// end game
					g.CurrentState = GAME_OVER
				}
				g.foodCollisionDetection()
			}
		default:
			// clear the buffers
			_ = <-moveLoop
		}
	}
}

func (g *Game) foodCollisionDetection() {
	// if the player is on food, the player becomes longer
	if Point(g.Food) == g.Player.positions[0] { // if the head touches the food
		g.Player.positions = append(g.Player.positions, Point(g.Food))
		g.Player.length++
		g.Score += FOOD_VALUE
		g.SpawnFood()
	}
}

func (g *Game) boundsCollisionDetection() bool {
	x, y := g.Player.positions[0].x, g.Player.positions[0].y
	return (x < 0 || x >= g.TileColumns || y < 0 || y > g.TileColumns)
}

// fun function name, isn't it.
func (g *Game) snakeEatsSnakeDetection() bool {
	head := g.Player.positions[0]
	for _, bodyPart := range g.Player.positions[1:] {
		if bodyPart == head {
			return true
		}
	}
	return false
}

// main render loop
func render(g *Game) {
	switch g.CurrentState {
	case RUNNING:
		renderRunning(g)
	default:
		renderGameOver(g)
	}
}

func renderGameOver(g *Game) {
	fmt.Println("rendering game over screen")
	renderBackground(g)

	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("font", "50px Arial")
	ctx.Set("fillStyle", "#fff")
	centreX, centreY := g.Width/2, g.Height/2
	ctx.Call("fillText", fmt.Sprintf("Game Over! %v points", g.Score), centreX, centreY)
	ctx.Call("fillText", "Press space to continue!", centreX, centreY+50)
}

func renderRunning(g *Game) {
	renderBackground(g)
	renderPlayer(g)
	renderFood(g)
	renderScore(g)
}

func renderScore(g *Game) {
	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("font", "20px Arial")
	ctx.Set("fillStyle", "#fff")
	ctx.Call("fillText", fmt.Sprintf("Score: %v", g.Score), 10, 50)
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

	sq := Point(g.Food).ToCanvasSquare(g)
	ctx.Call("fillRect", sq.x, sq.y, sq.w, sq.h)

}

func main() {
	run()
}

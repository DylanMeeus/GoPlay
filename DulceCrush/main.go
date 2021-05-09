package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

const (
	// POLLING_INTERVAL to check for user input??
	POLLING_INTERVAL = 100
)

type GameState int

const (
	RUNNING GameState = iota
	GAME_OVER
	PAUSE
)

var (
	SPACEBAR_PRESSED = false
	// global buffer for handling player clicks
	clickBuffer = ClickBuffer{}
)

var (
	// TODO: associate the candies with images instead of colours
	CandyImage = map[int]string{
		0: "red",
		1: "green",
		2: "blue",
		3: "white",
		4: "yellow",
		5: "orange",
	}

	CandyN = len(CandyImage)
)

type Board [][]int

// Point identifies a logical location on the board (81 tiles per game)
type Point struct {
	x, y int
}

// PixelPoint identifies a point by the pixel coordinates
type PixelPoint Point

type ClickBuffer [2]PixelPoint

// Push an element to the pixel buffer
func (c *ClickBuffer) Push(p PixelPoint) {
	if c[0] == (PixelPoint{}) {
		c[0] = p
	} else {
		c[1] = p
	}
}

func (c *ClickBuffer) Len() (count int) {
	for _, p := range c {
		if p != (PixelPoint{}) {
			count++
		}
	}
	return
}

// Empty the buffer by resetting to default zero values
func (c *ClickBuffer) Empty() {
	c[0], c[1] = PixelPoint{}, PixelPoint{}
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

// CoordinatesToPoint returns the point located at the coordinates
func CoordinatesToPoint(x, y int, g *Game) Point {

	return Point{}
}

type Game struct {
	CurrentState GameState
	Board        Board
	Score        int
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
	_, _ = x, y
}

func NewBoard(rows, columns int) Board {
	b := make([][]int, rows)

	for i := 0; i < len(b); i++ {
		b[i] = make([]int, columns)
	}

	return b
}

// Randomize enters random candies in the board
func (g *Game) PopulateBoard() {
	for i := 0; i < g.TileRows; i++ {
		for j := 0; j < g.TileColumns; j++ {
			candy := rand.Intn(CandyN)
			g.Board[i][j] = candy
		}
	}
}

func setupGame() *Game {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()
	rows, columns := 9, 9

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
	js.Global.Get("document").Call("addEventListener", "click", mouseClickEvent, true)

	// Call runs a function against the object
	body.Call("appendChild", canvas)
	return &Game{
		Canvas:       canvas,
		Board:        NewBoard(rows, columns),
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

func mouseClickEvent(e *js.Object) {
	clickX, clickY := e.Get("pageX").String(), e.Get("pageY").String()
	x, _ := strconv.Atoi(clickX)
	y, _ := strconv.Atoi(clickY)
	clickBuffer.Push(PixelPoint{x, y})
	fmt.Printf("%v\n", clickBuffer)
}

func keyPressEvent(e *js.Object) {
	fmt.Println(e.Get("keyCode"))
}

// TODO: mouse events

func run() {
	g := setupGame()
	g.PopulateBoard()
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
	actionLoop := time.Tick(POLLING_INTERVAL * time.Millisecond)
	for {
		switch g.CurrentState {
		case RUNNING:
			select {
			case <-actionLoop:
				g.processInputBuffer()
				// todo: add snake state here (dead || alive)
			}
		case GAME_OVER:
			_ = <-actionLoop
			g.pauseScreenLoop()
		default:
			// clear the buffers
			_ = <-actionLoop
		}
	}
}

// proccesInputBuffer checks the input buffer
// and takes appropriate actions
func (g *Game) processInputBuffer() {
	if clickBuffer.Len() == 2 {
		clickBuffer.Empty()
	}
	fmt.Printf("buf: %v\n", clickBuffer)
}

func (g *Game) resetGame() {
	// TODO: implement me
}

func (g *Game) pauseScreenLoop() {
	if SPACEBAR_PRESSED {
		SPACEBAR_PRESSED = false
		g.resetGame()
	}
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
	renderBoard(g)
	renderScore(g)
}

func renderScore(g *Game) {
	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("font", "20px Arial")
	ctx.Set("fillStyle", "#fff")
	ctx.Call("fillText", fmt.Sprintf("Score: %v", g.Score), 10, 50)
	ctx.Call("fillText", "use WASD to move", 10, 75)
}

func renderBackground(g *Game) {
	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "#000")
	ctx.Call("fillRect", 0, 0, g.Width, g.Height)
}

func renderBoard(g *Game) {

	ctx := g.Canvas.Call("getContext", "2d")
	ctx.Set("fillStyle", "white")

	for i := 0; i < g.TileRows; i++ {
		for j := 0; j < g.TileColumns; j++ {
			sq := Point{i, j}.ToCanvasSquare(g)
			ctx.Set("fillStyle", CandyImage[g.Board[i][j]])
			ctx.Call("fillRect", sq.x, sq.y, sq.w, sq.h)
		}
	}

}

func main() {
	run()
}

package main


import (
    "fmt"
)

// instead of starting at the starting point, start at the three corners?
type Point struct{
    x int
    y int
}

type Triangle struct{
    a Point
    b Point
    c Point
}

type Line struct{
    p Point
    q Point
}
// for testing
var ABC = Triangle{Point{-340,495}, Point{-153,-910}, Point{835,-947}}
var DEF = Triangle{Point{-175,41}, Point{-421,-714}, Point{574,-645}}


func main(){
    lines := createLines(ABC)
    for i := 0; i < len(lines); i++{
        walkLine(lines[i])
        break
    }
    //
    //t := Line{Point{x:0, y:0}, Point{x:10,y:10}}
    //walkLine(t)
}

func walkLine(line Line) bool{
    // determine equation

    // todo: find original direction of the point wrt the line.

    fmt.Println(line)
    // m = (y1 - y0) / (x1 - x0)
    // p = 1, q = 0
    slope := (line.p.y - line.q.y) / (line.p.x - line.q.x)
    startX := line.p.x
    originalX := line.q.x
    originalY := line.q.y
    for ; startX < line.q.x ; startX++{
        // xy = m(x-x0)+y0)
        xy := slope * (startX - originalX) + originalY
        point := Point{x: startX, y:xy}
    }

    return false
}

// create lines for the triangles
func createLines(triangle Triangle) []Line {
    ab := Line{p:triangle.a,q:triangle.b}
    bc := Line{p:triangle.b,q:triangle.c}
    ca := Line{p:triangle.c,q:triangle.a}
    lines := []Line{ab,bc,ca}
    return lines
}
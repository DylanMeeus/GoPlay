package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "math/rand"
    "strconv"
)


type Cell struct{
    alive bool // 0 == dead, 1 == alive
}

const rows int = 50
const cols int = 50
const timeout time.Duration = 50 // time in milliseconds between cycles

func main(){
    cells := make([][]Cell, rows)
    for i := range cells {
        cells[i] = make([]Cell, cols)
    }
    cells = initField(&cells, randomBool)

    for generation := 0; ; generation++ {
        clearScreen()
        printGame(&cells)
        timeStepTime := timeStep(&cells)
        fmt.Println("generation: " + strconv.Itoa(generation))
        fmt.Println("generated generation in: " + strconv.Itoa(int(timeStepTime)) + " Microseconds")
        time.Sleep(timeout * time.Millisecond)
    }
}

// a function that randomly returns either true or false
func randomBool() bool{
    rand.Seed(time.Now().UTC().UnixNano())
    i := rand.Int31()
    return i & 1 == 1
}


func shallowCopy(field *[][]Cell) [][]Cell{
    dereferenced := *field
    arr := make([][]Cell, rows)
    for i := range arr {
        arr[i] = make([]Cell, cols)
    }
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            arr[row][col] = dereferenced[row][col]
        }
    }
    return arr
}

// manipulate the field one 'step'
func timeStep(field *[][]Cell) float64{ // alters the array under the pointer, returns time for running method
    startTime := time.Now()
    dereferenced := *field
    arr := shallowCopy(field)
    for row := 0; row < rows; row++{
        for col := 0; col < cols; col++{
            cell := dereferenced[row][col]
            neighbours := 0
            // init movemenet

            startr := func(row int) int{
                if row > 0{
                    return row - 1
                }
                return row
            }(row)

            startc := func(col int) int{
                if col > 0{
                    return col -1
                }
                return col
            }(col)

            for rowi := startr; rowi <= startr + 2; rowi++{ // start from row-1, to row+1
                if rowi >= rows {
                    break
                }
                for coli := startc; coli <= startc + 2; coli++{
                    if coli >= cols || rowi == row && coli == col { // skip the current one!
                        continue
                    } else {
                        if dereferenced[rowi][coli].alive {
                            neighbours++
                        }
                    }
                }
            }

            newCell := getNewCellState(cell, neighbours, cell.alive)
            arr[row][col] = newCell
        }
    }

    // change field to the copy
    *field = arr

    return float64(time.Now().Sub(startTime)) / float64(1000)
}

func getNewCellState(cell Cell, neighbours int, oldstate bool) Cell {
    c := Cell{}
    switch neighbours {
    // just the cases for live, death is default
        case 2:
            if  oldstate {      // staying alive
                c.alive = true
            }
            break
        case 3:
            c.alive = true      // staying alive or being born
            break
        default: c.alive = false
            break
    }
    return c
}

/**
initialize a field based on a function
 */
func initField(field *[][]Cell, result func() bool) [][]Cell {
    // set them all to false
    dereferenced := *field
    for row := 0; row < rows; row++{
        for col := 0; col < cols; col++{
            dereferenced[row][col].alive = result()
        }
    }
    return *field
}


// print the game
func printGame(gameArea *[][]Cell){
    dereferenced := *gameArea
    for row := 0; row < rows; row++{
        for col := 0; col < cols; col++{
            cell := dereferenced[row][col]
            if cell.alive{
                fmt.Print(" * ")
            } else {
                fmt.Print(" . ")
            }
        }
        fmt.Print("\n")
    }
}

// Clear the CLI window for a repaint
func clearScreen(){
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

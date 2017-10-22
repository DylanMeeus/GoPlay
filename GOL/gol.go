package main

import (
    "fmt"
)


type Cell struct{
    alive bool // 0 == dead, 1 == alive
}

const rows = 10
const cols = 10

func main(){
    cells := make([][]Cell, rows)
    for i := range cells {
        cells[i] = make([]Cell, cols)
    }
    cells = initField(&cells)

    // todo: make this an infinite loop in the future
    for i := 0; i < 1; i++{
        printGame(&cells)
        timeStep(&cells)
    }
}

// manipulate the field one 'step'
func timeStep(field *[][]Cell){
    arr := make([][]Cell, rows)
    copy(arr, *field) // copy field to arr
    dereferenced := *field
    for row := 0; row < rows; row++{
        for col := 0; col < cols; col++{
            cell := dereferenced[row][col]
            neighbours := 0
            // init movemenet
            nextCol := col + 1
            prevCol := col - 1
            nextRow := row + 1
            prevRow := row - 1

            if !isLastColumn(col) {
                if dereferenced[row][nextCol].alive {
                    neighbours++
                }
            }
            if !isLastColumn(col) && !isLastRow(row) {
                if dereferenced[nextRow][nextCol].alive {
                    neighbours++
                }
            }
            if !isLastRow(row) {
                if dereferenced[nextRow][col].alive{
                    neighbours++
                }
            }
            if !isLastRow(row) && !isFirstColumn(col) {
                if dereferenced[nextRow][prevCol].alive{
                    neighbours++
                }
            }
            if !isFirstColumn(col) {
                if dereferenced[row][prevCol].alive {
                    neighbours++
                }
            }
            if !isFirstColumn(col) && !isFirstRow(row) {
                if dereferenced[prevRow][prevCol].alive {
                    neighbours++
                }
            }
            if !isFirstRow(row) {
                if dereferenced[prevRow][col].alive {
                    neighbours++
                }
            }
            if !isFirstRow(row) && !isLastColumn(col) {
                if dereferenced[prevRow][nextCol].alive {
                    neighbours++
                }
            }
            newCell := getNewCellState(cell, neighbours)
            arr[row][col] = newCell
        }
    }

    // change field to the copy
    *field = arr
}

func getNewCellState(cell Cell, neighbours int) Cell {
    fmt.Println(neighbours)
    c := Cell{}
    switch neighbours {
    case 3:
        c.alive = true
        break
    }
    return c
}

func isLeftTopCorner(row int, col int) bool {
    return isFirstRow(row) && isFirstColumn(col)
}

func isLeftBottomCorner(row int, col int) bool {
    return isLastRow(row) && isFirstColumn(col)
}

func isRightTopCorner(row int, col int) bool {
    return isFirstRow(row) && isLastColumn(col)
}

func isRightBottomCorner(row int, col int) bool {
    return isLastRow(row) && isLastColumn(col)
}

func isFirstRow(row int) bool {
    return row == 0
}

func isLastRow(row int) bool {
    return row == rows-1
}

func isFirstColumn(col int) bool {
    return col == 0
}

func isLastColumn(col int) bool {
    return col == cols-1
}

func countNeighbours(cell *Cell, field *[][]Cell) int{
    neighbours := 0
    return neighbours
}



func initField(field *[][]Cell) [][]Cell {
    // set them all to false
    dereferenced := *field
    for row := 0; row < rows; row++{
        for col := 0; col < cols; col++{
            dereferenced[row][col].alive = false
        }
    }

    // set some to true for testing
    dereferenced[5][5].alive = true
    dereferenced[5][6].alive = true
    dereferenced[6][5].alive = true

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

//func runForEachCell(f func()){
//    for row := 0; row < rows; row++{
//        for col := 0; col < cols; col++{
//            field[row][col] = f()
//        }
//    }
//}

// Clear the CLI window for a repaint
func clearScreen(){

}

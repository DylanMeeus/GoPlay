package main

import(
    "fmt"
)

const (
    n = 8
)

type board [][]bool

func main() {
    b := make(board, n)
    for i,_ := range b {
        b[i] = make([]bool,n)
    }
    b[0][1] = true
    b[1][0] = true
    fmt.Println(b)
    fmt.Printf("%v\n", valid(b))
}

func (b board) String() (s string) {
    for _, row := range b {
        for _, col := range row {
            if col {
                s += " 1 "
            } else {
                s += " 0 "
            }
        }
        s += "\n"
    }
    return 
}

// checks if the board is in a valid state
// this means: no Queens on same row, column, diagonal
func valid(b board) bool {
    return validDiagonal(b)
}

// check that all diagonals have only one queen..
func validDiagonal(b board) bool {
    for col := 0; col < len(b); col++ {
        qd := false
        startRow,startCol := 0, col
        // check rightward diagonal
        for startRow < len(b) && startCol < len(b) {
            if b[startRow][startCol] {
                if qd {
                    return false
                }
                qd = true
            }
            startRow += 1
            startCol += 1
        }
        startRow, startCol = 0, col
        // check leftward diagonal
        for startRow < len(b) && startCol > 0 {
            if b[startRow][startCol] {
                if qd {
                    return false
                }
                qd = true
            }
            startRow += 1
            startCol -= 1
        }
    }
    return true 
}

func validHoriVerti(b board) bool {
    rowQueens := make(map[int]bool)
    colQueens := make(map[int]bool)
    for ri,row := range b {
        for ci,col := range row {
            if col {
                if rowQueens[ri] {
                    return false
                }
                rowQueens[ri] = true
                if colQueens[ci] {
                    return false
                }
                colQueens[ci] = true
            }
        }   
    }
    return true
}

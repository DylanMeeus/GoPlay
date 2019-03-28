package main

import(
    "fmt"
)

const (
    n = 6
)

type board [][]bool

func main() {
    fmt.Println("solving..")
    b := make(board, n)
    for i,_ := range b {
        b[i] = make([]bool,n)
    }
    _, solved := (solve(b,0))
    fmt.Println(solved)
}


func solve(b board, queens int) (bool, *board) {
    if queens == n && valid(b) {
        return true, &b
    }
    if queens >= n {
        return false, nil
    }
    // place a queen on an empty slot..
    for r := 0; r < n; r++{
        // if this row already has a queen, continue..
        for _,tile := range b[r] {
            if tile {
                continue
            }
        }
        for c := 0; c < n; c++ {
            // we try to place one here, on a copy of the board
            // if the column already has a queen, ignore it..
            for row := 0; row < n; row++ {
                if b[row][c] {
                    continue
                }
            }
            if b[r][c] == true {
                continue
            }
            copyBoard := deepCopy(b)
            copyBoard[r][c] = true
            if s, sb := solve(copyBoard, queens + 1); s {
                return true, sb
            }
        }
    }
    return false, nil
}


func deepCopy(b board) board {
    newb := make(board, n) 
    for i,_:= range newb {
        row := make([]bool, n)
        for j,_ := range row {
            row[j] = b[i][j]
        }
        newb[i] = row
    }
    return newb
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
    return validDiagonal(b) && validHoriVerti(b)
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
        qd = false // sanity check really :)
        for startRow < len(b) && startCol >= 0 {
            if b[startRow][startCol] {
                if qd {
                    return false
                }
                qd = true
            }
            startRow += 1
            startCol -= 1
        }
        
        // and from the bottom..
        qd = false
        startRow, startCol = len(b) - 1, col
        for startRow >= 0 && startCol < len(b) {
            if b[startRow][startCol] {
                if qd {
                    return false
                }
                qd = true
            }
            startRow -= 1
            startCol += 1
        }

        qd = false
        startRow, startCol = len(b) - 1, col
        for startRow >= 0 && startCol >= 0 {
            if b[startRow][startCol] {
                if qd {
                    return false
                }
                qd = true
            }
            startRow -= 1
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

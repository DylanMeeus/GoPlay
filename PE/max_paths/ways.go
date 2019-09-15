package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

const (
    size = 80 
    index_size = 79 
)

var (
    lowest_score = -1
)

func main() {
    mat := load()
    solve(mat, 0, 0, 0)
    fmt.Printf("%v\n", lowest_score)
}

func solve(matrix [][]int, row, col, score int) {
    new_score := score + matrix[row][col]
    if lowest_score != -1 && new_score > lowest_score {
        return
    }
    if row == index_size && col == index_size {
        if lowest_score == -1 {
            lowest_score = new_score 
            return
        }
        if new_score < lowest_score {
            lowest_score = new_score
            return
        }
    }
    if row != index_size {
        go solve(matrix, row + 1, col, new_score)
    }
    if col != index_size {
        go solve(matrix, row, col + 1, new_score)
    }
}


func load() [][]int {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    data := make([][]int, size)
    for i,_ := range data {
        data[i] = make([]int, size)
    }
    lines := strings.Split(string(b), "\n")
    for i, line := range lines {
        nums := strings.Split(line, ",")
        for j, num := range nums {
            if num == "" {
                continue
            }
            n, err := strconv.Atoi(num)
            if err != nil {
                panic(err)
            }
            data[i][j] = n
        }
    }
    return data 
}

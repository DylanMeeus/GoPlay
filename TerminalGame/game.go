package main

/*
#include <stdio.h>
#include <unistd.h>
#include <termios.h>
char getch(){
    char ch = 0;
    struct termios old = {0};
    fflush(stdout);
    if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
    if( read(0, &ch,1) < 0 ) perror("read()");
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
    return ch;
}
*/
import "C"


import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

type field [][]string

const (
    width, height = 10,20
    padding = "            "
)

type point struct {
    x,y int
}

type tetromino struct {
    location []*point
}

type game struct {
    field field
    score int
    active *tetromino
}


// tetris game in the CLI
func main() {
    // start the game
    gameArea := make(field, height)
    for i,_ := range gameArea {
        gameArea[i] = make([]string, width)
    }
    gameArea[5][5] = "#"
    g := game{
        field: gameArea,
        score: 0,
    }
    for true {
        g.loop()
        clearScreen()
        g.render()
        time.Sleep(1 * time.Second)
    }
}

// main game-loop
func (g game) loop() {
    // read input
    //i := input()
    //fmt.Println(i)

    // move tetromino
    // check collision
}


func (g game) moveTetromino() {
    if g.active == nil {
        return
    }
}

// main render-loop. It really just prints to terminal
func (g game) render() {
    f := g.field
    fmt.Printf("\n\n\n")
    for _, row := range f {
        fmt.Printf("%v|", padding)
        for _, col := range row {
            if col == "" {
                fmt.Print(" ")
            } else {
                fmt.Print(col)
            }
        }   
        fmt.Printf("|\n")
    }
}

func input() string {
    in := C.getch()
    fmt.Printf("char: %v\n", in)
    return "" 
}

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}



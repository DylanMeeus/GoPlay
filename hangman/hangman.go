package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

// some drawings

var error1 string =
`



----`

var error2 string =
`
    |
    |
    |
    |
-----`

var error3 string =
`
----|
    |
    |
    |
-----`

var error4 string =
`
____
|   |
o   |
    |
    |
----
`

var error5 string =
`
____
|   |
o   |
|   |
    |
----
`

var error6 string =
`
_____
 |   |
 o   |
 |   |
/ \  |
-----
`

var error7 string =
`
_____
 |   |
 o   |
/|\  |
/ \  |
-----
`

func main(){
    RunGame()
}

func RunGame(){
    reader := bufio.NewReader(os.Stdin)
    currentGame := Game{Errors: 0, Word : "Esposita"}
    currentGame.Init()

    for {
        fmt.Print("Enter character: ")
        text, _ := reader.ReadString('\n')
        currentGame.ParseInput(text)
        gameWin := CheckWin(currentGame)
        gameLost := CheckLoss(currentGame)
        if gameWin{
            fmt.Println("Congratulations, you won!")
            os.Exit(1)
        }
        if gameLost{
            fmt.Println("You lost, try again!")
            os.Exit(1)
        }
    }
}

func CheckWin(g Game) bool{
    if g.Word == g.FoundWord{
        return true
    }
    return false
}

func CheckLoss(g Game) bool{
    if g.Errors == 7{
        return true
    }
    return false
}

func (game* Game) ParseInput(Input string) {
    if len(Input) != 2 { // 2 because of the newline
        fmt.Println("Please enter only one character")
        return
    }

    contains := strings.Contains(game.Word, string(Input[0]))
    if contains {
        newString := ""
        for i := 0; i < len(game.FoundWord); i++ {
            if string(game.Word[i]) == string(Input[0]){
                newString += string(game.Word[i])
            } else {
                newString += string(game.FoundWord[i])
            }
        }
        fmt.Println(newString)
        game.FoundWord = newString
    } else {
        fmt.Println("Wrong!!")
        game.Errors += 1
        game.DrawHangman()
    }
}

type Game struct {
    Errors int
    Word string
    FoundWord string // a word containing * for the unguesssed characters
}

func (g* Game) Init() {
    wordLen := len(g.Word) // here len == bytecount because we will just use standard english words
    g.FoundWord = strings.Repeat("*",wordLen)
}

func (g Game) DrawHangman() {
    var drawingString string
    switch g.Errors{
        case 1: drawingString = error1
        break
    case 2: drawingString = error2
        break
    case 3: drawingString = error3
        break
    case 4: drawingString = error4
        break
    case 5: drawingString = error5
        break
    case 6: drawingString = error6
        break
    case 7: drawingString = error7
        break
    }

    fmt.Println(drawingString)
}
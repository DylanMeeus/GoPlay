package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// url for current comic 
const currentComic = "https://xkcd.com/info.0.json" 

type Comic struct {
    Num int 
    Transcript string
    Alt string
    Title string
}

func main() {
    fmt.Printf("%v\n", fetchComic(currentComic))
}


// create an index of all the xkcd comics
func fetchComic(url string) *Comic {
    resp, err := http.Get(currentComic)
    if err != nil {
        fmt.Println("something went wrong!")
    }
    var current Comic
    err = json.NewDecoder(resp.Body).Decode(&current)
    if err != nil {
        fmt.Println("Something went wrong during decoding")
    }
    return &current
}

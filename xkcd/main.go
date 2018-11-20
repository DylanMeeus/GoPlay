package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    "io/ioutil"
)

// url for current comic 
const baseUrl = "https://xkcd.com/"
const jsonUrl = "info.0.json"
const currentComic = baseUrl + jsonUrl
const indexFile = "index.json"

type Comic struct {
    Num int 
    Transcript string
    Alt string
    Title string
}

func main() {
    //buildIndex()
    readIndex()
}

/*
* If no index was found: builds a new index of all XKCD comics
* Else: Add the missing comics to the index
*/
func buildIndex() {
    latest := *fetchComic(currentComic)
    latest.Num = 10 // testing
    comics := make([]*Comic, latest.Num)
    for i := 0; i < latest.Num; i++ {
        url := baseUrl + strconv.Itoa(i) + "/" + jsonUrl
        fmt.Println(url)
        comics[i] = fetchComic(url)
    }
    saveIndex(&comics)
}

func readIndex() []*Comic {
    var comics []Comic
    data, err := ioutil.ReadFile(indexFile)
    if err != nil {
        return make([]*Comic,0)
    }
    err = json.Unmarshal(data, &comics)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v\n",comics)
    return make([]*Comic, 0)
}

func saveIndex(comics *[]*Comic) {
    jsonBytes,_ := json.Marshal(comics)
    ioutil.WriteFile(indexFile,jsonBytes, 0644)
}

// create an index of all the xkcd comics
func fetchComic(url string) *Comic {
    resp, err := http.Get(url)
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

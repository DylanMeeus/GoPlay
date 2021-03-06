package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    "io/ioutil"
    "strings"
    "errors"
    "log"
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
    Img string
    Title string
}


func runServer() {
    comics := getComics()
    wordMap := buildSearchableMap(comics)
    handler := func (w http.ResponseWriter, r *http.Request) {
        search := r.URL.Path[1:]
        bestMatch,err := findBestMatch(&wordMap, search, selectClosestMatch)
        if err != nil {
            fmt.Fprint(w, "no comic found!")
        }
        fmt.Fprintf(w, "<body><img src=\"" + bestMatch.Img + "\"></body>")
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080",nil))
}

func main() {
    runServer()
    /*
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("enter text: ")
        text,_ := reader.ReadString('\n')
        bestMatch, err := findBestMatch(&wordMap, text, selectClosestMatch)
        if err != nil {
            fmt.Println("No comic found!")
        } else {
            fmt.Printf("%v\n", bestMatch.Num)
        }
    }
    */
}

func selectClosestMatch(matches map[Comic]int, search string) *Comic {
    // now return the closest matching comic
    // we prefer the title as matching factor
    searchTerms := strings.Split(search, " ")
    for k,_ := range matches {
        matchesAll := true 
        for _,term := range searchTerms {
            matchesAll = matchesAll && strings.Contains(strings.ToLower(k.Title), strings.ToLower(term))
        }
        if matchesAll {
            fmt.Println("title matches all")
            return &k
        }
    }
    // else look for frequency of matches
    var max int
    var frequentComic *Comic
    for k,v := range(matches) {
        if frequentComic == nil {
            frequentComic = &k
        }
        if v > max {
            max = v
            frequentComic = &k
        }
    }
    return frequentComic
}

func findBestMatch(wordMap *map[string][]Comic, 
                    search string,
                    selection func(map[Comic]int, string)*Comic) (*Comic, error) {
    // if no comic is found, then what?
    // for each word, find all the comics related. 
    // then find the most common occuring comic (it matches most strings)
    matchingComics := make(map[Comic]int, 0)
    search = strings.Replace(search,"\n", "", -1)
    for _,s := range(strings.Split(search, " ")) {
        if comics, ok := (*wordMap)[s]; ok {
            for _,c := range(comics){
                matchingComics[c]++
            }
        }
    }
    selected := selectClosestMatch(matchingComics, search) 
    if selected != nil {
        return selected, nil
    }
    return nil, errors.New("No comic found!")
}

func sanitizeString(in string) (out string) {
    out = in
    for _,s := range []string{"?","!","]","[",".", "\n", ")", "(", "*", "-"} {
        out = strings.Replace(out, s, " ", -1)
    }
    return 
}

func buildSearchableMap(comics *[]Comic) map[string][]Comic {
    // for each word in the comics
    // add it to a hashmap with a list of comics in which it appears
    // allows for O(1) lookup of those comics
    wordMap := make(map[string][]Comic, 1000) // 1000 chosen as 'best guess' for amount of words
    for _, c := range(*comics) { 
        var search string
        search += strconv.Itoa(c.Num) + " " 
        search += c.Transcript + " "
        search += c.Alt + " "
        search += c.Title + " "
        search = sanitizeString(strings.ToLower(search))
        for _,s := range strings.Split(search, " ") {
            if s == " " {
                continue
            }
            if slice, ok := wordMap[s]; ok {
                wordMap[s] = append(slice, c)
            } else {
                wordMap[s] = []Comic{c}
            }
        }
    }
    return wordMap
}

func getComics() *[]Comic{
    comics := readIndex()
    latestSavedComic := func(comics *[]Comic) int {
        var max int
        for _,c := range(*comics) {
            if c.Num > max{
                max = c.Num
            }
        }
        return max
    }(&comics)

    comicChan := make(chan Comic)
    go fetchComic(currentComic, comicChan)
    latest := <-comicChan 
    if latestSavedComic == latest.Num {
        // our index file is up to date!
        return &comics
    }
    buildIndex(latest)
    return getComics() // recursive call after building the index
}

/*
* If no index was found: builds a new index of all XKCD comics
* Else: Add the missing comics to the index
*/
func buildIndex(latest Comic) {
    comics := make([]Comic, latest.Num)
    comicChan := make(chan Comic)
    comics = append(comics, latest)
    for i := 1; i < latest.Num; i++ {
        url := baseUrl + strconv.Itoa(i) + "/" + jsonUrl
        fmt.Println(url)
        go fetchComic(url, comicChan)
        comics[i] = <-comicChan
    }
    saveIndex(&comics)
}

func readIndex() []Comic {
    var comics []Comic
    data, err := ioutil.ReadFile(indexFile)
    if err != nil {
        return make([]Comic,0)
    }
    err = json.Unmarshal(data, &comics)
    if err != nil {
        panic(err)
    }
    return comics
}

func saveIndex(comics *[]Comic) {
    jsonBytes,_ := json.Marshal(comics)
    ioutil.WriteFile(indexFile,jsonBytes, 0644)
}

// create an index of all the xkcd comics
func fetchComic(url string, channel chan Comic) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("something went wrong!")
    }
    var current Comic
    err = json.NewDecoder(resp.Body).Decode(&current)
    if err != nil {
        fmt.Println("Something went wrong during decoding")
    }
    channel <-current
}

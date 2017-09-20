package main


// file for the REST-API endpoints
import (
    "net/http"
    "fmt"
    "encoding/json"
)

func GetUsers(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintln(writer,"{['dylan','ana']}")
}

func GetTweets(writer http.ResponseWriter, request *http.Request){
    tweets := DatabaseTweets()
    str, _ := json.Marshal(tweets)
    fmt.Fprint(writer,string(str))
}
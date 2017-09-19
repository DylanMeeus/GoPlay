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
    tweets := Tweets{
        Tweet{Username: "Dylan", Tweetbody:"Hello world! this is my first tweet!"},
        Tweet{Username: "Ana", Tweetbody:"Hola Mundo! Esto es mi primero tweet!"},
    }
    str, _ := json.Marshal(tweets)
    fmt.Fprint(writer,string(str))

}
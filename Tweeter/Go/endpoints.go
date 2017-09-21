package main


// file for the REST-API endpoints
import (
    "net/http"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

func GetUsers(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintln(writer,"{['dylan','ana']}")
}

func GetTweets(writer http.ResponseWriter, request *http.Request){
    tweets := DatabaseTweets()
    str, _ := json.Marshal(tweets)
    fmt.Fprint(writer,string(str))
}

func Login(writer http.ResponseWriter, request *http.Request){
    err := request.ParseForm()
    if err != nil{
        panic(err)
    }
    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        panic(err)
    }

    var user User
    json.Unmarshal(body,&user)
    token := DatabaseLogin(user)
    fmt.Fprint(writer, token)
}
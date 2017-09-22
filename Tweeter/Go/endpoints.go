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

func ProfileTweets(writer http.ResponseWriter, request *http.Request){
    if(request.Method == "OPTIONS"){
        writer.Header().Add("Access-Control-Allow-Headers", "Bearer")
        return // don't do anything else if it was the option method
    }
    bearer := request.Header.Get("Bearer")

    type JwtBearer struct {
        Userjwt jwt
    }
    var jwtBearer JwtBearer
    json.Unmarshal([]byte(bearer),&jwtBearer)
    user, err := getUserFromToken(jwtBearer.Userjwt)
    if err != nil{
        panic(err)
    }
    tweets := DatabaseGetTweetsFromFollowers(user)
    tweetjson, err := json.Marshal(tweets)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(tweetjson))
    fmt.Fprint(writer,tweetjson)
}
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
    fmt.Println("trying to log in")
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

    user, err := getUserFromToken(getJwtBearer(request))
    if err != nil{
        panic(err)
    }

    tweets := DatabaseGetTweetsFromFollowers(user)
    tweetjson, err := json.Marshal(tweets)
    if err != nil {
        panic(err)
    }
    fmt.Fprint(writer,string(tweetjson))
}

func SendTweet(writer http.ResponseWriter, request *http.Request){
    if(request.Method == "OPTIONS"){
        writer.Header().Add("Acces-Control-Allow-Headers","Bearer")
        return
    }
    fmt.Println("Sending tweet!")
    user, err := getUserFromToken(getJwtBearer(request))
    if err != nil{
        panic(err)
    }

    err = request.ParseForm()
    if err != nil {
        panic(err)
    }

    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        panic(err)
    }

    type ContentStruct struct{
        Content string
    }
    var contentStruct ContentStruct
    json.Unmarshal(body,&contentStruct)
    DatabaseSendTweet(Tweet{Username:user.Username,Tweetbody:contentStruct.Content})
}

func getJwtBearer(request *http.Request) jwt{
    bearer := request.Header.Get("Bearer")

    type JwtBearer struct {
        Userjwt jwt
    }
    var jwtBearer JwtBearer
    json.Unmarshal([]byte(bearer),&jwtBearer)
    return jwtBearer.Userjwt
}
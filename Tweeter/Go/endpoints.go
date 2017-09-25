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

    type Tokenuser struct {
        Username string
        Token jwt
    }

    tokenuser := Tokenuser{Username:user.Username, Token: token}
    loginjson, err := json.Marshal(tokenuser)
    if err != nil{
        panic(err)
    }
    fmt.Fprint(writer, string(loginjson))
}

func ProfileTweets(writer http.ResponseWriter, request *http.Request){
    if(request.Method == "OPTIONS"){
        writer.Header().Add("Access-Control-Allow-Headers", "Username")
        return // don't do anything else if it was the option method
    }

    username := getUsernameFromRequest(request)
    user, err := getDatabaseUserByName(username)
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
    newTweet := DatabaseSendTweet(Tweet{Username:user.Username,Tweetbody:contentStruct.Content})
    jsonString, jsonErr := json.Marshal(newTweet)
    if jsonErr != nil {
        panic(jsonErr)
    }
    fmt.Fprint(writer, string(jsonString))
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

func getUsernameFromRequest(request *http.Request) string{
    username := request.Header.Get("Username")
    type usernameHeader struct{
        Username string
    }

    var usernameheader usernameHeader
    json.Unmarshal([]byte(username), &usernameheader)
    return usernameheader.Username
}
package main

import (
    "fmt"
    "encoding/json"
    "encoding/base64"
    "strings"
    "errors"
)

var secret string = "supersecret"

func generateToken(user User) jwt {
    // I need to base 64 encode user
    jsonuser, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonuser))

    type JWTHeader struct {
        Alg string
        Typ string
    }


    base64User := make([]byte, base64.StdEncoding.EncodedLen(len(jsonuser)))
    base64.StdEncoding.Encode(base64User, []byte(jsonuser))
    userText := string(base64User)


    jwtHeader := JWTHeader{Alg:"HS256", Typ:"JWT"}
    jsonjwtinfo, err := json.Marshal(jwtHeader)
    if err != nil {
        panic(err)
    }

    base64Header := make([]byte, base64.StdEncoding.EncodedLen(len(jsonjwtinfo)))
    base64.StdEncoding.Encode(base64Header, []byte(jsonjwtinfo))
    headerString := string(base64Header)

    concat := strings.Join([]string {headerString,userText,secret}, ".")
    return jwt(concat)
}

func getUserFromToken(token jwt) (User, error){
    parts := strings.Split(string(token),".")
    fmt.Println(parts)
    //headerinfo := parts[0] -- for now this is not important
    message := parts[1]
    secretPart := parts[2]

    var user User
    if strings.Compare(secret,secretPart) != 0{
        return user, errors.New("Secret does not match!")
    }

    // the secret matches
    data, err := base64.StdEncoding.DecodeString(message)
    if err != nil{
        panic(err)
    }
    json.Unmarshal(data,&user)
    return user, nil
}
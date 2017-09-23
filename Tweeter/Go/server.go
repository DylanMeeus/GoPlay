package main


import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)


func main() {
    SetupDatabase()
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/", GetUsers).Methods("GET")
    router.HandleFunc("/tweets", GetTweets).Methods("GET")
    router.HandleFunc("/login", Login).Methods("POST", "OPTIONS")
    router.HandleFunc("/profile/tweets", ProfileTweets).Methods("GET","OPTIONS")
    router.HandleFunc("/profile/sendtweet", SendTweet).Methods("POST", "OPTIONS")

    c := cors.New(cors.Options{
        AllowedMethods: []string{"GET","POST", "OPTIONS"},
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
        AllowedHeaders: []string{"Content-Type","Bearer" ,"content-type","Origin","Accept"},
        OptionsPassthrough: true,
    })

    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":8080", handler))
}

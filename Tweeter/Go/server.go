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
    router.HandleFunc("/users", GetUsers)
    router.HandleFunc("/tweets", GetTweets)
    handler := cors.Default().Handler(router)
    log.Fatal(http.ListenAndServe(":8080", handler))
}

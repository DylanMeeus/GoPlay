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
    router.HandleFunc("/tweets", GetTweets).Methods("GET")
    router.HandleFunc("/login", Login).Methods("POST")
    handler := cors.Default().Handler(router)
    log.Fatal(http.ListenAndServe(":8080", handler))
}

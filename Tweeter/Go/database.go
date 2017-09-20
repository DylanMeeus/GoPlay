package main


// file for the database stuff

import (
    "os"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func SetupDatabase(){
    os.Remove("./tweeterdb.db")
    db, err := sql.Open("sqlite3", "./tweeterdb.db")
    if err != nil{
        panic(err)
    }
    defer db.Close();


    statement := `
        create table users(id integer not null primary key, name text, password text);
    `
    _, err = db.Exec(statement)
    if err != nil{
        fmt.Println(err)
    }

    statement = `
    create table tweets(id integer not null primary key, userid integer, tweet text);
    `

    _, err = db.Exec(statement)
    if err != nil {
        panic(err)
    }

    fmt.Println("database setup done")
    populateDatabase(db)
}

// populate the db with some test data
func populateDatabase(db *sql.DB){
    fmt.Println("Populating database..")
    statement := `
        insert into users(name,password) values ("Dylan","Password");
        insert into users(name,password) values ("Ana","Password");
    `
    _, err := db.Exec(statement)
    if err != nil {
        panic(err)
    }


    statement = `
        insert into tweets(userid, tweet) values (1, "Hello world, this is my first tweet!");
        insert into tweets(userid, tweet) values (2, "Hola mundo, esto es mi primero tweet!");
        insert into tweets(userid, tweet) values (1, "Go is a pretty great language!");
    `

    _, err = db.Exec(statement)
    if err != nil{
        panic(err)
    }
}


func DatabaseTweets() Tweets{
    var tweets Tweets
    db, err := sql.Open("sqlite3", "./tweeterdb.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    querystring := `
        select name, tweet from users inner join tweets on users.id = tweets.userid;
    `

    rows, err := db.Query(querystring)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        var tweet string
        err = rows.Scan(&name, &tweet)
        if err != nil {
            fmt.Println(err)
        }
        newTweet := Tweet{Username:name, Tweetbody:tweet}
        tweets = append(tweets, newTweet)
    }
    return tweets;
}


func getDatabaseUsers() []string{
    db, err := sql.Open("sqlite3", "./tweeterdb.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    querystring := `
        select name from users;
    `

    rows, err := db.Query(querystring)
    if err != nil {
        panic(err)
    }
    defer rows.Close()


    var usernames []string
    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println(err)
        }
        usernames = append(usernames,name)
    }
    return usernames
}



































package main


// file for the database stuff

import (
    "os"
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "errors"
    "time"
)

type jwt string

func SetupDatabase(){
    fmt.Println("Cleaning up last database.")
    os.Remove("./tweeterdb.db")
    db := openDatabase()
    defer db.Close();

    fmt.Println("Creating database tables.")
    statement := `
        create table users(id integer not null primary key, name text unique, password text);
    `
    _, err := db.Exec(statement)
    if err != nil{
        fmt.Println(err)
    }

    statement = `
    create table tweets(id integer not null primary key, userid integer, tweet varchar(140), sendtime datetime not null);
    `

    _, err = db.Exec(statement)
    if err != nil {
        panic(err)
    }

    statement = `
        create table followers(userid integer, followinguserid integer);
    `

    _, err = db.Exec(statement)
    if err != nil {
        panic(err)
    }
    fmt.Println("Database tables created.")
    populateDatabase(db)
}

// populate the db with some test data
func populateDatabase(db *sql.DB){
    fmt.Println("Populating database..")
    statement := `
        insert into users(name,password) values ("Dylan","Password");
        insert into users(name,password) values ("Ana","Password");
        insert into users(name,password) values ("Sean","Password");
        insert into users(name,password) values ("Sander","Password");
        insert into users(name,password) values ("Dexter","Password");
        insert into users(name,password) values ("Luna","Password");
        insert into users(name,password) values ("Chris","Password");
    `
    _, err := db.Exec(statement)
    if err != nil {
        panic(err)
    }


    statement = `
        insert into tweets(userid, tweet, sendtime) values (1, "Hello world, this is my first tweet!", "2017-23-09 10:00:00");
        insert into tweets(userid, tweet, sendtime) values (2, "Hola mundo, esto es mi primero tweet!", "2017-22-09 10:00:00");
        insert into tweets(userid, tweet, sendtime) values (1, "Go is a pretty great language!", "2017-21-09 10:00:00");
        insert into tweets(userid, tweet, sendtime) values (3, "Lorem ipsum dolor sit amet!", "2017-20-09 10:00:00");
        insert into tweets(userid, tweet, sendtime) values (5, "I did not like this feeling of having feelings.", "2017-20-09 10:00:00");
        insert into tweets(userid, tweet, sendtime) values (5, "The mind picks some very bad times to take a walk, doesn't it?.", "2017-20-09 09:00:00");
    `

    _, err = db.Exec(statement)
    if err != nil{
        panic(err)
    }

    statement = `
        insert into followers(userid, followinguserid) values (1,2);
        insert into followers(userid, followinguserid) values (1,3);
        insert into followers(userid, followinguserid) values (1,5);
        insert into followers(userid, followinguserid) values (2,1);
    `

    _, err = db.Exec(statement)
    if err != nil {
        panic(err)
    }
    fmt.Println("Database populated.")
}


func openDatabase() *sql.DB {
    db, err := sql.Open("sqlite3","./tweeterdb.db")
    if err != nil {
        panic(err)
    }
    return db
}

func DatabaseLogin(user User) jwt {
    username := user.Username
    password := user.Password
    db := openDatabase()
    defer db.Close()
    rows, err := db.Query("select * from users where name = ? and password = ?",username,password)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    hasItems := false
    for rows.Next() {
        var id int
        var name string
        var password string
        rows.Scan(&id, &name, &password)
        fmt.Println(name + " " + password)
        user = User{Id:id, Username:name,Password:password}
        hasItems = true
        break
    }
    if hasItems {
        return generateToken(user)
    }
    return "failed"
}

func DatabaseTweets() Tweets{
    var tweets Tweets
    db := openDatabase()
    defer db.Close()

    querystring := `
        select name, tweet, sendtime from users inner join tweets on users.id = tweets.userid
        order by sendtime asc;
    `

    rows, err := db.Query(querystring)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        var tweet string
        var sendtime time.Time
        err = rows.Scan(&name, &tweet, &sendtime)
        if err != nil {
            fmt.Println(err)
        }
        newTweet := Tweet{Username:name, Tweetbody:tweet}
        tweets = append(tweets, newTweet)
    }
    return tweets;
}


func getDatabaseUsers() []string{
    db := openDatabase()
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


// get the tweets from the followers
func DatabaseGetTweetsFromFollowers(user User) Tweets{
    db := openDatabase()

    statement := `
        select name, tweet, sendtime from users inner join tweets on users.id = tweets.userid where userid in (
            select followinguserid from followers where userid = ?
        ) or tweets.userid = ?
        order by sendtime asc
    `

    rows, err := db.Query(statement, user.Id, user.Id)
    if err != nil {
        panic(err)
    }

    var tweets Tweets
    for rows.Next(){
        var name string
        var tweet string
        var sendtime time.Time
        err := rows.Scan(&name,&tweet, &sendtime)
        if err != nil {
            panic(err)
        }
        newTweet := Tweet{Username:name, Tweetbody:tweet}
        tweets = append(tweets,newTweet)
    }
    return tweets
}


// returns true if it was successful
func DatabaseSendTweet(tweet Tweet) Tweet{
    fmt.Println("Sending a tweet!")
    user, err := getDatabaseUserByName(tweet.Username)
    if err != nil{
        panic(err)
    }
    db := openDatabase()
    statement := "insert into tweets(userid, tweet, sendtime) values (?,?,?)"
    // set the time for the tweet
    tweet.Sendtime = time.Now()
    _, err = db.Exec(statement,user.Id,tweet.Tweetbody, tweet.Sendtime)
    if err != nil {
        panic(err)
    }
    return tweet
}


func getDatabaseUserByName(username string) (User,error){
    db := openDatabase()
    statement := "select * from users where name = ?"
    rows, err := db.Query(statement, username)
    if err != nil{
        panic(err)
    }
    defer rows.Close()
    hasItems := false

    var user User
    for rows.Next() {
        var id int
        var name string
        var password string
        rows.Scan(&id, &name, &password)
        user = User{Id:id, Username:name,Password:password}
        hasItems = true
        break
    }
    if hasItems {
        return user, nil
    }
    return user, errors.New("user not found")
}
































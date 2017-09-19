package main

type Tweet struct{
    Username    string     `json:"username"`
    Tweetbody   string    `json:"body"`
}

type Tweets []Tweet

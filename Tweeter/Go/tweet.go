package main

import "time"

type Tweet struct{
    Username    string     `json:"username"`
    Tweetbody   string    `json:"body"`
    Sendtime    time.Time        `json:sendtime`
}

type Tweets []Tweet

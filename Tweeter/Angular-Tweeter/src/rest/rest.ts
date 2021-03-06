/**
 * Created by dylan on 19.09.17.
 */

import { Tweet } from '../model/tweet'


export default class Rest {

  baseurl : string = "http://localhost:8080/"

  constructor(){
  }

  public async getTweets() {
    var url : string = this.baseurl + "tweets"
    var tweets : Tweet[] = [];
     return fetch(url).then(
      function(response){
          return response.text().then(
          function(body){
            var jsonObj = JSON.parse(body);
            for(var i = 0; i < jsonObj.length; i++) {
              var tweet : Tweet = new Tweet();
              tweet.username = jsonObj[i].username;
              tweet.tweetbody = jsonObj[i].body;
              tweets.push(tweet);
            }
            return tweets;
          }
        )}
    )
  }

  public async login(username : string, password : string){
    var url : string = this.baseurl + "login";
    var jsonstring = JSON.stringify({username,password});
    return fetch(url, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: "POST",
      body: jsonstring
    }).then(
      function(response){
        return response.text().then((text) =>{
          return text;
        });
      }
    )
  }

  public async loadFollowerTweets(username: string){
    var url : string = this.baseurl + "profile/tweets";
    var jsonusername = JSON.stringify({username});
    console.log("loading follower tweets");
    console.log(jsonusername);
    return fetch(url, {
      headers: {
        "Content-Type": 'application/text',
        "Username": jsonusername
      },
      method: "GET"
    }).then((response) =>{
      return response.text().then((text) => {
        var jsonObj = JSON.parse(text);
        var tweets : Tweet[] = []
        for(var i = 0; i < jsonObj.length; i++) {
          var tweet : Tweet = new Tweet();
          tweet.username = jsonObj[i].username;
          tweet.tweetbody = jsonObj[i].body;
          tweets.push(tweet);
        }
        return tweets;
      })
    });
  }

  // Send a tweet for a certain user with some content
  public async sendTweet(userjwt: string, content: string){
    if(content == ""){
      return; // no point in sending an empty tweet
    }

    var url : string = this.baseurl + "profile/sendtweet";
    var bearer = JSON.stringify({userjwt});
    var contentjson = JSON.stringify({content});
    return fetch(url, {
      headers: {
        "Content-Type": 'application/text',
        "Bearer" : bearer
      },
      method: "POST",
      body: contentjson
    }).then((response) => {
      return response.text().then((text) => {
        var jsonObj = JSON.parse(text);
        var tweet : Tweet = new Tweet();
        tweet.username = jsonObj.username;
        tweet.tweetbody = jsonObj.body;
        return tweet
      })
    })
  }
}

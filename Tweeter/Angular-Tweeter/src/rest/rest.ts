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

  public async loadFollowerTweets(userjwt: string){
    var url : string = this.baseurl + "profile/tweets";
    var bearer = JSON.stringify({userjwt});
    return fetch(url, {
      headers: {
        "Content-Type": 'application/text',
        "Bearer": bearer
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
}

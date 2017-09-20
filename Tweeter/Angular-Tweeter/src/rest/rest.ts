/**
 * Created by dylan on 19.09.17.
 */

import { Tweet } from '../model/tweet'


export default class Rest {

  baseurl : string = "localhost:8080/"

  constructor(){
  }

  public async getTweets() {
    var url : string = this.baseurl + "users"
    var tweets : Tweet[] = [];
     return fetch("http://localhost:8080/tweets").then(
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
}

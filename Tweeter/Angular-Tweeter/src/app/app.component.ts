import {Component, NgModule} from '@angular/core';
import Rest from '../rest/rest'
import {Tweet} from "../model/tweet";


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'Tweeter';
  users: string[] = ["one", "two"];
  rest : Rest;
  tweets : Tweet[];
  constructor(){
    this.rest = new Rest();
    console.log("created rest service");
    this.rest.getTweets().then( (y) =>{
      console.log(y);
      this.tweets = y;
      for(var i = 0; i < this.tweets.length; i++){
        console.log("body: " + this.tweets[i].tweetbody);
      }
    });
  }

}

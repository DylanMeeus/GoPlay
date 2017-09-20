import { Component, OnInit } from '@angular/core';
import Rest from "../../rest/rest";
import {Tweet} from "../../model/tweet";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  title = 'Tweeter';
  users: string[] = ["one", "two"];
  rest : Rest;
  tweets : Tweet[];

  constructor(){

  }

  ngOnInit() {
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

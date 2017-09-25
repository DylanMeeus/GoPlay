import {Component, OnInit, NgModule, Input} from '@angular/core';
import {Tweet} from "../../../model/tweet";

@Component({
  selector: 'tweetview',
  templateUrl: './tweetview.component.html',
  styleUrls: ['./tweetview.component.css']
})
export class TweetviewComponent implements OnInit {

  constructor() { }

  @Input() tweets : Tweet[]

  ngOnInit() {
  }

  viewReplies(event){
    console.log(event);
  }
}

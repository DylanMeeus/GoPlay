import {Component, OnInit, NgModule} from '@angular/core';
import {Router} from "@angular/router";
import Rest from "../../../rest/rest";
import {Tweet} from "../../../model/tweet";
import {TweetviewComponent} from "../../tweets/tweetview/tweetview.component";
import {FormControl, FormBuilder} from "@angular/forms";


@NgModule({
  imports: [TweetviewComponent]
})

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css'],
})
export class ProfileComponent implements OnInit {

  chartotal: number = 140;
  charleft: number = 140;
  tweets : Tweet[];

  public tweetForm = this.fb.group({
    tweetcontent: new FormControl(""),
  });

  constructor(public fb: FormBuilder, public router: Router) {
    this.tweetForm.valueChanges.subscribe( change =>{
      var text = change.tweetcontent;
      this.charleft = this.chartotal - text.length;
      if(this.charleft < 0){
        this.charleft = 0;
        text = text.substring(0,this.chartotal);
        this.tweetForm.controls['tweetcontent'].setValue(text);
      }
    });
  }

  ngOnInit() {
    var userjwt = localStorage.getItem("userjwt");
    if(userjwt == null){
      this.router.navigate(["user/login"]);
      return;
    }

    var rest : Rest = new Rest();
    rest.loadFollowerTweets(userjwt).then((result) => {
      this.tweets = result;
    });
  }

  sendTweet(event){
    var tweet = this.tweetForm.controls['tweetcontent'].value;
    var jwt = localStorage.getItem("userjwt");
    var rest : Rest = new Rest();
    rest.sendTweet(jwt,tweet).then( (result) => {
      console.log(result);
    })
  }

}

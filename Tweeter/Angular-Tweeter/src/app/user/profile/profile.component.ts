import {Component, OnInit, NgModule} from '@angular/core';
import {Router, ActivatedRoute} from "@angular/router";
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
  ownprofile: boolean = false;
  public tweetForm = this.fb.group({
    tweetcontent: new FormControl(""),
  });

  constructor(public fb: FormBuilder, public router: Router, public route: ActivatedRoute) {
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
    this.route.params.subscribe(params =>{
      // I need to get the current user profile..
      var username = params['user'];
      if(username == null){
        var localUsername = localStorage.getItem("username");
        if(localUsername == null){
          this.router.navigate(['/login']);
        }
        this.loadProfile(localUsername);
        return;
      }
      var localstorageUsername = localStorage.getItem("username");
      if(localstorageUsername == username){
        this.ownprofile = true;
      }
      this.loadProfile(username);
    });
  }

  loadProfile(username: string){
    var rest : Rest = new Rest();
    rest.loadFollowerTweets(username).then((result) => {
      this.tweets = result;
    });
  }

  sendTweet(event){
    var tweet = this.tweetForm.controls['tweetcontent'].value;
    var jwt = localStorage.getItem("userjwt");
    var rest : Rest = new Rest();
    rest.sendTweet(jwt,tweet).then( (result) => {
        // add this tweet to the 'tweets'
        var newarr = [];
        newarr.push(result); // this one should be in the front
        for(var i = 0; i < this.tweets.length; i++){
          newarr.push(this.tweets[i]);
        }
        this.tweets = newarr;
    });
  }
}

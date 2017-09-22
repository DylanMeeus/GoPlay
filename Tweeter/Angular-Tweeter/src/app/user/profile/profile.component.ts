import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Rest from "../../../rest/rest";
import {Tweet} from "../../../model/tweet";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  constructor(public router: Router) { }

  tweets : Tweet[]
  ngOnInit() {
    var userjwt = localStorage.getItem("userjwt");
    if(userjwt == null){
      this.router.navigate(["user/login"])
      return
    }

    var rest : Rest = new Rest();
    rest.loadFollowerTweets(userjwt).then((result) => {
      this.tweets = result
    })
  }

}

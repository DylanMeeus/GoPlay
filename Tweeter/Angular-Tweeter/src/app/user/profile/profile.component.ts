import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Rest from "../../../rest/rest";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  constructor(public router: Router) { }

  ngOnInit() {
    var userjwt = localStorage.getItem("userjwt");
    if(userjwt == null){
      this.router.navigate(["user/login"])
      return
    }

    // if we get here, we (normally) have a user JWT
    console.log(userjwt)
    var rest : Rest = new Rest();
    rest.loadFollowerTweets(userjwt).then((result) => {
      // pass
    })
  }

}

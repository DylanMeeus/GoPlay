import {Component, OnInit, NgModule} from '@angular/core';
import {FormControl, FormGroup, FormBuilder, ReactiveFormsModule} from "@angular/forms";
import Rest from "../../../rest/rest";
import {Router} from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})


export class LoginComponent implements OnInit {

  public loginForm = this.fb.group({
    username :new FormControl("username"),
    password :  new FormControl("password")
  });

  constructor(public fb: FormBuilder, public router: Router) { }

  ngOnInit() {
  }

  doLogin(event){
    var username = this.loginForm.controls.username.value;
    var password = this.loginForm.controls.password.value;
    var rest : Rest = new Rest();
    rest.login(username, password).then((token) => {
      // store the token in local storage
      if(token != "failed"){
        localStorage.setItem("userjwt",token);
        this.router.navigate(["/user/profile"])
      }
    });
  }
}

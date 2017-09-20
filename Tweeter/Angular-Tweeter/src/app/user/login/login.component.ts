import {Component, OnInit, NgModule} from '@angular/core';
import {FormControl, FormGroup, FormBuilder, ReactiveFormsModule} from "@angular/forms";

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

  constructor(public fb: FormBuilder) { }

  ngOnInit() {
  }

  doLogin(event){
    var username = this.loginForm.controls.username.value;
    var password = this.loginForm.controls.password.value;
    console.log(username);
    console.log(password);
  }
}

import {Component, NgModule} from '@angular/core';
import Rest from '../rest/rest'



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'Tweeter';
  users: string[] = ["one", "two"];
  rest : Rest;
  constructor(){
    this.rest = new Rest();
    console.log("created rest service");
    this.rest.getUsers();
  }

}

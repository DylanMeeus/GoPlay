import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RouterModule, Routes } from '@angular/router'

import { AppComponent } from './app.component';
import { TimelineComponent } from './user/timeline/timeline.component';
import { ProfileComponent } from './user/profile/profile.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './user/login/login.component';

import { ReactiveFormsModule } from '@angular/forms';
import { LogoutComponent } from './user/logout/logout.component'

const appRoutes: Routes = [
  { path: 'user/profile', component: ProfileComponent },
  { path: 'user/timeline', component: TimelineComponent },
  { path: 'user/login', component: LoginComponent },
  { path: 'user/logout', component: LogoutComponent },
  { path: '**', component: HomeComponent}
];


@NgModule({
  declarations: [
    AppComponent,
    TimelineComponent,
    ProfileComponent,
    HomeComponent,
    LoginComponent,
    LogoutComponent
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true} // debug
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

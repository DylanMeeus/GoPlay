import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RouterModule, Routes } from '@angular/router'

import { AppComponent } from './app.component';
import { TimelineComponent } from './user/timeline/timeline.component';
import { ProfileComponent } from './user/profile/profile.component';
import { HomeComponent } from './home/home.component';


const appRoutes: Routes = [
  { path: 'user/profile', component: ProfileComponent },
  { path: 'user/timeline',      component: TimelineComponent },
  { path: '**', component: HomeComponent}
];


@NgModule({
  declarations: [
    AppComponent,
    TimelineComponent,
    ProfileComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true} // debug
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

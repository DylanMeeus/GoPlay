import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TweetviewComponent } from './tweetview.component';

describe('TweetviewComponent', () => {
  let component: TweetviewComponent;
  let fixture: ComponentFixture<TweetviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TweetviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TweetviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

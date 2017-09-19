/**
 * Created by dylan on 19.09.17.
 */


export default class Tweet{

  private _username : string;
  private _tweetbody : string;

  constructor(json?: string){

  }

  set username(value: string){
    this._username = value;
  }

  set tweetbody(value: string){
    this._tweetbody = value;
  }

  get username(){
    return this._username;
  }

  get tweetbody(){
    return this._username;
  }

}

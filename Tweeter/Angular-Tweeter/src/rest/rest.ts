/**
 * Created by dylan on 19.09.17.
 */


export default class Rest {

  baseurl : string = "localhost:8080/"

  constructor(){
  }

  public getUsers(): string {

    var url : string = this.baseurl + "users"

    console.log("setting up fetch")
    fetch("http://localhost:8080/tweets").then(
      function(response){
        response.text().then(
          function(body){
            console.log(body);
            var jsonObj = JSON.parse(body);
            console.log(jsonObj);
          }
        )
      }
    )

    console.log("done with the call");
    return ""
  }


}

// A simple file for serving HTML files for different URLs
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

//The code that will be run
func main(){
  //Returns a file server that returns a a handler that server HTTP requests with the contents of the file system
  fs := http.FileServer(http.Dir("Templates"))
  //Handles the handler for the given pattern
  //Strip prefix strips the input from the URL pattern

  http.Handle("/Templates", http.StripPrefix("Templates", fs))
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
  http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("./JS"))))

  //The top lines will serve files from the Templates directory in this project
  //The top lines will serve files from the css directory in this project
  //The top lines will serve files from the JS directory in this project

  //This function enables the getIndex function for the "/" URL
  http.HandleFunc("/", getIndex)
  //This function enables the getGamePage function for the "/game" URL
  http.HandleFunc("/game", getGamePage)
  http.HandleFunc("/words", returnWords)
  http.HandleFunc("/lorem_ipsum", getLoremIpsum)

  fmt.Println("Server is running on localhost:8080")

  //Start the http server on localhost:8080
  http.ListenAndServe(":8080", nil)

}

//Function for returning the "Index.html" page
func getIndex(w http.ResponseWriter, r *http.Request){
  //template = Index.html, err = potential errors
  template, err := template.ParseFiles(filepath.Join("Templates", "Index.html"))
  //Handles the error
  if (err != nil){
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  //Serves the html files
  template.Execute(w, nil)
}

//Function for returning the "Game.html" page
func getGamePage(w http.ResponseWriter, r *http.Request){
  //template = Game.html, err = potential errors
  template, err := template.ParseFiles(filepath.Join("Templates", "Game.html"))
  //Handles the error
  if (err != nil){
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  //Serves the html files
  template.Execute(w, nil)
}

func getLoremIpsum(w http.ResponseWriter, r *http.Request){

  lorem_ipsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer auctor est eget nulla sagittis luctus. Duis turpis erat, vulputate sit amet luctus ut, mollis vitae odio. Mauris eget elit lacus. Nam nisl nulla, iaculis vitae velit eget, tristique condimentum lectus. Suspendisse tincidunt tempor rhoncus. Mauris commodo gravida consectetur. Praesent nec orci non erat condimentum ultrices. Quisque volutpat, augue in fermentum egestas, nibh erat aliquam enim, convallis faucibus ligula augue ac purus. Praesent vitae ex est. Maecenas fringilla purus ante, sed sagittis urna sollicitudin ac. Suspendisse tincidunt pretium enim quis auctor. Donec rhoncus pellentesque neque, nec venenatis dui iaculis et. Vestibulum lobortis accumsan consequat. Mauris tincidunt vitae nisl at ultricies. Quisque sed nisi eget nunc tincidunt malesuada. Praesent finibus, justo nec dapibus sagittis, urna sapien feugiat purus, id malesuada diam est eget nisl." 


  w.Header().Set("Content-type", "Application/json")
  if err := json.NewEncoder(w).Encode(lorem_ipsum); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }

}

func returnWords(w http.ResponseWriter, r *http.Request){
  words := []string{ "Apple", "Banana", "Cherry" }

  w.Header().Set("Content-type", "Application/json");

  if err := json.NewEncoder(w).Encode(words); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }
}

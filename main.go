// A simple file for serving HTML files for different URLs
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
	d "typinggame/internal"
	c "typinggame/multiplayer_backend"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var WordRepo d.WordRepo = d.NewWordRepo()
var store = sessions.NewCookieStore([]byte("secret"))

//The code that will be run
func main(){
  dao := d.NewUserDAO()
  dao.TestConn()
  server := c.NewGameServer()
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
  http.HandleFunc("/MPgame", getMPGamePage)
  http.HandleFunc("/multiplayer", getMultiplayerPage)
  http.HandleFunc("/words", returnWords)
  http.HandleFunc("/lorem_ipsum", getLoremIpsum)
  http.HandleFunc("/ten_words", getTenWords)
  http.HandleFunc("/ws", server.HandleConnections)
  http.HandleFunc("/lobby", getLobbyPage)
  http.HandleFunc("/join", joinLobbyHandler)
  http.HandleFunc("/joinGame", joinGameHandler)
  http.HandleFunc("/Login", getLoginPage)

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

func getLoginPage(w http.ResponseWriter, r *http.Request){
  template, err := template.ParseFiles(filepath.Join("Templates", "LoginPage.html"))
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  template.Execute(w, nil)
}


func getLobbyPage(w http.ResponseWriter, r *http.Request){
  template, err := template.ParseFiles(filepath.Join("Templates", "LobbyJoiner.html"))
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  template.Execute(w, nil)
}

//Function for returning the "Game.html" page
func getMPGamePage(w http.ResponseWriter, r *http.Request){
  //template = Game.html, err = potential errors
  template, err := template.ParseFiles(filepath.Join("Templates", "MPGame.html"))
  //Handles the error
  if (err != nil){
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  //Serves the html files
  template.Execute(w, nil)
}
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

func getMultiplayerPage(w http.ResponseWriter, r *http.Request){
  lobbyID := r.URL.Query().Get("lobbyid")
  clientID := r.URL.Query().Get("clientid")

  if lobbyID == "" {
    http.Error(w, "Lobby ID is required", http.StatusBadRequest)
    return
  }

  // Create the data structure to pass to the template
  data := PageData{
    LobbyID:  lobbyID,
    ClientID: clientID,
  }

  // Parse the template
  tmpl, err := template.ParseFiles(filepath.Join("Templates", "Lobby.html"))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  // Execute the template with the data
  if err := tmpl.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}


func getLoremIpsum(w http.ResponseWriter, r *http.Request){

  lorem_ipsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer auctor est eget nulla sagittis luctus. Duis turpis erat, vulputate sit amet luctus ut, mollis vitae odio. Mauris eget elit lacus. Nam nisl nulla, iaculis vitae velit eget, tristique condimentum lectus. Suspendisse tincidunt tempor rhoncus. Mauris commodo gravida consectetur. Praesent nec orci non erat condimentum ultrices. Quisque volutpat, augue in fermentum egestas, nibh erat aliquam enim, convallis faucibus ligula augue ac purus. Praesent vitae ex est. Maecenas fringilla purus ante, sed sagittis urna sollicitudin ac. Suspendisse tincidunt pretium enim quis auctor. Donec rhoncus pellentesque neque, nec venenatis dui iaculis et. Vestibulum lobortis accumsan consequat. Mauris tincidunt vitae nisl at ultricies. Quisque sed nisi eget nunc tincidunt malesuada. Praesent finibus, justo nec dapibus sagittis, urna sapien feugiat purus, id malesuada diam est eget nisl."
  w.Header().Set("Content-type", "Application/json")
  if err := json.NewEncoder(w).Encode(lorem_ipsum); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }
}

func getTenWords(w http.ResponseWriter, r *http.Request){

  session, _ := store.Get(r, "session-name")
  lobbyIDInterface := session.Values["lobbyid"]
  lobbyID, ok := lobbyIDInterface.(string)
  if !ok || lobbyID == ""{
    http.Error(w, "Lobby id is required", http.StatusInternalServerError);
  }
  lobbyidnr, err := strconv.Atoi(lobbyID)
  if err != nil{
    http.Error(w, "Lobby ID is required", http.StatusBadRequest)
  }
  words, err := WordRepo.GetTenWords(int64(lobbyidnr))
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }
  w.Header().Set("Content-type", "Application/json");

  if err := json.NewEncoder(w).Encode(words); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }
}
func returnWords(w http.ResponseWriter, r *http.Request){ words := []string{ "Apple", "Banana", "Cherry" }

  w.Header().Set("Content-type", "Application/json");

  if err := json.NewEncoder(w).Encode(words); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError);
  }
}

type PageData struct {
    LobbyID string
    ClientID string
}

func joinLobbyHandler(w http.ResponseWriter, r* http.Request){
  if r.Method == http.MethodPost {
    lobbyID := r.FormValue("lobbyid")
    clientid := GenerateClientID()
    if lobbyID == ""{
      http.Error(w, "Lobby id is requirede", http.StatusBadRequest)
      return
    }
    if clientid == ""{
      http.Error(w, "Lobby id is requirede", http.StatusBadRequest)
      return
    }
    session, _ := store.Get(r, "session-name")
    session.Values["lobbyid"] = lobbyID
    // session.Values["clientid"] = clientid
    session.Save(r,w)
    http.Redirect(w, r, "/multiplayer?lobbyid="+lobbyID+"&clientid="+clientid, http.StatusSeeOther)
    return
  }
  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func joinGameHandler(w http.ResponseWriter, r* http.Request){
  if r.Method == http.MethodPost {
    session, _ := store.Get(r, "session-name")
    lobbyIDInterface := session.Values["lobbyid"]
    lobbyID, ok := lobbyIDInterface.(string)
    if !ok || lobbyID == ""{
      http.Error(w, "Lobby id is required", http.StatusBadRequest)
      return
    }
    clientid := r.FormValue("clientid")
    fmt.Println(clientid)
    if clientid == ""{
      http.Error(w, "Client id is required", http.StatusBadRequest)
      return
    }
    http.Redirect(w, r, "/MPgame?lobbyid="+lobbyID+"&clientid="+clientid, http.StatusSeeOther)
    return
  }
  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// GenerateClientID generates a random client ID
func GenerateClientID() string {
  rand.Seed(time.Now().UnixNano())
  letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
  b := make([]rune, 10)
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

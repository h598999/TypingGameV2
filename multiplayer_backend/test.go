package multiplayerbackend

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool{
    return true
  },
}

type Client struct {
  id int64
  conn *websocket.Conn
  lobby *Lobby
  // Add other fields as needed, e.g., username, current progress, etc.
}

type Message struct{
  ID string `json:"id"`
  WPM string `json:"wpm"`
  CURSORPOSITION string `json:"cursor"`
}

type Lobby struct {
  clients map[*Client]bool
  broadcast chan []byte
  register chan *Client
  unregister chan *Client
  mutex sync.Mutex
}

type GameServer struct{
  idGen int64
  lobbies map[string] *Lobby
  mutex sync.Mutex
}

func NewGameServer() *GameServer{
  return &GameServer{
    idGen: 0,
    lobbies: make(map[string]*Lobby),
  }
}

func newLobby() *Lobby{
  return &Lobby{
    clients: make(map[*Client]bool),
    broadcast: make(chan []byte),
    register: make(chan *Client),
    unregister: make(chan *Client),
  }
}

func (lobby *Lobby) run(){
  fmt.Println("Lobby started running")
  for {
    select{
    case client := <-lobby.register:
      lobby.mutex.Lock()
      lobby.clients[client] = true
      lobby.mutex.Unlock()
    case client := <-lobby.unregister:
      lobby.mutex.Lock()
      if _, ok := lobby.clients[client]; ok {
        delete(lobby.clients, client)
        client.conn.Close()
      }
      lobby.mutex.Unlock()

    case message := <- lobby.broadcast:
      lobby.mutex.Lock()
      for client := range lobby.clients{
        err := client.conn.WriteMessage(websocket.TextMessage, message)
        if err != nil{
          log.Print("Error sending message: ", err)
          client.conn.Close()
          delete(lobby.clients, client)
        }
      }
      lobby.mutex.Unlock()

    }
  }
}

func (server *GameServer) getLobby(id string) *Lobby{
  server.mutex.Lock()
  defer server.mutex.Unlock()
  if lobby, ok := server.lobbies[id]; ok{
    return lobby
  }
  lobby := newLobby()
  server.lobbies[id] = lobby
  go lobby.run()
  return lobby
}

func (server *GameServer)HandleConnections(w http.ResponseWriter, r *http.Request){
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil{
    log.Printf("Error: %v", err)
    return
  }
  defer conn.Close()

   var initMessage struct {
     ID string `json:"id"`
     LobbyID string `json:"lobbyid"`
    }

  err = conn.ReadJSON(&initMessage)
  if err != nil{
    log.Printf("Error reading initial message: %v", err)
    return
  }

  lobby := server.getLobby(initMessage.LobbyID)
  if lobby == nil {
    log.Printf("Lobby with ID %s not found", initMessage.LobbyID)
    return
  }
  client := &Client{conn: conn, lobby: lobby, }
  lobby.register <- client

  defer func(){
    lobby.unregister <- client
  }()

  for {
    _, message, err := conn.ReadMessage()
    if err != nil{
      log.Printf("Error: %v", err)
      break
    }
    lobby.broadcast <- message
  }
}

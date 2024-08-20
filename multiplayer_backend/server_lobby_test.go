package multiplayerbackend

import (
	"net/http"
  "fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var server *GameServer

func setup() *httptest.Server {
	server = NewGameServer()
	httpServer := httptest.NewServer(http.HandlerFunc(server.HandleConnections))
  fmt.Println("Setting up")
	return httpServer
}

func TestWebSocketConnection(t *testing.T) {
	httpServer := setup()
	defer httpServer.Close()

  fmt.Println("TestWeb")
	// Create a WebSocket client
	u := "ws" + httpServer.URL[len("http"):]

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(u, nil)
	assert.NoError(t, err, "Dial should not return an error")
  fmt.Println("connection established")
	defer conn.Close()

	// Send an initial message to join a lobby
  initMessage := `{"type": "message", "id": "1","wpm":"1","cursor":"0"}`
	err = conn.WriteMessage(websocket.TextMessage, []byte(initMessage))
	assert.NoError(t, err, "WriteMessage should not return an error")

	// Read the message from the server to verify connection
  fmt.Println("Before conn.readmessage")
	_, message, err := conn.ReadMessage()
  fmt.Println("After conn.ReadMessage")
	assert.NoError(t, err, "ReadMessage should not return an error")
	assert.NotNil(t, message, "Message should not be nil")
}

func TestBroadcastMessage(t *testing.T) {
	httpServer := setup()
	defer httpServer.Close()

	u := "ws" + httpServer.URL[len("http"):]

	dialer := websocket.DefaultDialer

	// Connect the first client
	conn1, _, err := dialer.Dial(u, nil)
	assert.NoError(t, err, "Dial should not return an error")
	defer conn1.Close()
	initMessage1 := `{"id": "1", "lobbyid": "lobby1"}`
	err = conn1.WriteMessage(websocket.TextMessage, []byte(initMessage1))
	assert.NoError(t, err, "WriteMessage should not return an error")

	// Connect the second client
	conn2, _, err := dialer.Dial(u, nil)
	assert.NoError(t, err, "Dial should not return an error")
	defer conn2.Close()
	initMessage2 := `{"id": "2", "lobbyid": "lobby1"}`
	err = conn2.WriteMessage(websocket.TextMessage, []byte(initMessage2))
	assert.NoError(t, err, "WriteMessage should not return an error")

	// Send a message from the first client
	message := `{"type": "message", "id": "1", "wpm": "100", "cursor": "10"}`
	err = conn1.WriteMessage(websocket.TextMessage, []byte(message))
	assert.NoError(t, err, "WriteMessage should not return an error")

	// Both clients should receive the message
	_, received1, err := conn1.ReadMessage()
	assert.NoError(t, err, "ReadMessage should not return an error")
	assert.Equal(t, message, string(received1), "Messages should be equal")

	_, received2, err := conn2.ReadMessage()
	assert.NoError(t, err, "ReadMessage should not return an error")
	assert.Equal(t, message, string(received2), "Messages should be equal")
}

func TestClientRegistration(t *testing.T) {
	httpServer := setup()
	defer httpServer.Close()

	u := "ws" + httpServer.URL[len("http"):]

	dialer := websocket.DefaultDialer

	// Connect a client
	conn, _, err := dialer.Dial(u, nil)
	assert.NoError(t, err, "Dial should not return an error")
	defer conn.Close()
	initMessage := `{"id": "1", "lobbyid": "lobby1"}`
	err = conn.WriteMessage(websocket.TextMessage, []byte(initMessage))
	assert.NoError(t, err, "WriteMessage should not return an error")

	// Wait a moment for the client to register
	time.Sleep(100 * time.Millisecond)

	lobby := server.getLobby("lobby1")
	assert.NotNil(t, lobby, "Lobby should exist")
	assert.Equal(t, 1, len(lobby.clients), "There should be one client in the lobby")

	// Unregister the client
	conn.Close()
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 0, len(lobby.clients), "There should be no clients in the lobby")
}


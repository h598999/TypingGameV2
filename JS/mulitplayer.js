console.log("Hello from multiplayer");


export function setupWebSocket(lobbyId) {
   let ws = new WebSocket("ws://localhost:8080/ws");

    ws.addEventListener('open', function () {
        console.log("Connected to the WebSocket");
        const initialMessage = JSON.stringify({ lobbyid: lobbyId });
        ws.send(initialMessage);
    });

    ws.addEventListener('error', function (event) {
        console.error("WebSocket error observed:", event);
    });

    ws.addEventListener('close', function () {
        console.log("WebSocket is closed now.");
    });

    ws.onmessage = function (event) {
        const message = JSON.parse(event.data);
        console.log(message)
    };
  return ws
}

export function joinLobby() {
    const lobbyId = document.getElementById('join-input').value.trim();
    if (lobbyId) {
        setupWebSocket(lobbyId);
        document.getElementById('create').style.display = 'none';
    } else {
        alert("Please enter a Lobby ID to join.");
    }
}

export function createLobby() {
    const lobbyId = document.getElementById('create-input').value.trim();
    if (lobbyId) {
        setupWebSocket(lobbyId);
        document.getElementById('join').style.display = 'none';
    } else {
        alert("Please enter a Lobby ID to create.");
    }
}

export function sendMessage(ws, pressedchars, pressedspace) {
    const message = {
        wpm: pressedchars, // Replace with dynamic username
        index: pressedspace
    };
    if (ws) {
        ws.send(JSON.stringify(message));
    } else {
        alert("WebSocket connection is not established.");
    }
}

document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search);
  const lobbyId = params.get('lobbyid');
  let charspressed = 0
  let spacepressed = 0

  if (lobbyId) {
    const ws = setupWebSocket(lobbyId);
    const sendmsdbutton = document.getElementById("sendmsgbutton")
    sendmsdbutton.addEventListener('click', () => sendMessage(ws))
    document.addEventListener('keydown', function(event){
      if (event.key != ' '){
        charspressed++
      } else {
        spacepressed++
      }
      sendMessage(ws, charspressed, spacepressed)
    })
  } else {
    console.error("Lobby ID not found in query parameters.");
  }

});




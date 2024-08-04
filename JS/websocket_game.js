export class WebSocketManager {
  constructor(){}
  setupWebSocket(lobbyId) {
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
      console.log(message);
    };

    return ws;
  }

  sendMessage(ws, wpm, cursor){
    const message = {
        wpm: wpm, // Replace with dynamic username
        cursor: cursor 
    };
    if (ws) {
        ws.send(JSON.stringify(message));
    } else {
        alert("WebSocket connection is not established.");
    }
  }
}


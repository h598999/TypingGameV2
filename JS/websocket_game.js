
export class WebSocketBuilder {

  constructor(clientid, lobbyid){
    this.clientid = clientid;
    this.lobbyid = lobbyid;
    this.ws = null
  }

  setupWebSocket() {
    let ws = new WebSocket("ws://localhost:8080/ws");

    ws.addEventListener('open', function () {
      console.log("Connected to the WebSocket");
      const initialMessage = JSON.stringify({ lobbyid: self.lobbyId });
      ws.send(initialMessage);
    });

    ws.addEventListener('error', function (event) {
      console.error("WebSocket error observed:", event);
    });

    ws.addEventListener('close', function () {
      console.log("WebSocket is closed now.");
    });
    this.ws = ws
    return this
}

setOnMessage(f){
  this.ws.onmessage = f
  return this
}

getWS(){
  return this.ws
}
}


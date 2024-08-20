import {WebSocketBuilder} from "./websocket_game.js"

const params = new URLSearchParams(window.location.search);
const lobbyId = params.get('lobbyid');

const clientId = params.get('clientid')


// export function setupWebSocket(lobbyId) {
//    let ws = new WebSocket("ws://localhost:8080/ws");
//
//     ws.addEventListener('open', function () {
//         console.log("Connected to the WebSocket");
//         const initialMessage = JSON.stringify({ lobbyid: lobbyId });
//         ws.send(initialMessage);
//     });
//
//     ws.addEventListener('error', function (event) {
//         console.error("WebSocket error observed:", event);
//     });
//
//     ws.addEventListener('close', function () {
//         console.log("WebSocket is closed now.");
//     });
//
//     ws.onmessage = function (event) {
//         const message = JSON.parse(event.data);
//         if (message.type == "state"){
//           if (message.state == "start"){
//             window.location.href = "/MPgame?lobbyid="+lobbyId;
//           }
//         }
//         console.log(message)
//     };
//   return ws
// }
//
// export function joinLobby() {
//
//     const lobbyId = document.getElementById('join-input').value.trim();
//     if (lobbyId) {
//         setupWebSocket(lobbyId);
//         document.getElementById('create').style.display = 'none';
//     } else {
//         alert("Please enter a Lobby ID to join.");
//     }
// }
//
// export function createLobby() {
//     const lobbyId = document.getElementById('create-input').value.trim();
//     if (lobbyId) {
//         setupWebSocket(lobbyId);
//         document.getElementById('join').style.display = 'none';
//     } else {
//         alert("Please enter a Lobby ID to create.");
//     }
// }

export function sendMessage(ws, state) {
    const message = {
        type: "state",
        clientid: clientId,
        state: state
    };
    if (ws) {
        ws.send(JSON.stringify(message));
    } else {
        alert("WebSocket connection is not established.");
    }
}

function handleMessage(event){
  const message = JSON.parse(event.data);
  if (message.type == "state"){
    if (message.state == "start"){
      window.location.href = "/MPgame?lobbyid="+lobbyId+"&clientid="+clientId;
    }
  }
  console.log(message)
}

document.addEventListener("DOMContentLoaded", () => {
  if (lobbyId) {
    let ws = new WebSocketBuilder(clientId, lobbyId)
    .setupWebSocket()
    .setOnMessage((event) => handleMessage(event))
    .getWS()
    const sendmsdbutton = document.getElementById("sendmsgbutton")
    sendmsdbutton.addEventListener('click', () => Console.log("start"))
    const startgameButton = document.getElementById("StartGame")
    startgameButton.addEventListener('click', () => sendMessage(ws, "start"))
  } else {
    console.error("Lobby ID not found in query parameters.");
  }
})


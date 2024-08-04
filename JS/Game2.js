import { fetchLorem, fetchWords } from "./wordClient.js"
import { WPMCalculator } from "./WPMCalculator.js"
// import  {WebSocketManager} from "./websocket_game.js"
//
const params = new URLSearchParams(window.location.search);
const lobbyId = params.get('lobbyid');

function generateUniqueId() {
  return 'xxxxxx'.replace(/x/g, function() {
    return Math.floor(Math.random() * 16).toString(16);
  }) + Date.now().toString(16);
}

let enemyWPM = 0
const clientid = generateUniqueId()

function setupWebSocket(lobbyId, cursor, wordDiv) {
   let ws = new WebSocket("ws://localhost:8080/ws");

    ws.addEventListener('open', function () {
        console.log("Connected to the WebSocket");
      const initialMessage = JSON.stringify({clientid: clientid, lobbyId: lobbyId});
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
      if ( message.clientid  != clientid){
        console.log(message)
      
      if (message.type == "state"){
        if (message.state == "finish"){
          endGame(0,"Lost")
        }
      }
      if (message.type == "message"){
        enemyWPM = parseInt(message.wpm)
        updateEnemyCursorPosition(cursor, wordDiv, parseInt(message.index))
      }
      }
    };
  return ws
}

function sendMessage(ws, pressedchars, pressedspace) {
    const message = {
        type: "message",
        clientid: clientid,
        wpm: pressedchars, // Replace with dynamic username
        index: pressedspace
    };
    if (ws) {
        ws.send(JSON.stringify(message));
    } else {
        alert("WebSocket connection is not established.");
    }
}

function sendStateMessage(ws, state) {
    const message = {
        type: "state",
        clientid: clientid,
        state: state
    };
    if (ws) {
        ws.send(JSON.stringify(message));
    } else {
        alert("WebSocket connection is not established.");
    }
}


function wrapWords(words){
  var wordsStr = ""
  for (let i = 0; i<words.length; i++){
    wordsStr += words[i] + " "
  }
  return wordsStr.trimEnd()
}

function wrapWord(word) {
  const wordDiv = document.createElement('div');
  wordDiv.classList.add("wordsdiv");
  for (let i = 0; i < word.length; i++) {
    const charSpan = document.createElement('span');
    charSpan.classList.add('char');
    charSpan.innerHTML = word[i];
    wordDiv.appendChild(charSpan);
  }
  return wordDiv;
}

function newWord(gameContainer, wordCount, words) {
  if (wordCount < words.length) {
    gameContainer.innerHTML = '';
    const wordDiv = wrapWord(words[wordCount]);
    gameContainer.appendChild(wordDiv);
    return wordDiv;
  } else {
    console.log("No more words available!");
    return null;
  }
}

function endGame(wpm, result){
  document.body.innerHTML = ''
  // Create new elements and add new functionality
  const resultDiv = document.createElement('div');
  resultDiv.id = 'result';

  const wpmDisplay = document.createElement('h2');
  wpmDisplay.id = 'wpm-display';

  const enemydisplay = document.createElement('h2');
  enemydisplay.id = 'enemy-display';

  const resultdisplay = document.createElement('h2');
  resultdisplay.id = 'result-display';

  const restartButton = document.createElement('button');
  restartButton.textContent = 'Restart';
  restartButton.id = 'restart-button';

  // Add new elements to the body
  document.body.appendChild(resultDiv);
  document.body.appendChild(wpmDisplay);
  document.body.appendChild(enemydisplay);
  document.body.appendChild(resultdisplay);
  document.body.appendChild(restartButton);

  // Display results
  wpmDisplay.textContent = `WPM: ${wpm}`;
  enemydisplay.textContent = `EnemyWPM: ${enemyWPM}`

  if (result == "won"){
    resultdisplay.textContent = "You won!"
  } else {
    resultdisplay.textContent = "You Lost!"
  }

  // Add new functionality
  restartButton.addEventListener('click', () => {
    location.reload(); // This will refresh the page, restarting the game
  });
}
function updateEnemyCursorPosition(cursor,wordDiv, index) {
  const prevChar = wordDiv.children[index-1]
  if (!prevChar.classList.contains("green")){
    prevChar.classList.add("grey")
  }
  const currentChar = wordDiv.children[index];
  const rect = currentChar.getBoundingClientRect();
  const parentRect = wordDiv.getBoundingClientRect();
  cursor.style.left = `${rect.left - parentRect.left}px`;
  cursor.style.top = `${rect.top - parentRect.top}px`;
}

document.addEventListener('DOMContentLoaded', async () => {
  // const Lorem = await fetchLorem();
  const calc = new WPMCalculator();
  const Words = await fetchWords();
  // console.log(Lorem)
  // console.log(Words)
  const dbwords = wrapWords(Words)
  let index = 0;
  let wordCount = 0;
  let prevIndex = 0;
  const wpmcontainer = document.getElementById('WPM');
  wpmcontainer.classList.add('container');
  wpmcontainer.innerHTML = "WPM: " + 0
  const gameContainer = document.getElementById('game-container');
  let wordDiv = newWord(gameContainer, wordCount, [dbwords]);

  // Create the cursor element
  const cursor = document.createElement('span');
  cursor.classList.add('cursor');
  wordDiv.appendChild(cursor);

  const enemycursor = document.createElement('span');
  enemycursor.classList.add('enemyCursor')
  wordDiv.appendChild(enemycursor);

  const ws = setupWebSocket(lobbyId, enemycursor, wordDiv)


  function updateCursorPosition() {
    const currentChar = wordDiv.children[index];
    const rect = currentChar.getBoundingClientRect();
    const parentRect = wordDiv.getBoundingClientRect();
    cursor.style.left = `${rect.left - parentRect.left}px`;
    cursor.style.top = `${rect.top - parentRect.top}px`;
  }

  updateCursorPosition();

  function handleKeyDown(event){
    if (event.key === dbwords[index]) {
      if (index == 0){
        calc.start()
      }
      wordDiv.children[index].classList.remove("red");
      wordDiv.children[index].classList.add("green");
      if (event.key === ' ') {
        wordCount++;
        prevIndex = calc.wordCompleted(index, prevIndex)
        wpmcontainer.innerHTML = "WPM: " + calc.calculateWPM().wpm
      }
      index++;
      updateCursorPosition();
      sendMessage(ws, String(calc.calculateWPM().wpm), String(index))
    } else {
      wordDiv.children[index].classList.add("red");
    }

    if (index === dbwords.length) {
      console.log("You finished with a WPM of: " + calc.calculateWPM().wpm)
      sendStateMessage(ws, "finish")
      endGame(calc.calculateWPM().wpm, "won")
      document.removeEventListener('keydown', handleKeyDown)
    }
  }
  document.addEventListener('keydown', handleKeyDown);
});

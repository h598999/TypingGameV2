import { fetchLorem } from "./wordClient.js"
import { WPMCalculator } from "./WPMCalculator.js"

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


function endGame(wpm){
  document.body.innerHTML = ''
    // Create new elements and add new functionality
  const resultDiv = document.createElement('div');
  resultDiv.id = 'result';
  
  const wpmDisplay = document.createElement('h2');
  wpmDisplay.id = 'wpm-display';
  
  const accuracyDisplay = document.createElement('h2');
  accuracyDisplay.id = 'accuracy-display';
  
  const restartButton = document.createElement('button');
  restartButton.textContent = 'Restart';
  restartButton.id = 'restart-button';

  // Add new elements to the body
  document.body.appendChild(resultDiv);
  document.body.appendChild(wpmDisplay);
  document.body.appendChild(accuracyDisplay);
  document.body.appendChild(restartButton);

  // Display results
  wpmDisplay.textContent = `WPM: ${wpm}`;

  // Add new functionality
  restartButton.addEventListener('click', () => {
    location.reload(); // This will refresh the page, restarting the game
  });
}

document.addEventListener('DOMContentLoaded', async () => {
    const Lorem = await fetchLorem();
    const calc = new WPMCalculator()
    console.log(Lorem)
    let index = 0;
    let wordCount = 0;
    let prevIndex = 0;
    const wpmcontainer = document.getElementById('WPM');
    wpmcontainer.classList.add('container');
    wpmcontainer.innerHTML = "WPM: " + 0
    const gameContainer = document.getElementById('game-container');
    let wordDiv = newWord(gameContainer, wordCount, [Lorem]);
    

    // Create the cursor element
    const cursor = document.createElement('span');
    cursor.classList.add('cursor');
    wordDiv.appendChild(cursor);

    function updateCursorPosition() {
        const currentChar = wordDiv.children[index];
        const rect = currentChar.getBoundingClientRect();
        const parentRect = wordDiv.getBoundingClientRect();
        cursor.style.left = `${rect.left - parentRect.left}px`;
        cursor.style.top = `${rect.top - parentRect.top}px`;
    }

    updateCursorPosition();

  function handleKeyDown(event){
    if (event.key === Lorem[index]) {
      if (index == 0){
        calc.start()
      }
      wordDiv.children[index].classList.add("green");
      if (event.key === ' ') {
        wordCount++;
        // console.log("Start: " + calc.startTime + "End: " + new Date())
        // console.log("Index" + index, "Previndex" + prevIndex)
        prevIndex = calc.wordCompleted(index, prevIndex)
        // console.log("WPM: " + calc.calculateWPM().wpm)
        wpmcontainer.innerHTML = "WPM: " + calc.calculateWPM().wpm
      }
      index++;
      updateCursorPosition();
    }

    if (index === Lorem.length) {
      console.log("You finished with a WPM of: " + calc.calculateWPM().wpm)
      endGame(calc.calculateWPM().wpm)
      document.removeEventListener('keydown', handleKeyDown)
    }
  }

    document.addEventListener('keydown', handleKeyDown);
});

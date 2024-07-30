import { fetchWords } from "./wordClient.js"

function wrapWord(word){
  const wordDiv = document.createElement('div');
  wordDiv.id = "wordDiv";
  wordDiv.classList.add("wordsdiv");
  for (let i = 0; i<word.length; i++){
      const correctspan = document.createElement('span');
      correctspan.innerHTML = word[i];
      wordDiv.appendChild(correctspan);
  }
  const body = document.querySelector("body");
  body.appendChild(wordDiv);
  return wordDiv;
}

function newWord(wordDiv, wordCount, words) {
  if (wordCount < words.length) {
    const body = document.querySelector("body");
    body.removeChild(wordDiv);
    return wrapWord(words[wordCount]);
  } else {
    console.log("No more words available!");
    return null; // or handle the end of words list scenario
  }
}

document.addEventListener('DOMContentLoaded', async() =>{
  const words = await fetchWords();
  let index = 0;
  let wordCount = 0;
  let wordDiv = wrapWord(words[wordCount]);
  

  document.addEventListener('keydown', function(event){
    if (event.key === words[wordCount][index]){
      // wordDiv.classList.add("my-text");
      wordDiv.children[index].classList.add("green");
      // console.log("Correct");
      index = index + 1;
    }
if (index === words[wordCount].length){
      // wordDiv.classList.remove("my-text");
      wordCount = wordCount + 1;
      index = 0;
      // console.log(wordCount, index);
      wordDiv = newWord(wordDiv, wordCount, words);
    }
  });
});


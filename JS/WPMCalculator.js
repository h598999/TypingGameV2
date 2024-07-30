export class WPMCalculator {
  constructor() {
    this.startTime = null;
    this.wordCount = 0;
    this.charCount = 0;
    this.errors = 0;
  }

  start() {
    this.startTime = new Date();
  }

  wordCompleted(index, previndex) {
    if (this.startTime === null) {
      throw new Error("Timer not started");
    }
    this.wordCount++;
    this.charCount += index - previndex
    return index
  }

  calculateWPM() {
    if (this.startTime === null) {
      throw new Error("Timer not started");
    }
    const currentTime = new Date();
    const elapsedMinutes = (currentTime - this.startTime) / 60000; // Convert ms to minutes
    const grossWPM = this.charCount / 5 / elapsedMinutes; // Assume average word length of 5 characters
    const netWPM = grossWPM - (this.errors / elapsedMinutes);
    
    return {
      wpm: Math.round(netWPM),
      // accuracy: Math.round((1 - this.errors / this.charCount) * 100)
    };
  }
}

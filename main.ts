/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program uses a word search in a given sentence.
 */

console.log("Please enter a sentence?");
const sentence: string = prompt("Enter a sentence: ") ?? "";

console.log("Please enter a word to search for in your sentence?");
const word: string = prompt("Enter a word: ") ?? "";

if (sentence.includes(word)) {
  console.log("Hooray, the word " + word +
    " was found in the sentence: " + sentence);
} else {
  console.log("Sorry, the word " + word +
    " was not found in the sentence: " + sentence);
}
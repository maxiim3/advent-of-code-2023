"use strict";
// Path: day-1/index.ts
Object.defineProperty(exports, "__esModule", { value: true });
var fs = require("fs");
function readInputFile() {
    var PATH = "./day-1/input.txt";
    var file = fs.readFileSync(PATH, { encoding: 'utf-8' });
    // console.log(sample_file)
    return file;
}
function splitLinesFromFileInput(fileInput) {
    var lines = fileInput.split("\n");
    // console.log(lines)
    return lines;
}
function splitCharsInLine(line) {
    var listOfChars = line.split("\n");
    // console.log(listOfChars)
    return listOfChars;
}
function extractNumbersFromChars(chars) {
    var numbers = chars[0].split("")
        .filter(function (value) { return Number.parseInt(value); })
        .map(function (value) { return Number.parseInt(value); });
    return numbers;
}
function extractFirstAndLastNumber(listOfNumbers) {
    var first = listOfNumbers.at(0) || 0;
    var last = listOfNumbers.at(-1) || 0;
    return [first, last];
}
function mergeNumbers(twoDigitsArray) {
    var first = twoDigitsArray[0], last = twoDigitsArray[1];
    var stringFromDigits = "".concat(first).concat(last);
    var parsedString = Number.parseInt(stringFromDigits);
    // console.log(twoDigitsArray, parsedString);
    return parsedString;
}
function calculateTotal(listOfValues) {
    var result = listOfValues.reduce(function (accumulator, current) { return accumulator + current; }, 0);
    // console.log(result)
    return result;
}
// Had to add all matching values (including numbers)
var matchedNumbers = {
    "1": 1,
    "one": 1,
    "2": 2,
    "two": 2,
    "3": 3,
    "three": 3,
    "4": 4,
    "four": 4,
    "5": 5,
    "five": 5,
    "6": 6,
    "six": 6,
    "7": 7,
    "seven": 7,
    "8": 8,
    "eight": 8,
    "9": 9,
    "nine": 9,
};
var spelled = Object.keys(matchedNumbers);
var regEx = RegExp(spelled.join("|"), "ig");
function matchAndReplace(string) {
    // @ts-ignore
    var result = string.replaceAll(regEx, function (matchedString) { return matchedNumbers[matchedString]; });
    console.log(result, string);
    return result;
}
function pipePartOne(lines) {
    var numbers = lines.map(function (line) {
        var chars = splitCharsInLine(line);
        var nums = extractNumbersFromChars(chars);
        var digits = extractFirstAndLastNumber(nums);
        var merged = mergeNumbers(digits);
        return merged;
    });
    return calculateTotal(numbers);
}
function pipePartTwo(lines) {
    var numbers = lines.map(function (line) {
        console.log(line);
        var nums = matchAllNumbers(line);
        console.log(nums);
        var digits = extractFirstAndLastNumber(nums);
        console.log(digits);
        var merged = mergeNumbers(digits);
        console.log(merged);
        return merged;
    });
    return calculateTotal(numbers);
}
/**
 * This function searches through 'str' to find all occurrences of numbers based on a regex pattern.
 * When a match is found, 'regEx.lastIndex' is automatically updated to the position after the last character of the match.
 * However, we need to manually reset 'regEx.lastIndex' to 'match.index + 1' to ensure the regex starts searching from the next character after the current match.
 * This prevents skipping potential matches in subsequent iterations.
 */
function matchAllNumbers(str) {
    var matches = [];
    var match;
    while ((match = regEx.exec(str)) !== null) {
        // @ts-ignore
        matches.push(matchedNumbers[match[0].toLowerCase()]);
        // console.log(match[0], "last index regex", regEx.lastIndex)
        // console.log("match index", match.index)
        regEx.lastIndex = match.index + 1;
    }
    return matches;
}
function main() {
    var file = readInputFile();
    var list_of_lines = splitLinesFromFileInput(file);
    var partOneResult = pipePartOne(list_of_lines);
    var partTwoResult = pipePartTwo(list_of_lines);
    console.log(partOneResult, partTwoResult);
}
main();
console.log(regEx);

// Path: day-1/index.ts

import * as fs from "fs";


function readInputFile() {
    const PATH = "./day-1/input.txt";
    const file = fs.readFileSync(PATH, {encoding: 'utf-8'})

    // console.log(sample_file)
    return file;
}

function splitLinesFromFileInput(fileInput: string): string[] {
    const lines = fileInput.split("\n");

    // console.log(lines)
    return lines;
}

function splitCharsInLine(line: string): string[] {
    const listOfChars = line.split("\n");

    // console.log(listOfChars)
    return listOfChars
}

function extractNumbersFromChars(chars: string[]): number[] {
    const numbers = chars[0].split("")
        .filter(value => Number.parseInt(value))
        .map(value => Number.parseInt(value));

    return numbers
}

function extractFirstAndLastNumber(listOfNumbers: number[]): [number, number] {
    const first = listOfNumbers.at(0) || 0
    const last = listOfNumbers.at(-1) || 0

    return [first, last]
}

function mergeNumbers(twoDigitsArray: [number, number]): number {
    const [first, last] = twoDigitsArray;
    const stringFromDigits = `${first}${last}`
    const parsedString = Number.parseInt(stringFromDigits);

    // console.log(twoDigitsArray, parsedString);
    return parsedString
}

function calculateTotal(listOfValues: number[]): number {
    const result = listOfValues.reduce((accumulator, current) => accumulator + current, 0)

    // console.log(result)
    return result;
}

// Had to add all matching values (including numbers)
const matchedNumbers = {
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
}

const spelled = Object.keys(matchedNumbers);
const regEx = RegExp(spelled.join("|"), "ig")

function matchAndReplace(string: string): string {
    // @ts-ignore
    let result = string.replaceAll(regEx, (matchedString) => matchedNumbers[matchedString]!)
    console.log(result, string)

    return result
}

function pipePartOne(lines: string[]): number {

    const numbers = lines.map(line => {
        const chars = splitCharsInLine(line);
        const nums = extractNumbersFromChars(chars)
        const digits = extractFirstAndLastNumber(nums)
        const merged = mergeNumbers(digits);

        return merged
    })
    return calculateTotal(numbers)

}

function pipePartTwo(lines: string[]): number {

    const numbers = lines.map(line => {
        console.log(line)
        const nums = matchAllNumbers(line)
        console.log(nums)
        const digits = extractFirstAndLastNumber(nums)
        console.log(digits)
        const merged = mergeNumbers(digits);
        console.log(merged)

        return merged
    })
    return calculateTotal(numbers)

}

/**
 * This function searches through 'str' to find all occurrences of numbers based on a regex pattern.
 * When a match is found, 'regEx.lastIndex' is automatically updated to the position after the last character of the match.
 * However, we need to manually reset 'regEx.lastIndex' to 'match.index + 1' to ensure the regex starts searching from the next character after the current match.
 * This prevents skipping potential matches in subsequent iterations.
 */
function matchAllNumbers(str: string): number[] {

    const matches: number[] = [];

    let match;

    while ((match = regEx.exec(str)) !== null) {

        // @ts-ignore
        matches.push(matchedNumbers[match[0].toLowerCase()])
        // console.log(match[0], "last index regex", regEx.lastIndex)
        // console.log("match index", match.index)
        regEx.lastIndex = match.index + 1
    }
    return matches
}

function main() {

    const file = readInputFile();
    const list_of_lines = splitLinesFromFileInput(file);
    const partOneResult = pipePartOne(list_of_lines)

    const partTwoResult = pipePartTwo(list_of_lines)
    console.log(partOneResult, partTwoResult)

}

main()


console.log(regEx)

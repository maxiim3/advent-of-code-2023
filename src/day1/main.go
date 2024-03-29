package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Day 1!")

	file := readFile("./src/day1/input.txt")

	lines := readLines(file)

	// Accumulatored result
	var acc int

	for _, currentLine := range lines {

		firstItem, lastItem := string(currentLine[0]), string(currentLine[len(currentLine)-1])
		_, errFirst := strconv.Atoi(firstItem)
		_, errLast := strconv.Atoi(lastItem)

		if errFirst == nil && errLast == nil {
			digits := mergeDigits(Digits{
				values: [2]string{firstItem, lastItem}})
			acc += digits
		} else {
			currentLine = replaceSpelledByNumbers(currentLine)
			nums := extractNumbersFromLine(currentLine)
			digits := mergeDigits(nums)
			acc += digits
		}

	}

	fmt.Println(acc)

}

func readFile(path string) []byte {
	file, err := os.ReadFile(path)

	if err != nil {
		panic("Cannot read the file ")
	}

	return file
}

func readLines(file []byte) []string {
	lines := strings.Split(string(file), "\n")

	return lines
}

// Replace the spelled numbers from the line into numbers
func replaceSpelledByNumbers(line string) string {
	dictionary := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	regEx := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")

	matchedItems := regEx.FindAllString(line, 10)

	if matchedItems == nil {
		return line
	}

	for _, match := range matchedItems {
		line = strings.ReplaceAll(line, match, dictionary[match])
	}

	return line
}

type Digits struct {
	values [2]string
}

// Filter out the numbers from the line
//
// Returns a slice of numbers as string type
func extractNumbersFromLine(line string) Digits {
	allNums := []string{}

	for _, chars := range line {
		chars := string(chars)
		if _, err := strconv.Atoi(chars); err == nil {
			allNums = append(allNums, chars)
		}
	}

	switch len(allNums) {
	case 0:
		return Digits{values: [2]string{"0", "0"}}
	case 1:
		return Digits{values: [2]string{allNums[0], allNums[0]}}
	default:
		return Digits{values: [2]string{allNums[0], allNums[len(allNums)-1]}}
	}
}

// Create a digit from the first and last element of the slice : { X Y W Z } -> XZ
//
// # If the slice contains no item, returns 0
//
// Ifgf there is only one element in the slice, return a digit from the element : { X } -> XX
func mergeDigits(digit Digits) int {

	first, last := digit.values[0], digit.values[1]
	digits, err := strconv.Atoi(fmt.Sprint(first, last))

	if err != nil {
		panic(err)
	}

	return digits
}

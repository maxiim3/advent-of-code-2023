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

	file := readFile("./src/day1/test-2.txt")

	lines := readLines(file)

	// Accumulatored result
	acc := 0

	for _, currentLine := range lines {

		firstItem, lastItem := string(currentLine[0]), string(currentLine[len(currentLine)-1])
		_, errFirst := strconv.Atoi(firstItem)
		_, errLast := strconv.Atoi(lastItem)

		if errFirst == nil && errLast == nil {
			digits := mergeDigits([]string{firstItem, lastItem})
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

// Filter out the numbers from the line
//
// Returns a slice of numbers as string type
func extractNumbersFromLine(line string) []string {
	nums := []string{}

	for j := 0; j < len(line); j++ {
		// Cast element byte -> string
		item := string(line[j])

		// Cast to integer
		if _, err := strconv.Atoi(item); err == nil {
			nums = append(nums, item)
		}
	}
	return nums
}

// Create a digit from the first and last element of the slice : { X Y W Z } -> XZ
//
// # If the slice contains no item, returns 0
//
// Ifgf there is only one element in the slice, return a digit from the element : { X } -> XX
func mergeDigits(nums []string) int {
	fmt.Println(nums)

	if len(nums) == 0 {
		return 0
	}

	if len(nums) > 1 {
		first := nums[0]
		last := nums[len(nums)-1]

		digits, err := strconv.Atoi(fmt.Sprint(first, last))

		if err != nil {
			panic(err)
		}

		return digits
	}
}

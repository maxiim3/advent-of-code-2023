package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func main() {
	fmt.Println("Hello, Day 1!")
	file, err := os.ReadFile("./src/day1/test-2.txt")

	if err != nil {
		panic("Cannot read the file ")
	}

	lines := strings.Split(string(file), "\n")

	acc := 0
	for i := 0; i < len(lines); i++ {
		lines[i] = transformSpelledNumbers(lines[i])
		num := filterNumbers(lines[i])
		digits := extractFirstAndLast(num)
		acc += digits
	}

	fmt.Println(acc)

}

func transformSpelledNumbers(line string) string {
	hash := map[string]string{
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

func filterNumbers(slice string) []string {
	nums := []string{}
	for j := 0; j < len(slice); j++ {
		item := string(slice[j])

		if _, err := strconv.Atoi(item); err == nil {
			nums = append(nums, item)
		}
	}
	return nums
}

func extractFirstAndLast(nums []string) int {
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

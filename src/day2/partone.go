package main

import (
	"fmt"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func partone() {
	fmt.Println("Part one\n")
	testFile := fileutils.ParseFile("./src/day2/input.txt")

	var lines []Line

	for _, line := range testFile.Lines {
		l := CreateLine(line)
		if l.Bluecube.Max <= 14 && l.Redcube.Max <= 12 && l.Greencube.Max <= 13 {
			lines = append(lines, l)
		}
	}

	total := 0
	for _, l := range lines {
		total += l.Index
	}

	fmt.Println("The result is", total)
}

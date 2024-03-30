package main

import (
	"fmt"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func parttwo() {
	fmt.Println("Part two \n")
	testFile := fileutils.ParseFile("./src/day2/input.txt")

	total := 0
	for _, line := range testFile.Lines {
		l := CreateLine(line)
		pow := l.Bluecube.Max * l.Redcube.Max * l.Greencube.Max
		total += pow
	}

	fmt.Println("The result is", total)
}

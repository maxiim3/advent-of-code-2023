package main

import (
	"regexp"
	"strconv"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

type Row = [140]Item
type Matrix = [140]Row

func createMatrix(f fileutils.File) [140][140]Item {
	var matrix [140][140]Item
	for abs, el := range f.Lines {
		runes := []rune(el)

		var row [140]Item

		for ord, r := range runes {
			isNumber := runeIsNumber(r)
			newItem := Item{
				x: abs,
				y: ord,
				value: func() any {
					if isNumber {
						number, _ := strconv.ParseUint(string(r), 10, 64)
						return uint(number)
					}
					return r
				}(),
				isInt:        isNumber,
				isSymbol:     !isNumber && r == '*',
				isPartNumber: false,
			}

			row[ord] = newItem
		}
		matrix[abs] = row
	}

	return matrix

}

func runeIsNumber(r rune) bool {
	digitRegex := regexp.MustCompile(`[0-9]`)
	s := string(r)

	return digitRegex.MatchString(s)
}

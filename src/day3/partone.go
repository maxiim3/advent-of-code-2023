package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func partone() {
	fmt.Println("")
	fmt.Println("\tPart One is on...")

	f := fileutils.ParseFile("./src/day3/input-sample.txt")

	var matrix [][]Item

	for x, el := range f.Lines {

		items := strings.Split(el, "")

		var row []Item

		for y, item := range items {
			value, isSymbol := strconv.Atoi(item)

			if isSymbol != nil {
				if item == "." {
					newItem := Item{
						x:            x,
						y:            y,
						value:        item,
						isInt:        false,
						isSymbol:     false,
						isPartNumber: false,
					}
					row = append(row, newItem)
				} else {
					newItem := Item{
						x:            x,
						y:            y,
						value:        item,
						isInt:        false,
						isSymbol:     true,
						isPartNumber: false,
					}
					row = append(row, newItem)
				}
			} else {
				newItem := Item{
					x:            x,
					y:            y,
					value:        value,
					isInt:        true,
					isSymbol:     false,
					isPartNumber: false,
				}
				row = append(row, newItem)
			}
		}
		matrix = append(matrix, row)
	}

	// fmt.Println(matrix)

	colsize, rowsize := len(matrix), len(matrix[0])
	fmt.Println(colsize, rowsize)

	result := 0
	for k := 0; k < colsize; k++ {
		nums := [131]LineNumbers{}
		numIndex := 0
		for v := 0; v < rowsize; v++ {
			item := matrix[k][v]
			if item.isInt {
				var s Surrounding

				if k == 0 && v == 0 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         false,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   false,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  matrix[k+1][v+1].isSymbol,
					}

				} else if k == 0 && v == rowsize-1 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         matrix[k][v-1].isSymbol,
						right:        false,
						bottomleft:   matrix[k+1][v-1].isSymbol,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  false,
					}
				} else if k == 0 && v < rowsize-1 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         matrix[k][v-1].isSymbol,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   matrix[k+1][v-1].isSymbol,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  matrix[k+1][v+1].isSymbol,
					}
				} else if k == colsize-1 && v == 0 {
					s = Surrounding{
						topleft:      false,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     matrix[k-1][v+1].isSymbol,
						left:         false,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   false,
						bottomcenter: false,
						bottomright:  false,
					}
				} else if k == colsize-1 && v < rowsize-1 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isSymbol,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     matrix[k-1][v+1].isSymbol,
						left:         matrix[k][v-1].isSymbol,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   false,
						bottomcenter: false,
						bottomright:  false,
					}
				} else if k == colsize-1 && v == rowsize-1 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isSymbol,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     false,
						left:         matrix[k][v-1].isSymbol,
						right:        false,
						bottomleft:   false,
						bottomcenter: false,
						bottomright:  false,
					}
				} else if v == 0 {
					s = Surrounding{
						topleft:      false,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     matrix[k-1][v+1].isSymbol,
						left:         false,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   false,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  matrix[k+1][v+1].isSymbol,
					}
				} else if v == colsize-1 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isSymbol,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     false,
						left:         matrix[k][v-1].isSymbol,
						right:        false,
						bottomleft:   matrix[k+1][v-1].isSymbol,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  false,
					}
				} else {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isSymbol,
						topcenter:    matrix[k-1][v].isSymbol,
						topright:     matrix[k-1][v+1].isSymbol,
						left:         matrix[k][v-1].isSymbol,
						right:        matrix[k][v+1].isSymbol,
						bottomleft:   matrix[k+1][v-1].isSymbol,
						bottomcenter: matrix[k+1][v].isSymbol,
						bottomright:  matrix[k+1][v+1].isSymbol,
					}
				}

				nums[numIndex].i = fmt.Sprint(nums[numIndex].i, item.value)

				if s.hasAdjascentSymbol() {
					nums[numIndex].hasPartNum = true
					item.isPartNumber = true
				}

			} else {
				numIndex++
			}
		}

		for _, v := range nums {
			if v.hasPartNum {
				result += v.convertToNumber()
			}
		}
	}

	fmt.Println(result)
}

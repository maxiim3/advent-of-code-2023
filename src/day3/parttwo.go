package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func parttwo() {
	debug.FreeOSMemory()
	debug.SetMemoryLimit(4096 * 1 << 20)
	fmt.Println("")
	fmt.Println("\tPart TWO Let's go!")
	fmt.Println("\t==================")
	fmt.Println("")

	f := fileutils.ParseFile("./src/day3/test.txt")

	matrix := createMatrix(f)

	colsize, rowsize := len(matrix), len(matrix[0])
	fmt.Println("\tSIZE", colsize, rowsize)

	gears := []Gear{}
	for k := 0; k < colsize; k++ {
		fmt.Println("new row", k, "of", len(matrix[k]))
		for v := 0; v < rowsize; v++ {
			item := matrix[k][v]
			fmt.Println("parsing: ", item.value, "{", k, ";", v, "}")
			if item.isSymbol {
				gear := Gear{
					partNums: []uint64{},
					x:        item.x,
					y:        item.y,
				}

				var s Surrounding

				if k == 0 && v == 0 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         false,
						right:        matrix[k][v+1].isInt,
						bottomleft:   false,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  matrix[k+1][v+1].isInt,
					}

				} else if k == 0 && v == rowsize-1 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         matrix[k][v-1].isInt,
						right:        false,
						bottomleft:   matrix[k+1][v-1].isInt,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  false,
					}
				} else if k == 0 && v < rowsize-1 {
					s = Surrounding{
						topleft:      false,
						topcenter:    false,
						topright:     false,
						left:         matrix[k][v-1].isInt,
						right:        matrix[k][v+1].isInt,
						bottomleft:   matrix[k+1][v-1].isInt,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  matrix[k+1][v+1].isInt,
					}
				} else if k == colsize-1 && v == rowsize-1 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isInt,
						topcenter:    matrix[k-1][v].isInt,
						topright:     false,
						left:         matrix[k][v-1].isInt,
						right:        false,
						bottomleft:   false,
						bottomcenter: false,
						bottomright:  false,
					}
				} else if k == colsize-1 && v < rowsize-1 && v != 0 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isInt,
						topcenter:    matrix[k-1][v].isInt,
						topright:     matrix[k-1][v+1].isInt,
						left:         matrix[k][v-1].isInt,
						right:        matrix[k][v+1].isInt,
						bottomleft:   false,
						bottomcenter: false,
						bottomright:  false,
					}
				} else if v == 0 {
					s = Surrounding{
						topleft:      false,
						topcenter:    matrix[k-1][v].isInt,
						topright:     matrix[k-1][v+1].isInt,
						left:         false,
						right:        matrix[k][v+1].isInt,
						bottomleft:   false,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  matrix[k+1][v+1].isInt,
					}
				} else if v == colsize-1 {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isInt,
						topcenter:    matrix[k-1][v].isInt,
						topright:     false,
						left:         matrix[k][v-1].isInt,
						right:        false,
						bottomleft:   matrix[k+1][v-1].isInt,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  false,
					}
				} else {
					s = Surrounding{
						topleft:      matrix[k-1][v-1].isInt,
						topcenter:    matrix[k-1][v].isInt,
						topright:     matrix[k-1][v+1].isInt,
						left:         matrix[k][v-1].isInt,
						right:        matrix[k][v+1].isInt,
						bottomleft:   matrix[k+1][v-1].isInt,
						bottomcenter: matrix[k+1][v].isInt,
						bottomright:  matrix[k+1][v+1].isInt,
					}
				}
				if s.left {
					n := parseBackward(matrix[k], v-1)
					gear.partNums = append(gear.partNums, uint64(n))
				}
				if s.right {
					n := parseForward(matrix[k], v+1, rowsize-1)
					gear.partNums = append(gear.partNums, uint64(n))
				}

				if s.topcenter || s.topleft || s.topright {
					if s.topleft && s.topcenter && !s.topright {
						n := parseBackward(matrix[k-1], v)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.topcenter && s.topright && !s.topleft {
						n := parseForward(matrix[k-1], v, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.topcenter && s.topright && s.topleft {
						n := parseBothWays(matrix[k-1], v, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.topcenter && !s.topleft && !s.topright {
						numstr := fmt.Sprint(matrix[k-1][v].value)
						num, err := strconv.Atoi(numstr)
						if err != nil {
							errmsg := fmt.Sprint("Something went wrong with line above", k, v)
							panic(errmsg)
						}
						gear.partNums = append(gear.partNums, uint64(num))
					} else if !s.topcenter && s.topleft && !s.topright {
						n := parseBackward(matrix[k-1], v-1)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if !s.topcenter && s.topright && !s.topleft {
						n := parseForward(matrix[k-1], v+1, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					}
				}
				if s.bottomcenter || s.bottomleft || s.bottomright {
					if s.bottomleft && s.bottomcenter && !s.bottomright {
						n := parseBackward(matrix[k+1], v)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.bottomcenter && s.bottomright && !s.bottomleft {
						n := parseForward(matrix[k+1], v, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.bottomcenter && s.bottomright && s.bottomleft {
						n := parseBothWays(matrix[k+1], v, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if s.bottomcenter && !s.bottomleft && !s.bottomright {
						numstr := fmt.Sprint(matrix[k+1][v].value)
						num, err := strconv.Atoi(numstr)
						if err != nil {
							errmsg := fmt.Sprint("Something went wrong with  line bellow", k, v)
							panic(errmsg)
						}
						gear.partNums = append(gear.partNums, uint64(num))
					} else if !s.bottomcenter && s.bottomleft && !s.bottomright {
						n := parseBackward(matrix[k+1], v-1)
						gear.partNums = append(gear.partNums, uint64(n))
					} else if !s.bottomcenter && s.bottomright && !s.bottomleft {
						n := parseForward(matrix[k+1], v+1, rowsize)
						gear.partNums = append(gear.partNums, uint64(n))
					}
				}

				if gear.isGear() {
					fmt.Println(gear.partNums)
					if len(gears) == 0 {
						gears = append(gears, gear)
					} else {
						for _, g := range gears {
							if g.x != gear.x && g.y != gear.y {
								gears = append(gears, gear)
							}
						}
					}
				}
			}
		}
	}

	total := uint64(0)
	for _, g := range gears {
		total += g.calculateGearRatio()
	}

	fmt.Println(total)
}

func parseBothWays(row []Item, idx int, max int) int {
	forwardNum := parseForward(row, idx+1, max)
	middleNum := row[idx].value
	backwardNum := parseBackward(row, idx-1)

	b := fmt.Sprint(backwardNum)
	m := fmt.Sprint(middleNum)
	f := fmt.Sprint(forwardNum)

	items := []string{b, m, f}
	concat := strings.Join(items, "")
	val, err := strconv.Atoi(concat)

	if err != nil {
		panic("Error while parsing the element in both ways")
	}

	return val
}

func parseForward(row []Item, idx int, max int) int {
	i := idx

	var items []string
parser:
	for {
		if !row[i].isInt {
			break parser
		}

		item := fmt.Sprint(row[i].value)
		items = append(items, item)

		if i == max {
			break parser
		}

		i++
	}

	concat := strings.Join(items, "")

	val, err := strconv.Atoi(concat)

	if err != nil {
		panic("Error while parsing forward the element")
	}

	return val
}

func parseBackward(row []Item, idx int) int {
	i := idx
	var items []string
parser:
	for {
		if !row[i].isInt {
			break parser
		}

		item := fmt.Sprint(row[i].value)
		items = append(items, item)

		if i == 0 {
			break parser
		}

		i--
	}

	reversed := reverseItems(items)
	concat := strings.Join(reversed, "")

	val, err := strconv.Atoi(concat)

	if err != nil {
		panic("Error while parsing backward the element")
	}

	return val
}

func reverseItems(input []string) []string {
	if len(input) == 0 {
		return input
	}

	return append(reverseItems(input[1:]), input[0])
}

type Gear struct {
	partNums []uint64
	x        int
	y        int
}

func (g *Gear) isGear() bool {
	return len(g.partNums) == 2
}

func (g *Gear) calculateGearRatio() uint64 {
	return g.partNums[0] * g.partNums[1]
}

func createMatrix(f fileutils.File) [][]Item {
	var matrix [][]Item
	for x, el := range f.Lines {

		items := strings.Split(el, "")

		var row []Item

		for y, item := range items {
			value, err := strconv.Atoi(item)

			if err != nil {
				if item == "*" {
					newItem := Item{
						x:            x,
						y:            y,
						value:        item,
						isInt:        false,
						isSymbol:     true,
						isPartNumber: false,
					}
					row = append(row, newItem)
				} else {
					newItem := Item{
						x:            x,
						y:            y,
						value:        item,
						isInt:        false,
						isSymbol:     false,
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

	return matrix

}

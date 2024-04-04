package main

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/maxiim3/aoc2023/src/fileutils"
)

func parttwo() {
	debug.FreeOSMemory()
	debug.SetMemoryLimit(4096 * 1 << 50)
	fmt.Println("")
	fmt.Println("\tPart TWO Let's go!")
	fmt.Println("\t==================")
	fmt.Println("")

	f := fileutils.ParseFile("./src/day3/input.txt")

	matrix := createMatrix(f)

	colsize, rowsize := len(matrix), len(matrix[0])
	fmt.Printf("\t Size matrix %dx%d\n", rowsize, colsize)

	total := uint64(0)
	var wg sync.WaitGroup
	wg.Add(144)

	gears := []Gear{}

	for k, row := range matrix {
		go func(k int, row Row) {
			defer wg.Done()
			for v, data := range row {
				if data.value == '*' {
					data.isSymbol = true
					partNumbers := getPartNumbers(&matrix, k, v, colsize, rowsize)
					gear, isGear, gearRatio := createGear(data.x, data.y, partNumbers)

					if isGear {
						if len(gears) == 0 {
							gears = append(gears, gear)
						} else {
							for _, g := range gears {
								if g.x != gear.x && g.y != gear.y {
									gears = append(gears, gear)
									total += *gearRatio
								}
							}
						}
					}

				}
			}

		}(k, row)
		fmt.Printf("parsed row %d/%d, ", k, colsize)
	}

	fmt.Printf("\n TOTAL : %d\n", total)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func parseBothWays(row [140]Item, idx int, max int) int {
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

func parseForward(row [140]Item, idx int, max int) int {
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

func parseBackward(row [140]Item, idx int) int {
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
func parseNumbers(s Surrounding, matrix *Matrix, k int, v int, rowsize int) []uint {
	partNumbers := []uint{}
	if s.left {
		n := parseBackward((*matrix)[k], v-1)
		partNumbers = append(partNumbers, uint(n))
	}
	if s.right {
		n := parseForward((*matrix)[k], v+1, rowsize-1)
		partNumbers = append(partNumbers, uint(n))
	}

	if s.topcenter || s.topleft || s.topright {
		if s.topleft && s.topcenter && !s.topright {
			n := parseBackward((*matrix)[k-1], v)
			partNumbers = append(partNumbers, uint(n))
		} else if s.topcenter && s.topright && !s.topleft {
			n := parseForward((*matrix)[k-1], v, rowsize)
			partNumbers = append(partNumbers, uint(n))
		} else if s.topcenter && s.topright && s.topleft {
			n := parseBothWays((*matrix)[k-1], v, rowsize)
			partNumbers = append(partNumbers, uint(n))
		} else if s.topcenter && !s.topleft && !s.topright {
			numstr := fmt.Sprint((*matrix)[k-1][v].value)
			num, err := strconv.Atoi(numstr)
			if err != nil {
				errmsg := fmt.Sprint("Something went wrong with line above", k, v)
				panic(errmsg)
			}
			partNumbers = append(partNumbers, uint(num))
		} else if !s.topcenter && s.topleft && !s.topright {
			n := parseBackward((*matrix)[k-1], v-1)
			partNumbers = append(partNumbers, uint(n))
		} else if !s.topcenter && s.topright && !s.topleft {
			n := parseForward((*matrix)[k-1], v+1, rowsize)
			partNumbers = append(partNumbers, uint(n))
		}
	}
	if s.bottomcenter || s.bottomleft || s.bottomright {
		if s.bottomleft && s.bottomcenter && !s.bottomright {
			n := parseBackward((*matrix)[k+1], v)
			partNumbers = append(partNumbers, uint(n))
		} else if s.bottomcenter && s.bottomright && !s.bottomleft {
			n := parseForward((*matrix)[k+1], v, rowsize)
			partNumbers = append(partNumbers, uint(n))
		} else if s.bottomcenter && s.bottomright && s.bottomleft {
			n := parseBothWays((*matrix)[k+1], v, rowsize)
			partNumbers = append(partNumbers, uint(n))
		} else if s.bottomcenter && !s.bottomleft && !s.bottomright {
			numstr := fmt.Sprint((*matrix)[k+1][v].value)
			num, err := strconv.Atoi(numstr)
			if err != nil {
				errmsg := fmt.Sprint("Something went wrong with  line bellow", k, v)
				panic(errmsg)
			}
			partNumbers = append(partNumbers, uint(num))
		} else if !s.bottomcenter && s.bottomleft && !s.bottomright {
			n := parseBackward((*matrix)[k+1], v-1)
			partNumbers = append(partNumbers, uint(n))
		} else if !s.bottomcenter && s.bottomright && !s.bottomleft {
			n := parseForward((*matrix)[k+1], v+1, rowsize)
			partNumbers = append(partNumbers, uint(n))
		}
	}
	return partNumbers
}

func parseSurrounding(matrix *Matrix, k int, v int, rowsize int, colsize int) Surrounding {
	s := Surrounding{
		topleft:      false,
		topcenter:    false,
		topright:     false,
		left:         false,
		right:        false,
		bottomleft:   false,
		bottomcenter: false,
		bottomright:  false,
	}

	if k == 0 {
		switch {
		case v == 0:
			s.right = (*matrix)[k][v+1].isInt
			s.bottomcenter = (*matrix)[k+1][v].isInt
			s.bottomright = (*matrix)[k+1][v+1].isInt

		case v == rowsize-1:
			s.left = (*matrix)[k][v-1].isInt
			s.bottomleft = (*matrix)[k+1][v-1].isInt
			s.bottomcenter = (*matrix)[k+1][v].isInt

		case v < rowsize-1:
			s.left = (*matrix)[k][v-1].isInt
			s.right = (*matrix)[k][v+1].isInt
			s.bottomleft = (*matrix)[k+1][v-1].isInt
			s.bottomcenter = (*matrix)[k+1][v].isInt
			s.bottomright = (*matrix)[k+1][v+1].isInt
		case v == rowsize-1:
			s.topleft = (*matrix)[k-1][v-1].isInt
			s.topcenter = (*matrix)[k-1][v].isInt
			s.left = (*matrix)[k][v-1].isInt
		}
	} else if k == colsize-1 && v < rowsize-1 && v != 0 {
		s.topleft = (*matrix)[k-1][v-1].isInt
		s.topcenter = (*matrix)[k-1][v].isInt
		s.topright = (*matrix)[k-1][v+1].isInt
		s.left = (*matrix)[k][v-1].isInt
		s.right = (*matrix)[k][v+1].isInt
	} else if v == 0 {
		s.topcenter = (*matrix)[k-1][v].isInt
		s.topright = (*matrix)[k-1][v+1].isInt
		s.right = (*matrix)[k][v+1].isInt
		s.bottomcenter = (*matrix)[k+1][v].isInt
		s.bottomright = (*matrix)[k+1][v+1].isInt
	} else if v == colsize-1 {
		s.topleft = (*matrix)[k-1][v-1].isInt
		s.topcenter = (*matrix)[k-1][v].isInt
		s.left = (*matrix)[k][v-1].isInt
		s.bottomleft = (*matrix)[k+1][v-1].isInt
		s.bottomcenter = (*matrix)[k+1][v].isInt
	} else {
		s.topleft = (*matrix)[k-1][v-1].isInt
		s.topcenter = (*matrix)[k-1][v].isInt
		s.topright = (*matrix)[k-1][v+1].isInt
		s.left = (*matrix)[k][v-1].isInt
		s.right = (*matrix)[k][v+1].isInt
		s.bottomleft = (*matrix)[k+1][v-1].isInt
		s.bottomcenter = (*matrix)[k+1][v].isInt
		s.bottomright = (*matrix)[k+1][v+1].isInt
	}

	return s
}

func getPartNumbers(matrix *Matrix, k int, v int, colsize int, rowsize int) []uint {
	var wg sync.WaitGroup

	wg.Add(1)

	var s Surrounding
	var partNumbers []uint
	surrChan := make(chan Surrounding, 1)
	partnumChan := make(chan []uint, 50)

	go func() {
		fmt.Println("1")
		s = parseSurrounding(matrix, k, v, rowsize, colsize)
		surrChan <- s
		fmt.Println("2")
		partNumbers = parseNumbers(<-surrChan, matrix, k, v, rowsize)
		partnumChan <- partNumbers
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("GOGOGO", partNumbers)
	return <-partnumChan

	// s := parseSurrounding(matrix, k, v, rowsize, colsize)
	// partNumbers := parseNumbers(s, matrix, k, v, rowsize)
	// return partNumbers
}

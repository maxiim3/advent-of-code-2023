package main

import "strconv"

type Item struct {
	x            int
	y            int
	value        any
	isInt        bool
	isSymbol     bool
	isPartNumber bool
}

type Surrounding struct {
	topleft      bool
	topcenter    bool
	topright     bool
	left         bool
	right        bool
	bottomleft   bool
	bottomcenter bool
	bottomright  bool
}

func (s *Surrounding) hasAdjascentNumber() bool {
	if s.bottomcenter || s.bottomleft || s.bottomright || s.left || s.right || s.topcenter || s.topleft || s.topright {
		return true
	} else {
		return false
	}
}

func (s *Surrounding) hasAdjascentSymbol() bool {
	if s.bottomcenter || s.bottomleft || s.bottomright || s.left || s.right || s.topcenter || s.topleft || s.topright {
		return true
	} else {
		return false
	}
}

type LineNumbers struct {
	i          string
	hasPartNum bool
}

func (l *LineNumbers) convertToNumber() int {
	val, err := strconv.Atoi(l.i)
	if err != nil {
		panic("some number in line... isn't a number...")
	}
	return val
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Line struct {
	Index     int
	Bluecube  Cube
	Redcube   Cube
	Greencube Cube
	Draws     int
}

func CreateLine(line string) Line {
	gamedraw := strings.Split(line, ": ")
	if len(gamedraw) != 2 {
		panic(fmt.Sprintf("Error while parsing the line %s\n", line))
	}

	index := strings.Replace(gamedraw[0], "Game ", "", 1)
	idx, _ := strconv.Atoi(index)
	draws := strings.Split(gamedraw[1], "; ")

	blueBoxes := CreateCube(Blue)
	redBoxes := CreateCube(Red)
	greenBoxes := CreateCube(Green)

	for _, draw := range draws {
		boxes := strings.Split(draw, ", ")

		for _, box := range boxes {
			split := strings.Split(box, " ")
			stringValue, color := split[0], split[1]
			value, err := strconv.Atoi(stringValue)

			if err != nil {
				panic("The box value must be a number")
			}

			switch Color(color) {
			case "blue":
				blueBoxes.UpdateCube(value)
			case "red":
				redBoxes.UpdateCube(value)
			case "green":
				greenBoxes.UpdateCube(value)
			}
		}
	}

	return Line{
		Index:     idx,
		Bluecube:  blueBoxes,
		Redcube:   redBoxes,
		Greencube: greenBoxes,
		Draws:     len(draws),
	}
}

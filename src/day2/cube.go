package main

type Cube struct {
	Color Color
	Max   int
}

func CreateCube(color Color) Cube {
	return Cube{
		Color: Color(color),
		Max:   0,
	}
}

func (box *Cube) UpdateCube(value int) {
	if box.Max < value {
		box.Max = value
	}
}

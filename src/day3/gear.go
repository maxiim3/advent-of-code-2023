package main

type Gear struct {
	partNums []uint
	x        int
	y        int
}

// Is a gear if has exactly two adjascent part numbers
func (g *Gear) isGear() bool {
	return len(g.partNums) == 2
}

func (g *Gear) calculateGearRatio() uint64 {
	return uint64(g.partNums[0]) * uint64(g.partNums[1])
}

func createGear(x int, y int, partNumbers []uint) (gear Gear, isGear bool, gearRatio *uint64) {
	g := Gear{
		partNums: partNumbers,
		x:        x,
		y:        y,
	}
	valid := g.isGear()

	if valid {
		ratio := g.calculateGearRatio()
		return g, valid, &ratio
	}

	return g, valid, nil

}

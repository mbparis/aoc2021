package navigation

type Position struct {
	x int
	y int
}

func (p Position) check() int {
	return p.x * p.y
}

package navigation

type Position struct {
	x int
	y int
}

func (p Position) check() int {
	return p.x * p.y
}

func (p *Position) forward(units int) {
	p.x = p.x + units
}

func (p *Position) down(units int) {
	p.y = p.y + units
}

func (p *Position) up(units int) {
	p.y = p.y - units
}

package navigation

type Position struct {
	x int
	y int
}

func (p Position) Check() int {
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

type Bearing struct {
	Direction string
	Unit      int
}

type Course []*Bearing

func Chart(course Course) Position {
	position := Position{x: 0, y: 0}

	for _, bearing := range course {
		switch bearing.Direction {
		case "forward":
			position.forward(bearing.Unit)
		case "down":
			position.down(bearing.Unit)
		case "up":
			position.up(bearing.Unit)
		default:
			panic("What direction ?")
		}
	}

	return position

}

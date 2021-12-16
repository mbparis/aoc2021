package navigation

import "testing"

func TestPosition(t *testing.T) {
	expected := 0

	position := Position{x: 0, y: 0}

	actual := position.check()

	if expected != actual {
		t.Errorf("Position.check() = %v; want %v", actual, expected)

	}

}

func TestNavigate(t *testing.T) {
	expected := 150

	position := Position{x: 0, y: 0}

	position.forward(5)
	position.down(5)
	position.forward(8)
	position.up(3)
	position.down(8)
	position.forward(2)

	actual := position.check()

	if expected != actual {
		t.Errorf("Position.check() = %v; want %v", actual, expected)

	}

}

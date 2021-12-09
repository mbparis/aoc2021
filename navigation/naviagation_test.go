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

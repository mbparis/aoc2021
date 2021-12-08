package sonar

import "testing"

func TestSweep(t *testing.T) {
	inputSlice := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	expected := 7
	actual := Sweep(inputSlice)

	if expected != actual {
		t.Errorf("Sweep() = %v; want %v", actual, expected)
	}

}

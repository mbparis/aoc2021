package sonar

func Sweep(measurements []int) int {
	increases := 0
	for i := 1; i < len(measurements); i++ {
		curr := measurements[i]
		prev := measurements[i-1]

		diff := curr - prev

		if diff > 0 {
			increases += 1
		}
	}

	return increases
}

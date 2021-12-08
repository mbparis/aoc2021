package utils

import (
	"os"
	"strconv"
	"strings"
)

func ParseIntSlice(filePath string) ([]int, error) {
	// https://stackoverflow.com/a/9863218
	data, _ := os.ReadFile(filePath)

	lines := strings.Split(string(data), "\n")
	depths := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		depths = append(depths, n)

	}

	return depths, nil

}

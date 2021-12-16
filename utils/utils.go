package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/aoc2021/navigation"
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

func ParseCourse(filePath string) (navigation.Course, error) {

	data, _ := os.ReadFile(filePath)

	lines := strings.Split(string(data), "\n")
	course := make([]*navigation.Bearing, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineSplit := strings.Split(line, " ")
		unit, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			return nil, err
		}
		course = append(course, &navigation.Bearing{Direction: lineSplit[0], Unit: unit})
	}

	return course, nil

}

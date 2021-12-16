package main

import (
	"fmt"

	"github.com/aoc2021/navigation"
	"github.com/aoc2021/sonar"
	"github.com/aoc2021/utils"
)

func main() {
	fmt.Println("Hello, Advent of Code 2021!")
	day1Input, err := utils.ParseIntSlice("inputs/day1")
	if err != nil {
		return
	}
	fmt.Printf("Day1 answer: %d \n", sonar.Sweep(day1Input))

	day2input, err := utils.ParseCourse("inputs/day2")
	if err != nil {
		return
	}

	fmt.Printf("Day2 answer: %d \n", navigation.Chart(day2input).Check())
}

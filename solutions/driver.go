package solutions

import (
	"fmt"
	"os"

	"github.com/andey-robins/aoc-2023-go/solutions/day01"
)

func Day1() {
	testInput, err := os.ReadFile("./inputs/tests/01.txt")
	if err != nil {
		panic(err)
	}
	testInputb, err := os.ReadFile("./inputs/tests/01b.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("test part1: %v\n", day01.Part1(string(testInput)))
	fmt.Printf("test part2: %v\n", day01.Part2(string(testInputb)))

	input, err := os.ReadFile("./inputs/01.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %v\n", day01.Part1(string(input)))
	fmt.Printf("part2: %v\n", day01.Part2(string(input)))
}

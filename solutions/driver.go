package solutions

import (
	"fmt"
	"os"

	"github.com/andey-robins/aoc-2023-go/solutions/day01"
	"github.com/andey-robins/aoc-2023-go/solutions/day02"
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

func Day2() {
	testInput, err := os.ReadFile("./inputs/tests/02.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("test part1: %v\n", day02.Part1(string(testInput)))

	input, err := os.ReadFile("./inputs/02.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part1: %v\n", day02.Part1(string(input)))
}

func Day3() {
	// TODO: Implement Day3
}

func Day4() {
	// TODO: Implement Day4
}

func Day5() {
	// TODO: Implement Day5
}

func Day6() {
	// TODO: Implement Day6
}

func Day7() {
	// TODO: Implement Day7
}

func Day8() {
	// TODO: Implement Day8
}

func Day9() {
	// TODO: Implement Day9
}

func Day10() {
	// TODO: Implement Day10
}

func Day11() {
	// TODO: Implement Day11
}

func Day12() {
	// TODO: Implement Day12
}

func Day13() {
	// TODO: Implement Day13
}

func Day14() {
	// TODO: Implement Day14
}

func Day15() {
	// TODO: Implement Day15
}

func Day16() {
	// TODO: Implement Day16
}

func Day17() {
	// TODO: Implement Day17
}

func Day18() {
	// TODO: Implement Day18
}

func Day19() {
	// TODO: Implement Day19
}

func Day20() {
	// TODO: Implement Day20
}

func Day21() {
	// TODO: Implement Day21
}

func Day22() {
	// TODO: Implement Day22
}

func Day23() {
	// TODO: Implement Day23
}

func Day24() {
	// TODO: Implement Day24
}

func Day25() {
	// TODO: Implement Day25
}

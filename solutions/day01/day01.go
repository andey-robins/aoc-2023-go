package day01

import (
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input string) int {
	calibrationVals := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		calibrationRunes := make([]rune, 0)

		getFirstRune := func(line string) rune {
			for _, r := range line {
				if isNumeric(r) {
					return rune(r)
				}
			}
			return '0'
		}

		r := getFirstRune(line)
		calibrationRunes = append(calibrationRunes, r)
		lineRunes := []rune(line)
		slices.Reverse(lineRunes)
		r = getFirstRune(string(lineRunes))
		calibrationRunes = append(calibrationRunes, r)

		calibrationInt, err := strconv.Atoi(string(calibrationRunes))
		check(err)

		calibrationVals = append(calibrationVals, calibrationInt)
	}

	sum := 0
	for _, val := range calibrationVals {
		sum += val
	}
	return sum
}

func Part2(input string) int {

	calibrationVals := make([]int, 0)
	input = input + "\n"

	getFirstNumber := func(line *CalibrationLine) int {
		for _, element := range line.Element {
			if element.Number != nil && element.Number.Literal != nil {
				return element.Number.Literal.NumberLiteral
			} else if element.Number != nil && element.Number.Word != nil {
				if element.Number.Word.One != nil {
					return 1
				} else if element.Number.Word.Two != nil {
					return 2
				} else if element.Number.Word.Three != nil {
					return 3
				} else if element.Number.Word.Four != nil {
					return 4
				} else if element.Number.Word.Five != nil {
					return 5
				} else if element.Number.Word.Six != nil {
					return 6
				} else if element.Number.Word.Seven != nil {
					return 7
				} else if element.Number.Word.Eight != nil {
					return 8
				} else if element.Number.Word.Nine != nil {
					return 9
				} else if element.Number.Word.Zero != nil {
					return 0
				}
			}
		}
		return 0
	}

	parser, err := getParser()
	check(err)
	revParser, err := getParserReversed()
	check(err)

	inputLines := strings.Split(input, "\n")
	reversedInput := ""
	for _, line := range inputLines {
		lineRunes := []rune(line)
		slices.Reverse(lineRunes)
		reversedInput += string(lineRunes) + "\n"
	}
	reversedInput += "\n"

	calibration, err := parser.ParseString("", input)
	check(err)
	for _, line := range calibration.CalibrationLines {
		calibrationVals = append(calibrationVals, getFirstNumber(line))
	}

	calibration, err = revParser.ParseString("", reversedInput)
	check(err)
	for i, line := range calibration.CalibrationLines {
		if i < len(calibrationVals) {
			calibrationVals[i] = calibrationVals[i]*10 + getFirstNumber(line)
		}
	}

	sum := 0
	for _, val := range calibrationVals {
		sum += val
	}

	return sum
}

func isNumeric(x rune) bool {
	return x >= '0' && x <= '9'
}

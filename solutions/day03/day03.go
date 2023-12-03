package day03

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input string) int {
	schematicParser, err := getParser()
	check(err)

	schematic, err := schematicParser.ParseString("", input)
	check(err)
	schematic.CalculateValueLengths()
	schematic.CalculatePositions()

	partNumbersSum := 0

	for _, el := range schematic.EngineLines {
		for _, e := range el.Elements {
			if e.Value != nil {
				adjacentToValue := schematic.GetAdjacentElements(e)

				for _, a := range adjacentToValue {
					if a.Symbol != "" {
						partNumbersSum += e.Value.Number
						break
					}
				}
			}
		}
	}

	return partNumbersSum
}

func Part2(input string) int {
	schematicParser, err := getParser()
	check(err)

	schematic, err := schematicParser.ParseString("", input)
	check(err)
	schematic.CalculateValueLengths()
	schematic.CalculatePositions()

	gearRatioSum := 0

	for _, el := range schematic.EngineLines {
		for _, e := range el.Elements {
			if e.Symbol == "*" {
				adjacentElements := schematic.GetAdjacentElements(e)
				adjacentPartNumbers := make([]*Element, 0)

				for _, a := range adjacentElements {
					if a.Value != nil {
						adjacentPartNumbers = append(adjacentPartNumbers, a)
					}
				}

				if len(adjacentPartNumbers) == 2 {
					gearRatioSum += adjacentPartNumbers[0].Value.Number * adjacentPartNumbers[1].Value.Number
				}
			}
		}
	}

	return gearRatioSum
}

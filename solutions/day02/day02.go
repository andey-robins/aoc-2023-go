package day02

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input string) int {
	parser, err := getParser()
	check(err)

	games, err := parser.ParseString("", input)
	check(err)

	max := make(map[string]int, 0)
	max["red"] = 12
	max["blue"] = 14
	max["green"] = 13

	excludedGames := make([]int, 0)

	for _, game := range games.Games {
		for _, draw := range game.Draws {
			for _, cube := range draw.Cubes {
				color := cube.Color.cubeColorString()
				if cube.Count > max[color] {
					excludedGames = append(excludedGames, game.Id)
				}
			}
		}
	}

	sum := 0
	for _, game := range games.Games {
		if !contains(excludedGames, game.Id) {
			sum += game.Id
		}
	}

	return sum
}

func Part2(input string) int {
	parser, err := getParser()
	check(err)

	games, err := parser.ParseString("", input)
	check(err)

	powerSum := 0

	for _, game := range games.Games {
		max := make(map[string]int, 0)
		max["red"] = 0
		max["blue"] = 0
		max["green"] = 0

		for _, draw := range game.Draws {
			for _, cube := range draw.Cubes {
				color := cube.Color.cubeColorString()
				if cube.Count > max[color] {
					max[color] = cube.Count
				}
			}
		}

		powerSum += max["red"] * max["blue"] * max["green"]
	}

	return powerSum
}

func (c *Color) cubeColorString() string {
	if c.Red {
		return "red"
	} else if c.Blue {
		return "blue"
	} else if c.Green {
		return "green"
	}
	// never to be encountered unless we expand the grammar
	return ""
}

func contains(ll []int, i int) bool {
	for _, v := range ll {
		if v == i {
			return true
		}
	}
	return false
}

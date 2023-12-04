package day04

import (
	"math"

	"github.com/alecthomas/participle/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1(input string) int {
	cardParser, err := participle.Build[CardPile]()
	check(err)

	cardPile, err := cardParser.ParseString("", input)
	check(err)

	in := func(ll []int, vv int) bool {
		for _, v := range ll {
			if v == vv {
				return true
			}
		}

		return false
	}

	value := 0
	for _, c := range cardPile.Cards {
		cardWinners := c.getWinners()
		cardNumbers := c.getNums()
		winnersOnThisCard := make([]int, 0)

		for _, n := range cardNumbers {
			if in(cardWinners, n) {
				winnersOnThisCard = append(winnersOnThisCard, n)
			}
		}

		if len(winnersOnThisCard) != 0 {
			cardValue := int(math.Round(math.Pow(2, float64(len(winnersOnThisCard)-1))))
			value += cardValue
		}
	}

	return value
}

func Part2(input string) int {
	cardParser, err := participle.Build[CardPile]()
	check(err)

	cardPile, err := cardParser.ParseString("", input)
	check(err)

	cardCounts := make(map[int]int)
	for i := 1; i <= len(cardPile.Cards); i++ {
		// we can traverse this way because the cards are provided in order
		// from 1..n
		cardCounts[i] = 1
	}

	in := func(ll []int, vv int) bool {
		for _, v := range ll {
			if v == vv {
				return true
			}
		}

		return false
	}

	for _, c := range cardPile.Cards {
		cardId := c.CardNum

		// calculate the winners on this card before determining which cards we win
		cardWinners := c.getWinners()
		cardNumbers := c.getNums()
		winningNumbersOnThisCard := 0
		for _, n := range cardNumbers {
			if in(cardWinners, n) {
				winningNumbersOnThisCard++
			}
		}

		for i := cardId + 1; i <= winningNumbersOnThisCard+cardId; i++ {
			cardCounts[i] += cardCounts[cardId]
		}
	}

	// count the total number of cards
	totalCards := 0
	for _, v := range cardCounts {
		totalCards += v
	}

	return totalCards
}

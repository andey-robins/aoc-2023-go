package day04

type CardPile struct {
	Cards []*Card `@@*`
}

type Card struct {
	CardNum int       `"Card" @Int ":"`
	Winners []*Number `@@*`
	Numbers []*Number `"|" @@*`
}

type Number struct {
	Num int `@Int`
}

func (c *Card) getWinners() []int {
	winners := make([]int, 0)

	for _, w := range c.Winners {
		winners = append(winners, w.Num)
	}

	return winners
}

func (c *Card) getNums() []int {
	nums := make([]int, 0)

	for _, n := range c.Numbers {
		nums = append(nums, n.Num)
	}

	return nums
}

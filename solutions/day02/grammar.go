package day02

type GameList struct {
	Games []*Game `@@*`
}

type Game struct {
	Id    int     `Game @Int Colon`
	Draws []*Draw `@@*`
}

type Draw struct {
	Cubes []*CubeDraw `@@* Next`
}

type CubeDraw struct {
	Count int    `@Int`
	Color *Color `@@`
	Comma bool   `@Comma?`
}

type Color struct {
	Red   bool `  @"red"`
	Blue  bool `| @"blue"`
	Green bool `| @"green"`
}

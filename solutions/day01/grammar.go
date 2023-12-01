package day01

type CalibrationFile struct {
	CalibrationLines []*CalibrationLine `@@*`
}

type CalibrationLine struct {
	Element []*LineElement `@@* Newline`
}

type LineElement struct {
	Char   *Char   `  @@`
	Number *Number `| @@`
}

type Number struct {
	Literal *NumberLiteral `  @@`
	Word    *NumberWord    `| @@`
}

type Char struct {
	Char string `@Letter`
}

type NumberLiteral struct {
	NumberLiteral int `@Int`
}

type NumberWord struct {
	One   *NOne   `  @@`
	Two   *NTwo   `| @@`
	Three *NThree `| @@`
	Four  *NFour  `| @@`
	Five  *NFive  `| @@`
	Six   *NSix   `| @@`
	Seven *NSeven `| @@`
	Eight *NEight `| @@`
	Nine  *NNine  `| @@`
	Zero  *NZero  `| @@`
}

// NUMBERS //

type NOne struct {
	One string `@One`
}

type NTwo struct {
	Two string `@Two`
}

type NThree struct {
	Three string `@Three`
}

type NFour struct {
	Four string `@Four`
}

type NFive struct {
	Five string `@Five`
}

type NSix struct {
	Six string `@Six`
}

type NSeven struct {
	Seven string `@Seven`
}

type NEight struct {
	Eight string `@Eight`
}

type NNine struct {
	Nine string `@Nine`
}

type NZero struct {
	Zero string `@"zero"`
}

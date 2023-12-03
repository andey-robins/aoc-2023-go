package day03

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func getParser() (*participle.Parser[EngineSchematic], error) {
	configLexer := lexer.MustSimple([]lexer.SimpleRule{
		{"Number", `\d+`},
		{"Dot", `\.`},
		{"Newline", `\n`},
		{"Symbol", `[*#\+\-\$\&\^\/\=@\%]`},
	})

	parser, err := participle.Build[EngineSchematic](
		participle.Lexer(configLexer),
	)

	return parser, err
}

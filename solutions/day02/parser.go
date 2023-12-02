package day02

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func getParser() (*participle.Parser[GameList], error) {
	configLexer := lexer.MustSimple([]lexer.SimpleRule{
		{"Game", `Game`},
		{"Int", `[0-9]+`},
		{"Id", `[a-zA-Z]+`},
		{"Next", `[;\n]`},
		{"whitespace", `[ ]+`},
		{"Colon", `:`},
		{"Comma", `,`},
	})

	parser, err := participle.Build[GameList](
		participle.Lexer(configLexer),
	)

	return parser, err
}

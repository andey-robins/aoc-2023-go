package day01

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func getParser() (*participle.Parser[CalibrationFile], error) {
	configLexer := lexer.MustSimple([]lexer.SimpleRule{
		{"One", `one`},
		{"Two", `two`},
		{"Three", `three`},
		{"Four", `four`},
		{"Five", `five`},
		{"Six", `six`},
		{"Seven", `seven`},
		{"Eight", `eight`},
		{"Nine", `nine`},
		{"Zero", `zero`},
		{"Int", `[0-9]`},
		{"Letter", `[a-zA-Z]`},
		{"Newline", `\n`},
	})

	parser, err := participle.Build[CalibrationFile](
		participle.Lexer(configLexer),
	)

	return parser, err
}

func getParserReversed() (*participle.Parser[CalibrationFile], error) {
	configLexer := lexer.MustSimple([]lexer.SimpleRule{
		{"One", `eno`},
		{"Two", `owt`},
		{"Three", `eerht`},
		{"Four", `ruof`},
		{"Five", `evif`},
		{"Six", `xis`},
		{"Seven", `neves`},
		{"Eight", `thgie`},
		{"Nine", `enin`},
		{"Zero", `orez`},
		{"Int", `[0-9]`},
		{"Letter", `[a-zA-Z]`},
		{"Newline", `\n`},
	})

	parser, err := participle.Build[CalibrationFile](
		participle.Lexer(configLexer),
	)

	return parser, err
}

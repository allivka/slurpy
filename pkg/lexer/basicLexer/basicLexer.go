package basiclexer

import (
	"github.com/allivka/slurpy/pkg/formatter"
)

type Keyword interface {
	ProcessBlock(string) (string, error)
}

type Keywords = map[string]Keyword

func Lex(src formatter.WordSlice) {
}

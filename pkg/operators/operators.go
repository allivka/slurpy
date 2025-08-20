package operators

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
)

var OperatorTokens = bts.TokenMap {
	"+": Plus{},
	"-": Minus{},
	"++": Increment{},
	"--": Decrement{},
	"*": Multiply{},
	"**": Power{},
	"/": Divide{},
	"%": Module{},
	";": Eol{},
}

var BracketTokens = bts.TokenMap {
	"(": BracketCircleOpen{},
	")": BracketCircleClose{},
	"[": BracketSquareOpen{},
	"]": BracketSquareClose{},
	"{": BracketFigureOpen{},
	"}": BracketFigureClose{},
}

var TokenMaps = []bts.TokenMap{OperatorTokens, BracketTokens}

func init() {
	for _, tm := range TokenMaps {
		for k, v := range tm {
			tm[k], _ = bts.InitToken(v, k)
		}
	}
}

type Plus struct {
	bts.BasicToken
}

type Minus struct {
	bts.BasicToken
}

type Increment struct {
	bts.BasicToken
}

type Decrement struct {
	bts.BasicToken
}

type Multiply struct {
	bts.BasicToken
}

type Power struct {
	bts.BasicToken
}

type Divide struct {
	bts.BasicToken
}

type Module struct {
	bts.BasicToken
}

type Eol struct {
	bts.BasicToken
}

type BracketCircleOpen struct {
	bts.BasicToken
}

type BracketCircleClose struct {
	bts.BasicToken
}

type BracketSquareOpen struct {
	bts.BasicToken
}

type BracketSquareClose struct {
	bts.BasicToken
}

type BracketFigureOpen struct {
	bts.BasicToken
}

type BracketFigureClose struct {
	bts.BasicToken
}

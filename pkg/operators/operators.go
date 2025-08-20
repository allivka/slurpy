package operators

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
)

var OperatorTokens = bts.TokenMap {
	"+": Plus{},
	"-": Minus{},
	"*": Multiply{},
	"/": Divide{},
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

type Plus struct {
	bts.UncreatableToken
}

type Minus struct {
	bts.UncreatableToken
}

type Multiply struct {
	bts.UncreatableToken
}

type Divide struct {
	bts.UncreatableToken
}

type Eol struct {
	bts.UncreatableToken
}

type BracketCircleOpen struct {
	bts.UncreatableToken
}

type BracketCircleClose struct {
	bts.UncreatableToken
}

type BracketSquareOpen struct {
	bts.UncreatableToken
}

type BracketSquareClose struct {
	bts.UncreatableToken
}

type BracketFigureOpen struct {
	bts.UncreatableToken
}

type BracketFigureClose struct {
	bts.UncreatableToken
}

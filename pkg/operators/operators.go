package operators

import (
	"github.com/allivka/slurpy/pkg/tokens"
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
	tokens.UncreatableToken
}

type Minus struct {
	tokens.UncreatableToken
}

type Multiply struct {
	tokens.UncreatableToken
}

type Divide struct {
	tokens.UncreatableToken
}

type Eol struct {
	tokens.UncreatableToken
}

type BracketCircleOpen struct {
	tokens.UncreatableToken
}

type BracketCircleClose struct {
	tokens.UncreatableToken
}

type BracketSquareOpen struct {
	tokens.UncreatableToken
}

type BracketSquareClose struct {
	tokens.UncreatableToken
}

type BracketFigureOpen struct {
	tokens.UncreatableToken
}

type BracketFigureClose struct {
	tokens.UncreatableToken
}

package operators

import (
	"github.com/allivka/slurpy/pkg/tokens"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
)

var OperatorTokens = bts.TokenMap {
	"+": OperatorPlus{},
	"-": OperatorMinus{},
	";": OperatorMinus{},
}

var BracketTokens = bts.TokenMap {
	"(": BracketCircleOpen{},
	")": BracketCircleClose{},
	"[": BracketSquareOpen{},
	"]": BracketSquareClose{},
	"{": BracketFigureOpen{},
	"}": BracketFigureClose{},
}

type OperatorPlus struct {
	tokens.UncreatableToken
}

type OperatorMinus struct {
	tokens.UncreatableToken
}

type OperatorEol struct {
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

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


type OperatorPlus struct {
	tokens.UncreatableToken
}

type OperatorMinus struct {
	tokens.UncreatableToken
}

type OperatorEol struct {
	tokens.UncreatableToken
}



package basicstatements

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/lexer"
)

const (
	statementParameterTokenOpenWord = '['
	statementParameterTokenCloseWord = ']'
)

var (
	StatementParametersAssertionsSeparators = bts.TokenSlice{
		lexer.SpecifiedTokens[","],
	}

	StatementParametersAssertionsPartsSeparators = bts.TokenSlice{
		lexer.SpecifiedTokens["="],
	}

)


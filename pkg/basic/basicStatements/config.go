package basicstatements

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/lexer"
)

const (
	statementParameterTokenOpenWord  = '['
	statementParameterTokenCloseWord = ']'
	statementBodyTokenOpenWord = '{'
	statementBodyTokenCloseWord = '}'
)

var (
	StatementParametersAssertionsSeparators = bts.TokenSlice{
		lexer.SpecifiedTokens[","],
	}

	StatementParametersAssertionsPartsSeparators = bts.TokenSlice{
		lexer.SpecifiedTokens["="],
	}
	
	StatementHeaderSeparators = bts.TokenSlice {
		lexer.SpecifiedTokens[";"],
	}
)

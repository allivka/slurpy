package basicstatements

import (
	wp "github.com/allivka/slurpy/pkg/words"
	"github.com/allivka/slurpy/pkg/lexer"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	bps "github.com/allivka/slurpy/pkg/basic/basicParser"
	"errors"
	"fmt"
)

const (
	StatementParameterTokenOpenWord = '['
	StatementParameterTokenCloseWord = ']'
)

type Context = interface{}

type Statement interface {
	Execute(any) (error)
	NewFromTokenSlice(wp.WordSlice) (Statement, error)
}

type BasicStatementCreationContext struct {
	identificator bts.Token
	parameters map[bts.Token]bts.Token
	header bts.TokenSlice
	body []BasicStatementCreationContext
	context Context
}

func TokensToBasicStatementContext(src bts.TokenSlice, defaultIdentificator bts.Token, defaultBodyIdentificator bts.Token, statementIdentificators map[string]bts.Token) (context BasicStatementCreationContext, err error) {
	if src == nil || len(src) == 0 {
		return BasicStatementCreationContext{}, errors.New("Invalid src parameter for creation of BasicStatementCreationContext for later statement creation")
	}
	
	if defaultIdentificator == nil {
		return BasicStatementCreationContext{}, errors.New("Default identificator token for statement processing")
	}
	
	if defaultBodyIdentificator == nil {
		defaultBodyIdentificator = defaultIdentificator
	}
	
	var (
		identificator bts.Token
		ok bool
	)
	
	if identificator, ok = statementIdentificators[src[0].GetWord()]; !ok {
		identificator = defaultIdentificator
	}
	
	context.identificator = identificator
	
	offset, parameterSlice, err := bps.ParseBlockBetween(src[1:], lexer.SpecifiedTokens[StatementParameterTokenOpenWord], lexer.SpecifiedTokens[StatementParameterTokenCloseWord])
	
	if err != nil {
		return nil, fmt.Errorf("Failed turning tokens slice into statement creation context, failed parsing statement parameters: %w", err)
	}
	
	return
}


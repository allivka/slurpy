package basicstatements

import (
	"errors"
	"fmt"
	bps "github.com/allivka/slurpy/pkg/basic/basicParser"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/lexer"
	wp "github.com/allivka/slurpy/pkg/words"
)

type Context = interface{}

type Statement interface {
	Execute(any) error
	NewFromTokenSlice(wp.WordSlice) (Statement, error)
}

type BasicStatementCreationContext struct {
	identificator bts.Token
	parameters    bps.Parameters
	header        bts.TokenSlice
	body          []BasicStatementCreationContext
	context       Context
}

func TokensToBasicStatementContext(src bts.TokenSlice, defaultIdentificator bts.Token, defaultBodyIdentificator bts.Token, defaultParameters bps.Parameters, statementIdentificators map[string]bts.Token) (context BasicStatementCreationContext, err error) {
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
		ok            bool
	)

	if identificator, ok = statementIdentificators[src[0].GetWord()]; !ok {
		identificator = defaultIdentificator
	}

	context.identificator = identificator

	offset, parameterSlice, err := bps.ParseBlockBetween(src[1:], lexer.SpecifiedTokens[string(statementParameterTokenOpenWord)], lexer.SpecifiedTokens[string(statementParameterTokenCloseWord)])

	if err != nil {
		return BasicStatementCreationContext{}, fmt.Errorf("Failed turning tokens slice into statement creation context, failed parsing statement parameters: %w", err)
	}

	if offset != 0 {
		context.parameters = defaultParameters
	} else {
		parameters, err := bps.ParameterizeBlock(parameterSlice, StatementParametersAssertionsSeparators, StatementParametersAssertionsPartsSeparators, false, func(x int) string {
			return fmt.Sprint(x)
		})

		if err != nil {
			return BasicStatementCreationContext{}, fmt.Errorf("Failed turning source tokens into statement creation context, could not parameterize statements arguments: %w", err)
		}
	}

	return
}

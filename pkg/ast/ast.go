package ast

import (
	wp "github.com/allivka/slurpy/pkg/words"
	tp "github.com/allivka/slurpy/pkg/tokens"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"errors"
)

type Context = interface{}

type Keyword = tp.Identificator

type Statement interface {
	Execute(any) (error)
	NewFromTokenSlice(wp.WordSlice) (Statement, error)
}

type BasicStatementCreationContext struct {
	Keyword Keyword
	parameters map[bts.Token]bts.Token
	header bts.TokenSlice
	body []Statement
}

func TokensToBasicStatementContext(src bts.TokenSlice, defaultKeyword Keyword) (context BasicStatementCreationContext, err error) {
	if src == nil || len(src) == 0 {
		return BasicStatementCreationContext{}, errors.New("Invalid src parameter for creation of BasicStatementCreationContext for later statement creation")
	}
	
	var keyword Keyword
	
	if src[0].GetWordType() != wp.Identificator {
		keyword = defaultKeyword
	} else if keyword, ok := src[0].(Keyword); !ok {
		return BasicStatementCreationContext{}, errors.New("The first token of src parameter with word type Identificator cannot be represented as identificator(Something unexpected happened)")
	}
	
	
	return
}

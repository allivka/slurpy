package lexer

import (
	"errors"
	"fmt"
	base "github.com/allivka/slurpy/pkg/basic/basicLexer"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	wp "github.com/allivka/slurpy/pkg/words"
	"github.com/allivka/slurpy/pkg/tokens"
	"github.com/allivka/slurpy/pkg/operators"
)


var SpecifiedTokens bts.TokenMap

func init() {
	SpecifiedTokens = bts.MergeTokenMaps(operators.OperatorTokens, operators.SingleOperatorTokens)
}

type tokenizer struct {}

func(tokenizer) TokenFromWord(word string) (token bts.Token, err error) {
	
	if len(word) == 0 {
		return nil, errors.New("Impossible to tokenize empty word")
	}
	
	if v, ok := SpecifiedTokens[word]; ok {
		token = v
		token, err = token.NewFromWord(word)
		
		if err != nil {
			return nil, fmt.Errorf("Failed tokenizing word '%s' as specified token:  %w", word, err)
		}
		
		return token, nil
	}
	
	wordType, err := wp.GetWordType(word)
	
	if err != nil || wordType == wp.Empty || wordType == wp.Invalid	 {
		return nil, fmt.Errorf("Failed tokenizing word '%s' is either unknown or invalid or empty:  %w", word, err)
	}
	
	switch wordType {
		case wp.Integer: token = tokens.Integer{}
		case wp.Float: token = tokens.Float{}
		case wp.Identificator: token = tokens.Identificator{}
	
		default: return nil, fmt.Errorf("Failed tokenizing word '%s' is either unknown or invalid or empty: %w", word, err)
		
	}
	
	bts.InitToken(&token, word)
	
	return token, nil
}

func Lex(src wp.WordSlice) (bts.TokenSlice, error) {
	
	return base.Lex(src, tokenizer{})
}
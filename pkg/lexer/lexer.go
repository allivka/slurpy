package lexer

import (
	"errors"
	"fmt"
	base "github.com/allivka/slurpy/pkg/basic/basicLexer"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/runes"
	"github.com/allivka/slurpy/pkg/tokens"
	"github.com/allivka/slurpy/pkg/operators"
)


var SpecifiedTokens bts.TokenMap

func init() {
	SpecifiedTokens = bts.MergeTokenMaps(operators.OperatorTokens)
}

func tokenizer(word string) (token bts.Token, err error) {
	
	if len(word) == 0 {
		return nil, errors.New("Impossible to tokenize empty word")
	}
	
	if v, ok := SpecifiedTokens[word]; ok {
		token = v
		err = token.NewFromWord(word)
		
		if err != nil {
			return nil, fmt.Errorf("Failed tokenizing word as specified token:  %w", err)
		}
		
		return token, nil
	}
	
	
}

func Lex(words []string) {
	
}
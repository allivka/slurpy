package basiclexer

import (
	"fmt"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/words"
)

type TokenDetector interface {
	TokenFromWord(word string) (bts.Token, error)
}

func Lex(src words.WordSlice, tokenizer TokenDetector) (result bts.TokenSlice, err error) {

	result = make(bts.TokenSlice, len(src))

	for i, v := range src {
		result[i], err = tokenizer.TokenFromWord(v)

		if err != nil {
			return nil, fmt.Errorf("failed basic lexing of words, failed getting token from word '%s': %w", v, err)
		}
	}

	return result, nil
}

package basiclexer

import (
	"fmt"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"github.com/allivka/slurpy/pkg/words"
)


type TokenDetector interface {
	TokenFromWord(word string) (bts.Token, error)
}

func Lex(src words.WordSlice, detector TokenDetector) (result bts.TokenSlice, err error) {
	
	result = make(bts.TokenSlice, len(src))
	
	for i, v := range src {
		result[i], err = detector.TokenFromWord(v)
		
		if err != nil {
			return nil, fmt.Errorf("Failed basic lexing of words, failed getting token from word: %w", err)
		}
	}
	
	return result, nil
}

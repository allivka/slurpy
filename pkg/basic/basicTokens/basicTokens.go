package basictokens

import (
	wp "github.com/allivka/slurpy/pkg/words"
)

type Token interface {
	NewFromWord(string) (Token, error)
	GetWord() string
	GetWordType() wp.WordType
}

type BasicToken struct {
	Word string
	WordType wp.WordType
}

func(token BasicToken) NewFromWord(word string) (_ Token, err error) {
	token.Word = word
	token.WordType, err = wp.GetWordType(word)
	
	if err != nil {
		return nil, err
	}
	
	return token, nil
}

func(token BasicToken) GetWord() string {
	return token.Word
}

func(token BasicToken) GetWordType() wp.WordType {
	return token.WordType
}

func InitToken(token *Token, word string) (err error) {
	
	temp, err := (*token).NewFromWord(word)
	
	if err != nil {
		return err
	}
	
	*token = temp
	
	return nil
}

type TokenSlice = []Token

type TokenMap = map[string]Token

func MergeTokenMaps(overlays ...TokenMap) (result TokenMap) {
	
	result = make(TokenMap)
	
	for _, overlay := range overlays {
		for k, v := range overlay {
			result[k] = v
		}
	}
	
	return
}


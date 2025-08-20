package basictokens

import (
)

type Token interface {
	NewFromWord(string) (Token, error)
	GetWord() string
}

type BasicToken struct {
	Word string
}

func(token BasicToken) NewFromWord(word string) (Token, error) {
	token.Word = word
	return token, nil
}

func(token BasicToken) GetWord() string {
	return token.Word
}

func InitToken(token Token, word string) Token {
	token.NewFromWord(word)
	return token
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


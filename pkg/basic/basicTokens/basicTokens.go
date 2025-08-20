package basictokens

import (
)

type Token interface {
	NewFromWord(string) error
}

type UncreatableToken struct {}
func(UncreatableToken) NewFromWord(string) error{return nil}

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


package basicparser

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"fmt"
	"container/list"
)

func ParseBlockBetween(src bts.TokenSlice, startToken, endToken bts.Token) (int, bts.TokenSlice, error){ 
	
	if len(src) < 2 {
		return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v as provided block token slice length is less then 2 elements", src, startToken, endToken)
	}
	
	var (
		blockOffset int
		start bts.Token
	)
	
	for ; start != startToken; blockOffset++ {
		start = src[blockOffset]
	}
	
	if start.GetWord() == "" {
		return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no starting token '%+v' was found", src, startToken, endToken, startToken)
	}
	
	if blockOffset == len(src) {
		return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no paired ending token '%+v' was found", src, startToken, endToken, endToken)
	}
	
	var bracketStack list.List
	bracketStack.Init()
	
	bracketStack.PushBack(start)
	
	for i := blockOffset + 1; i < len(src); i++ {
		
		if endToken == startToken && endToken == src[i] {
			bracketStack.Remove(bracketStack.Back())
			goto endCheck
		}
		
		if (src[i] == startToken || src[i] == endToken) && src[i] != bracketStack.Back().Value {
			bracketStack.Remove(bracketStack.Back())
			goto endCheck
		}
		
		bracketStack.PushBack(src[i])
		
		endCheck:
		if bracketStack.Len() == 0 {
			return blockOffset, src[blockOffset:i+1], nil
		}
		
	}
	
	return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no paired ending token '%+v' was found", src, startToken, endToken, endToken)
}

package basicparser

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"fmt"
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
	
	if blockOffset >= len(src) - 1 {
		return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no paired ending token '%+v' was found", src, startToken, endToken, endToken)
	}
	
	var openedCounter = 1
	
	for i := blockOffset + 1; i < len(src); i++ {
		
		switch src[i] {
		case endToken: openedCounter--
		case startToken: openedCounter++
		}
		
		if openedCounter == 0 {
			return blockOffset, src[blockOffset:i+1], nil
		}
		
		if openedCounter < 0 {
			return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no paired starting token '%+v' was found", src, startToken, endToken, startToken)
		}
		
	}
	
	return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v no paired ending token '%+v' was found", src, startToken, endToken, endToken)
}

package basicparser

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	"fmt"
)

type StatementParameters = map[string]bts.TokenSlice

func ParseBlockBetween(src bts.TokenSlice, startToken, endToken bts.Token) (int, bts.TokenSlice, error){ 
	
	if len(src) < 2 {
		return 0, nil, fmt.Errorf("Failed parsing block of tokens '%+v' in block between %+v and %+v as provided block token slice length is less then 2 elements", src, startToken, endToken)
	}
	
	if startToken == nil  || endToken == nil {
		return 0, nil, fmt.Errorf("Border token cannot be nil")
	}
	
	var (
		blockOffset int
		start bts.Token
	)
	
	for ; start.GetWord() != startToken.GetWord(); blockOffset++ {
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
		
		switch src[i].GetWord() {
		case endToken.GetWord(): openedCounter--
		case startToken.GetWord(): openedCounter++
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

func ParseBlockWithSeparators(src bts.TokenSlice, separatorTokens []bts.Token) (result []bts.TokenSlice, err error) {
	
	if src == nil {
		return nil, fmt.Errorf("Source tokens slice cannot be nil")
	}
	
	if separatorTokens == nil {
		return nil, fmt.Errorf("Separator tokens cannot be nil")
	}
	
	var start int
	
	separatorSet := map[string]struct{}{}
	
	for _, v := range separatorTokens {
		separatorSet[v.GetWord()] = struct{}{}
	}
	
	for end := 0; end < len(src); end++ {
		if _, ok := separatorSet[src[end].GetWord()]; ok {
			result = append(result, src[start:end])
			start = end + 1
		}
	}
	
	if start < len(src) {
		result = append(result, src[start:])
	}
	
	return result, nil
}

func ParameterizeBlock(src bts.TokenSlice, separatorTokens bts.TokenSlice, assertionTokens bts.TokenSlice, singleRightAssertionPart bool, defaultParameters StatementParameters, countParameterParser func(int)string) (result StatementParameters, err error) {
	if src == nil {
		return nil, fmt.Errorf("Source tokens slice cannot be nil")
	}
	
	assertions, err := ParseBlockWithSeparators(src, separatorTokens)
	
	if err != nil {
		return nil, fmt.Errorf("Could not parameterize block, error during separating assertions: %w", err)
	}
	
	if len(assertions) == 0 {
		return nil, fmt.Errorf("Could not parameterize block, no assertions received from the block")
	}
	
	result = StatementParameters{}
	
	for k, v := range defaultParameters {
		result[k] = v
	}
	
	var (
		assertionParts []bts.TokenSlice
	)
	
	for i, assertion := range assertions {
		assertionParts, err = ParseBlockWithSeparators(assertion, assertionTokens)
		
		if err != nil {
			return nil, fmt.Errorf("Could not parameterize block, assertion parts separating failed: %w", err)
		}
		
		if len(assertionParts) == 0 {
			return nil, fmt.Errorf("Could not parameterize block, empty assertion somehow")
		}
		
		if len(assertionParts) == 1 {
			result[countParameterParser(i)] = assertionParts[0]
			continue
		}
		
		if len(assertionParts) == 2 && len(assertionParts[0]) == 1 {
			if singleRightAssertionPart && len(assertionParts[1]) > 1 {
				return nil, fmt.Errorf("Could not parameterize block, few tokens on the right side of the assertion are not allowed: '%+v'", assertion)
			} else if len(assertionParts[1]) == 0 {
				return nil, fmt.Errorf("Could not parameterize block, empty right side of the assertion '%+v'", assertion)
			}
			
			result[assertionParts[0][0].GetWord()] = assertionParts[1]
			continue
		}
		
		return nil, fmt.Errorf("Could not parameterize block, invalid parameter settings syntax")
	}
	
	return
}

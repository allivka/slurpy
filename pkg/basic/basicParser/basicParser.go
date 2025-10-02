package basicparser

import (
	"fmt"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
)

type Parameters = map[string]bts.TokenSlice

func ParseBlockBetween(src bts.TokenSlice, startToken, endToken bts.Token) (int, bts.TokenSlice, error) {
	if len(src) < 2 {
		return 0, nil, fmt.Errorf("failed parsing block: token slice length must be at least 2")
	}

	if startToken == nil || endToken == nil {
		return 0, nil, fmt.Errorf("border token cannot be nil")
	}

	startIndex := -1
	for i, token := range src {
		if token.GetWord() == startToken.GetWord() {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		return -1, nil, fmt.Errorf("failed parsing block: start token '%s' not found", startToken.GetWord())
	}

	openedCounter := 1

	for i := startIndex + 1; i < len(src); i++ {
		currentWord := src[i].GetWord()

		switch currentWord {
		case endToken.GetWord():
			openedCounter--
		case startToken.GetWord():
			openedCounter++
		}

		if openedCounter == 0 {
			return startIndex, src[startIndex : i+1], nil
		}

		if openedCounter < 0 {
			return 0, nil, fmt.Errorf("failed parsing block: unbalanced tokens, too many end tokens '%s'", endToken.GetWord())
		}
	}

	return 0, nil, fmt.Errorf("failed parsing block: end token '%s' not found to match start token", endToken.GetWord())
}

func ParseBlockWithSeparators(src bts.TokenSlice, separatorTokens []bts.Token) (result []bts.TokenSlice, err error) {

	if src == nil {
		return nil, nil
	}

	if separatorTokens == nil {
		return []bts.TokenSlice{src}, nil
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

func ParameterizeBlock(src bts.TokenSlice, separatorTokens bts.TokenSlice, assertionTokens bts.TokenSlice, singleRightAssertionPart bool, countParameterParser func(int) string) (result Parameters, err error) {
	if src == nil {
		return nil, fmt.Errorf("source tokens slice cannot be nil")
	}

	assertions, err := ParseBlockWithSeparators(src, separatorTokens)

	if err != nil {
		return nil, fmt.Errorf("could not parameterize block, error during separating assertions: %w", err)
	}

	if len(assertions) == 0 {
		return nil, fmt.Errorf("could not parameterize block, no assertions received from the block")
	}

	result = Parameters{}

	var (
		assertionParts []bts.TokenSlice
	)

	for i, assertion := range assertions {
		assertionParts, err = ParseBlockWithSeparators(assertion, assertionTokens)

		if err != nil {
			return nil, fmt.Errorf("could not parameterize block, assertion parts separating failed: %w", err)
		}

		if len(assertionParts) == 0 {
			return nil, fmt.Errorf("could not parameterize block, empty assertion somehow")
		}

		if len(assertionParts) == 1 {
			result[countParameterParser(i)] = assertionParts[0]
			continue
		}

		if len(assertionParts) == 2 && len(assertionParts[0]) == 1 {
			if singleRightAssertionPart && len(assertionParts[1]) > 1 {
				return nil, fmt.Errorf("could not parameterize block, few tokens on the right side of the assertion are not allowed: '%+v'", assertion)
			} else if len(assertionParts[1]) == 0 {
				return nil, fmt.Errorf("could not parameterize block, empty right side of the assertion '%+v'", assertion)
			}

			result[assertionParts[0][0].GetWord()] = assertionParts[1]
			continue
		}

		return nil, fmt.Errorf("could not parameterize block, invalid parameter settings syntax")
	}

	return
}

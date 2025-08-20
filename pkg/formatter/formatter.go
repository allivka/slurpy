package formatter

import (
	"fmt"
	"strings"
	wp "github.com/allivka/slurpy/pkg/words"
)

type lineSlice = []string

func clearComments(src string) (result lineSlice) {
	
	for _, line := range strings.Split(src, "\n") {
		if !strings.HasPrefix(strings.TrimSpace(line), `//`) {
			result = append(result, line)
		}
	}
	
	return
}

func Format(t lineSlice) (lineSlice, error) {

	lines := []string{}

	for _, v := range t {
		line := strings.TrimSpace(v)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	var result string

	for i := 0; i < len(lines); i++ {

		line := lines[i]

		if !(line[len(line)-1] == '\\') {
			result += strings.Join(strings.Fields(line), " ")
			if i != len(lines) - 1 {
				result += "\n"
			}
			continue

		}

		if i == len(lines)-1 {
			return nil, fmt.Errorf("No line for concatenation with \\ operator at line %d", i+1)
		}

		lines[i+1] = strings.Join(strings.Fields(strings.TrimSpace(string(line[:len(line)-1])+lines[i+1])), " ")
	}
	
	result = strings.TrimSpace(result)
	
	return strings.Split(result, "\n"), nil
}

func separateWords(word string) (words wp.WordSlice, err error) {
	
	runes := []rune(word)
	
	lastType, err := wp.GetWordType(string(runes[0]))
	
	if err != nil {
		return nil, fmt.Errorf("Failed separating words: invalid word: %w", err)
	}
	
	
	var (
		currentType wp.WordType
		subWord string
		lastIdx int
	)
	
	for i := 1; i < len(runes) + 1; i++ {
		
		subWord = string(runes[lastIdx:i+1])
		
		currentType, _ = wp.GetWordType(subWord)
		
		
		// fmt.Printf("Subword: %s type: %d\n", subWord, currentType)
		
		
		if currentType == wp.Bracket || currentType != lastType{
			words = append(words, string(runes[lastIdx:i]))
			lastIdx = i
			if i < len(runes) { lastType, err = wp.GetWordType(string(runes[i])) }
			
			if err != nil {
				return nil, fmt.Errorf("Failed separating words: invalid word: %w", err)
			}
			// fmt.Println(words)
		}
		
	}
	
	return
}

func WordsFromSrcString(src string) (wp.WordSlice, error) {
	
	
	lines, err := Format(clearComments(src))
	
	if err != nil {
		return wp.WordSlice{}, err
	}
	
	var buff = wp.WordSlice{}
	
	for _, line := range lines {
		for _, word := range strings.Fields(line) {
			subWords, err := separateWords(word)
			
			if err != nil {
				return nil, err
			}
			
			buff = append(buff, subWords...)
		}
		buff = append(buff, ";")
	}
	
	words := make(wp.WordSlice, len(buff))
	
	copy(words, buff)
	
	return words, nil
	
}

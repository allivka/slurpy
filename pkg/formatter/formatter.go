package formatter

import (
	"fmt"
	"strings"
	rns "github.com/allivka/slurpy/pkg/runes"
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

func separateWords(word string) (words wp.WordSlice) {
	
	runes := []rune(word)
	
	var (
		lastType rns.RuneType = rns.GetRuneType(runes[0])
	
		lastIdx int = 0
	)
	
	for i, v := range runes {
		if rns.GetRuneType(v) == lastType {continue}
		words = append(words, string(runes[lastIdx:i]))
		lastType =rns.GetRuneType(v)
		lastIdx = i
	}
	
	words = append(words, string(runes[lastIdx:]))
	
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
			buff = append(buff, separateWords(word)...)
		}
		buff = append(buff, ";")
	}
	
	words := make(wp.WordSlice, len(buff))
	
	copy(words, buff)
	
	return words, nil
	
}

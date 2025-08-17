package formatter

import (
	"fmt"
	"strings"
	unc "unicode"
)

var operatorRunes = "+-"

type Words = []string

func Format(src string) (string, error) {

	t := strings.Split(src, "\n")

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
			return "", fmt.Errorf("No line for concatenation with \\ operator at line %d", i+1)
		}

		lines[i+1] = strings.Join(strings.Fields(strings.TrimSpace(string(line[:len(line)-1])+lines[i+1])), " ")
	}

	return strings.TrimSpace(result), nil
}

func separateWords(word string) (words Words) {
	
	const (
		digit = iota
		letter
		op
		unknown
	)
	
	type typeType = int
	
	getType := func(r rune) typeType {
		switch {
			case strings.Contains(operatorRunes, string(r)): return op
			case unc.IsDigit(r): return digit
			case unc.IsLetter(r): return digit
			default: return unknown
		}
	}
	
	runes := []rune(word)
	
	var (
		lastType typeType = getType(runes[0])
	
		lastIdx int = 0
	)
	
	for i, v := range runes {
		if getType(v) == lastType {continue}
		words = append(words, string(runes[lastIdx:i]))
		lastIdx = i
	}
	
	words = append(words, string(runes[lastIdx:]))
	
	return
}

func WordsFromString(src string) (Words, error) {
	src, err := Format(src)
	
	if err != nil {
		return Words{}, err
	}
	
	lines := strings.Split(src, "\n")
	
	var buff = Words{}
	
	for _, line := range lines {
		for _, word := range strings.Fields(line) {
			buff = append(buff, separateWords(word)...)
		}
		buff = append(buff, ";")
	}
	
	words := make(Words, len(buff))
	
	copy(words, buff)
	
	return words, nil
	
}

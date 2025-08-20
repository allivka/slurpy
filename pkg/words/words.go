package words

import (
	"fmt"
	"strconv"
	"unicode"
	rp "github.com/allivka/slurpy/pkg/runes"
)

type WordSlice = []string

type WordType = int

const (
	Empty = iota
	Integer
	Float
	Boolean
	Identificator
	Operator
	Bracket
	Invalid
)

func ValidateWords(words WordSlice) (err error) {
	
	var wt WordType
	
	for _, word := range words {
		wt, err = GetWordType(word)
		
		if wt == Invalid || wt == Empty || err != nil {
			return fmt.Errorf("Detected invalid word '%s': %w", word, err)
		}
	}
	
	return nil
}

func GetWordType(word string) (result WordType, err error) {
	
	if len(word) == 0 {
		return Empty, nil
	}
	
	_, err = strconv.ParseInt(word, 10, 64)
	
	if err == nil {
		return Integer, nil
	}
	
	_, err = strconv.ParseFloat(word, 64)
	
	if err == nil {
		return Float, nil
	}
	
	_, err = strconv.ParseBool(word)
	
	if err == nil {
		return Boolean, nil
	}
	
	runes := []rune(word)
	
	if unicode.IsDigit(runes[0]) {
		return Invalid, fmt.Errorf("Invalid word '%s' starts from a digit but can't be parsed as an integer or float", word)
	}
	
	isIdentificator := true
	
	for _, v:= range runes {
		if unicode.IsDigit(v) || unicode.IsLetter(v) || v == '_' {continue}
		isIdentificator = false
		break
	}
	
	if isIdentificator {
		return Identificator, nil
	}
	
	isOperator := true
	
	for _, v:= range runes {
		if rp.GetRuneType(v) == rp.Operator {continue}
		isOperator = false
		break
	}
	
	if isOperator {
		return Operator, nil
	}
	
	if len(runes) == 1 && rp.GetRuneType(runes[0]) == rp.Bracket {
		return Bracket, nil
	}
	
	return Invalid, fmt.Errorf("Word '%s' is invalid", word)
}

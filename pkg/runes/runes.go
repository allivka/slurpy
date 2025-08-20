package runes

import (
	"strings"
	"unicode"
	"github.com/allivka/slurpy/pkg/operators"
)

const (
	Digit = iota
	Letter
	Bracket
	Operator
	Unknown
)

func init() {
	for k := range operators.OperatorTokens {
		operatorRunes += k
	}
	for k := range operators.BracketTokens {
		bracketRunes += k
	}
}

var operatorRunes string
var bracketRunes string

type RuneType = int

func GetRuneType(r rune) RuneType {
	switch {
	case strings.Contains(bracketRunes, string(r)): return Bracket
		case strings.Contains(operatorRunes, string(r)): return Operator
		case unicode.IsDigit(r): return Digit
		case unicode.IsLetter(r): return Letter
		default: return Unknown
	}
}
package runes

import (
	"strings"
	"unicode"
	"github.com/allivka/slurpy/pkg/operators"
)

const (
	Digit = iota
	Letter
	Operator
	Unknown
)

func init() {
	for k := range operators.OperatorTokens {
		operatorRunes += k
	}
}

var operatorRunes string

type RuneType = int

func GetRuneType(r rune) RuneType {
	switch {
		case strings.Contains(operatorRunes, string(r)): return Operator
		case unicode.IsDigit(r): return Digit
		case unicode.IsLetter(r): return Letter
		default: return Unknown
	}
}
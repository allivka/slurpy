package runes

import (
	"strings"
	"unicode"
	"github.com/allivka/slurpy/pkg/operators"
)

const (
	Digit = iota
	Letter
	SingleOperator
	Operator
	IdentificatorRune
	Unknown
)

func init() {
	for k := range operators.OperatorTokens {
		operatorRunes += k
	}
	for k := range operators.SingleOperatorTokens {
		SingleOperatorRunes += k
	}
}

var operatorRunes string
var SingleOperatorRunes string

type RuneType = int

func GetRuneType(r rune) RuneType {
	switch {
	case unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_': return IdentificatorRune
		case strings.Contains(SingleOperatorRunes, string(r)): return SingleOperator
		case strings.Contains(operatorRunes, string(r)): return Operator
		case unicode.IsDigit(r): return Digit
		case unicode.IsLetter(r): return Letter
		default: return Unknown
	}
}
package keywords

import (
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
	tp "github.com/allivka/slurpy/pkg/tokens"
	"github.com/allivka/slurpy/pkg/ast"	
)


type Keyword interface {
	NewFromIdentificator(tp.Identificator) (Keyword, error)	
	ParseBlock(bts.TokenSlice) (ast.Ast, error)
}

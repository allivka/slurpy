package ast

import (
	wp "github.com/allivka/slurpy/pkg/words"
	tp "github.com/allivka/slurpy/pkg/tokens"
	bts "github.com/allivka/slurpy/pkg/basic/basicTokens"
)

type AstNode interface {
	Execute(any) (error)
	NewFromTokenSlice(wp.WordSlice) (AstNode, error)
} 

type Context interface{}

type Statement struct {
	keyword tp.Identificator
	parameters map[tp.Identificator]bts.Token
	header bts.TokenSlice
	body []Statement
	defaultStatementsKeyword tp.Identificator
	
}

func(s Statement) Execute(context any) (err error) {
	
	for _, statement := range s.body {
		 err = statement.Execute(context)
		 
		 if err != nil {
			return err
		 }
	}
	
	return nil
}

type Ast struct {
	Nodes []AstNode
}

func (tree Ast) Execute(context any) (err error) {
	
	for _, node := range tree.Nodes {
		 err = node.Execute(context)
		 
		 if err != nil {
			return err
		 }
	}
	
	return nil
	
}

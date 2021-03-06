// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package expressions

import (
	"base/docs"
	"datavalues"
)

type AliasedExpression struct {
	name string
	expr IExpression
}

func ALIASED(name string, expr IExpression) IExpression {
	return NewAliasedExpression(name, expr)
}

func NewAliasedExpression(name string, expr IExpression) *AliasedExpression {
	return &AliasedExpression{
		name: name,
		expr: expr,
	}
}

func (e *AliasedExpression) Eval(params IParams) (*datavalues.Value, error) {
	return e.expr.Eval(params)
}

func (e *AliasedExpression) Walk(visit Visit) error {
	return Walk(visit, e.expr)
}

func (e *AliasedExpression) String() string {
	return e.name
}

func (e *AliasedExpression) Document() docs.Documentation {
	return nil
}

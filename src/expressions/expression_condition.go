// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package expressions

import (
	"base/docs"
	"datavalues"
)

func LT(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "<",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Less than."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp == datavalues.LessThan), nil
		},
	}
}

func LTE(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "<=",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Less than or equal to."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp < datavalues.GreaterThan), nil
		},
	}
}

func EQ(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "=",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Equal."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp == datavalues.Equal), nil
		},
	}
}

func NEQ(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "!=",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Not equal."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp != datavalues.Equal), nil
		},
	}
}

func GT(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: ">",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Greater than."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp == datavalues.GreaterThan), nil
		},
	}
}

func GTE(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: ">=",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Greater than or equal to."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			cmp, err := datavalues.Compare(left, right)
			if err != nil {
				return nil, err
			}
			return datavalues.MakeBool(cmp > datavalues.LessThan), nil
		},
	}
}

func AND(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "AND",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Logic AND."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			return datavalues.ToValue(left.AsBool() && right.AsBool()), nil
		},
	}
}

func OR(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "OR",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("Logic OR."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			return datavalues.ToValue(left.AsBool() || right.AsBool()), nil
		},
	}
}

func LIKE(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "LIKE",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("LIKE."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			return datavalues.ToValue(datavalues.Like(right.AsString(), left)), nil
		},
	}
}

func NOT_LIKE(left interface{}, right interface{}) IExpression {
	exprs := expressionsFor(left, right)
	return &BinaryExpression{
		name: "NOT LIKE",
		argumentNames: [][]string{
			{"left", "right"},
		},
		description: docs.Text("NOT LIKE."),
		validate:    All(),
		left:        exprs[0],
		right:       exprs[1],
		evalFn: func(left *datavalues.Value, right *datavalues.Value) (*datavalues.Value, error) {
			return datavalues.ToValue(!datavalues.Like(right.AsString(), left)), nil
		},
	}
}

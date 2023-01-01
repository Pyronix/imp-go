package parser

import (
	"fmt"
	"strconv"
)

import . "imp/types"

type Parser struct {
	tokens   *Tape[Token]
	position int
}

func NewParser(input string) *Parser {
	return &Parser{
		tokens:   TokenizeString(input),
		position: 0,
	}
}

func (p *Parser) ParseProgram() (Expression, error) {
	return p.ParseExpression()
}

func (p *Parser) ParseExpression() (Expression, error) {
	return p.parseDisjunction()
}

type ExpressionParser func() (Expression, error)
type BinaryExpressionCollector func(left Expression, right Expression) Expression

func (p *Parser) parseBinaryOperation(operatorToken TokenType, higherPrecedenceParser ExpressionParser, collector BinaryExpressionCollector) (Expression, error) {
	sub, err := higherPrecedenceParser()
	if err != nil {
		return nil, err
	}

	terms := []Expression{sub}

	for p.tokens.Peek().Type == operatorToken {
		p.tokens.Next()
		sub, err := higherPrecedenceParser()
		if err != nil {
			return nil, err
		}

		terms = append(terms, sub)
	}

	if len(terms) > 1 {
		for i := len(terms) - 2; i >= 0; i-- {
			terms[i] = collector(terms[i], terms[i+1])
		}
	}

	return terms[0], nil
}

func (p *Parser) parseDisjunction() (Expression, error) {
	return p.parseBinaryOperation(OR, p.parseConjunction, func(left Expression, right Expression) Expression {
		return Or(left, right)
	})
}

func (p *Parser) parseConjunction() (Expression, error) {
	return p.parseBinaryOperation(AND, p.parseEquality, func(left Expression, right Expression) Expression {
		return And(left, right)
	})
}

func (p *Parser) parseEquality() (Expression, error) {
	return p.parseBinaryOperation(EQUAL, p.parseComparison, func(left Expression, right Expression) Expression {
		return Equal(left, right)
	})
}

func (p *Parser) parseComparison() (Expression, error) {
	return p.parseBinaryOperation(LESS, p.parseTerm, func(left Expression, right Expression) Expression {
		return Lesser(left, right)
	})
}

func (p *Parser) parseTerm() (Expression, error) {
	return p.parseBinaryOperation(ADD, p.parseFactor, func(left Expression, right Expression) Expression {
		return Plus(left, right)
	})
}

func (p *Parser) parseFactor() (Expression, error) {
	return p.parseBinaryOperation(MUL, p.parseUnary, func(left Expression, right Expression) Expression {
		return Mult(left, right)
	})
}

func (p *Parser) parseUnary() (Expression, error) {
	if p.tokens.Peek().Type == NOT {
		p.tokens.Next()
		unary, err := p.parseUnary()
		if err != nil {
			return nil, err
		}

		return Negation(unary), nil
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (Expression, error) {
	switch next := p.tokens.Next(); {
	case next.Type == INT:
		intVar, err := strconv.Atoi(next.Value)
		if err != nil {
			return nil, err
		}

		return NumberExpression(intVar), nil
	case next.Type == BOOL:
		if next.Value == "true" {
			return Bool(true), nil
		} else if next.Value == "false" {
			return Bool(false), nil
		} else {
			return nil, fmt.Errorf("invalid bool literal: %q", next.Value)
		}
	case next.Type == OPEN:
		exp, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}

		after := p.tokens.Next()

		if after.Type != CLOSE {
			return nil, fmt.Errorf("expected closing ), got: %q", after.Value)
		} else {
			return Grouping(exp), nil
		}
	case next.Type == IDENTIFIER:
		return Variable(next.Value), nil
	default:
		p.tokens.Rewind()
	}

	return nil, fmt.Errorf("exptected an primary expression, got %q", p.tokens.Peek())
}

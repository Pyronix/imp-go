package parser

import (
	"fmt"
	"strconv"
)

import . "imp/types"

type Parser struct {
	tokens   []Token
	position int
}

func NewParser(input string) *Parser {
	lexer, _ := NewLexer(input)

	return &Parser{
		tokens:   lexer.AllTokens(),
		position: 0,
	}
}

func (p *Parser) current() Token {
	return p.tokens[p.position]
}

func (p *Parser) next() Token {
	if p.position >= len(p.tokens) {
		return p.current()
	}

	next := p.current()
	p.position += 1
	if p.position >= len(p.tokens) {
		p.position = len(p.tokens) - 1
	}

	return next
}

func (p *Parser) rewind() Token {
	if p.position == 0 {
		return p.current()
	}

	p.position -= 1

	return p.current()
}

func (p *Parser) peek() Token {
	return p.current()
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

	for p.peek().Type == operatorToken {
		p.next()
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
	if p.peek().Type == NOT {
		p.next()
		unary, err := p.parseUnary()
		if err != nil {
			return nil, err
		}

		return Negation(unary), nil
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() (Expression, error) {
	switch next := p.next(); {
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

		after := p.next()

		if after.Type != CLOSE {
			return nil, fmt.Errorf("expected closing ), got: %q", after.Value)
		} else {
			return Grouping(exp), nil
		}
	case next.Type == IDENTIFIER:
		return Variable(next.Value), nil
	default:
		p.rewind()
	}

	return nil, fmt.Errorf("exptected an primary expression, got %q", p.peek())
}

package parser

import (
	"fmt"
	"io"
	"strconv"
)

import . "imp/types"
import . "imp/statements"

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

func NewParserFromReader(input io.Reader) *Parser {
	return &Parser{
		tokens:   TokenizeFromReader(input),
		position: 0,
	}
}

func (p *Parser) ParseProgram() ProgramStatement {
	statement := p.ParseSequence()

	if p.tokens.position != p.tokens.size-1 && p.tokens.Peek().Type != EOF {
		panic(fmt.Errorf("unexpected token %q encountered, expected syntax not matched", p.tokens.Peek()))
	}

	return Program(statement)
}

func (p *Parser) ParseBlock() BlockStatement {
	if p.tokens.Next().Type != BLOCKOPEN {
		p.tokens.Rewind()
		panic(fmt.Errorf("expected '{', got %q", p.tokens.Peek()))
	}

	statement := p.ParseSequence()

	if p.tokens.Next().Type != BLOCKCLOSE {
		p.tokens.Rewind()
		panic(fmt.Errorf("expected '}', got %q", p.tokens.Peek()))
	}

	return Block(statement)
}

func (p *Parser) ParseSequence() Statement {
	statement := p.ParseStatement()

	stmts := []Statement{statement}

	for p.tokens.Peek().Type == SEMICOLON {
		p.tokens.Next()

		statement := p.ParseStatement()

		stmts = append(stmts, statement)
	}

	if len(stmts) > 1 {
		for i := len(stmts) - 2; i >= 0; i-- {
			stmts[i] = Sequence(stmts[i], stmts[i+1])
		}
	}

	return stmts[0]
}

func (p *Parser) ParseStatement() Statement {
	switch next := p.tokens.Next(); {
	case next.Type == WHILE:
		exp := p.ParseExpression()

		block := p.ParseBlock()

		return While(exp, block)
	case next.Type == IF:
		exp := p.ParseExpression()

		thenBlock := p.ParseBlock()

		if p.tokens.Next().Type != ELSE {
			p.tokens.Rewind()
			panic(fmt.Errorf("expected literal else, got %q", p.tokens.Peek()))
		}

		elseBlock := p.ParseBlock()

		return Ite(exp, thenBlock, elseBlock)
	case next.Type == PRINT:
		return Print(p.ParseExpression())
	case next.Type == IDENTIFIER:
		switch operator := p.tokens.Next(); {
		case operator.Type == ASSIGMENT:
			return Assignment(next.Value, p.ParseExpression())
		case operator.Type == DECLARATION:
			return Declaration(next.Value, p.ParseExpression())
		default:
			panic(fmt.Errorf("expected declaration or assignment, got %q", operator))
		}
	default:
		p.tokens.Rewind()
		panic(fmt.Errorf("expected a statement, got %q", p.tokens.Peek()))
	}
}

func (p *Parser) ParseExpression() Expression {
	return p.parseDisjunction()
}

type ExpressionParser func() Expression
type BinaryExpressionCollector func(left Expression, right Expression) Expression

func (p *Parser) parseBinaryOperation(operatorToken TokenType, higherPrecedenceParser ExpressionParser, collector BinaryExpressionCollector) Expression {
	sub := higherPrecedenceParser()

	terms := []Expression{sub}

	for p.tokens.Peek().Type == operatorToken {
		p.tokens.Next()
		sub := higherPrecedenceParser()

		terms = append(terms, sub)
	}

	if len(terms) > 1 {
		for i := len(terms) - 2; i >= 0; i-- {
			terms[i] = collector(terms[i], terms[i+1])
		}
	}

	return terms[0]
}

func (p *Parser) parseDisjunction() Expression {
	return p.parseBinaryOperation(OR, p.parseConjunction, func(left Expression, right Expression) Expression {
		return Or(left, right)
	})
}

func (p *Parser) parseConjunction() Expression {
	return p.parseBinaryOperation(AND, p.parseEquality, func(left Expression, right Expression) Expression {
		return And(left, right)
	})
}

func (p *Parser) parseEquality() Expression {
	return p.parseBinaryOperation(EQUAL, p.parseComparison, func(left Expression, right Expression) Expression {
		return Equal(left, right)
	})
}

func (p *Parser) parseComparison() Expression {
	return p.parseBinaryOperation(LESS, p.parseTerm, func(left Expression, right Expression) Expression {
		return Lesser(left, right)
	})
}

func (p *Parser) parseTerm() Expression {
	return p.parseBinaryOperation(ADD, p.parseFactor, func(left Expression, right Expression) Expression {
		return Plus(left, right)
	})
}

func (p *Parser) parseFactor() Expression {
	return p.parseBinaryOperation(MUL, p.parseUnary, func(left Expression, right Expression) Expression {
		return Mult(left, right)
	})
}

func (p *Parser) parseUnary() Expression {
	if p.tokens.Peek().Type == NOT {
		p.tokens.Next()
		unary := p.parseUnary()

		return Negation(unary)
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Expression {
	switch next := p.tokens.Next(); {
	case next.Type == INT:
		intVar, err := strconv.Atoi(next.Value)
		if err != nil {
			panic(err)
		}

		return NumberExpression(intVar)
	case next.Type == BOOL:
		if next.Value == "true" {
			return Bool(true)
		} else if next.Value == "false" {
			return Bool(false)
		} else {
			panic(fmt.Errorf("invalid bool literal: %q", next.Value))
		}
	case next.Type == OPEN:
		exp := p.ParseExpression()

		after := p.tokens.Next()

		if after.Type != CLOSE {
			panic(fmt.Errorf("expected closing ), got: %q", after.Value))
		} else {
			return Grouping(exp)
		}
	case next.Type == IDENTIFIER:
		return Variable(next.Value)
	default:
		p.tokens.Rewind()
	}

	panic(fmt.Errorf("expected an primary expression, got %q", p.tokens.Peek()))
}

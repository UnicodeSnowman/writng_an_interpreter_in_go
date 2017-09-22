package parser

import (
	"fmt"

	"github.com/unicodesnowman/writing_an_interpreter_in_go/ast"
	"github.com/unicodesnowman/writing_an_interpreter_in_go/lexer"
	"github.com/unicodesnowman/writing_an_interpreter_in_go/token"
)

type Parser struct {
	l *lexer.Lexer

	errors       []string
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so currentToken and peekToken are both set
	//	p.nextToken()
	//	p.nextToken()
	// NOTE the above feels kind janky... what about...
	p.currentToken = p.l.NextToken()
	p.peekToken = p.l.NextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Advance
	p.nextToken()

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Advance
	p.nextToken()

	// TODO handle expressions, skipping until
	// semicolon for now
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	// NOTE Book version has this mutating p (by calling nextToken)
	// but I don't like this happening in an expect*** function
	if p.peekToken.Type != t {
		p.addPeekError(t)
		return false
	}

	return true
}

func (p *Parser) addPeekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currentToken.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

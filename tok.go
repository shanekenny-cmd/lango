package main

var TokenCats = [...]TokenCat{"IDENT", "INTLIT", "STRLIT", "MULOP", "DIVOP", "ADDOP", "SUBOP", "EQOP", "ASNOP", "LPAREN", "RPAREN", "LCBRAC", "RCBRAC", "KYWORD", "EOF"}

const (
	IDENT = iota
	INTLIT
	STRLIT
	MULOP
	DIVOP
	ADDOP
	SUBOP
	EQOP
	ASNOP
	LPAREN
	RPAREN
	LCBRAC
	RCBRAC
	KYWORD
	EOF
)

type TokenCat string

type TokenValue struct {
	numval int
	lexeme string
}

type Token struct {
	tocat TokenCat
	value TokenValue
}

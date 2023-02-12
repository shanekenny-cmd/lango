package main

import (
	"fmt"
	"strconv"
)

func IsAlpha(c rune) bool {
	return (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'A')
}

func IsAlphaNum(c rune) bool {
	return IsAlpha(c) || (c <= '9' && c >= '0')
}

func IsWhiteSpace(c rune) bool {
	return c == rune(' ') || c == rune('	')
}

func IsDigit(c rune) bool {
	return (c <= '9' && c >= '0')
}

func ToDigit(c rune) int {
	return int(c - '0')
}

var CurTok *Token
var index int
var text string
var lexSuccess bool

func Advance() {
	lexSuccess = true
	if index >= len(text) {
		var tok = new(Token)
		tok.value = TokenValue{0, "EOF"}
		tok.tocat = TokenCats[EOF]
		CurTok = tok
		return
	}
	var c rune = rune(text[index])
	var tok = new(Token)
	for IsWhiteSpace(c) && index < len(text) {
		index++
		if index >= len(text) {
			break
		}
		c = rune(text[index])
	}

	if IsAlpha(c) {
		var lexeme string
		for IsAlphaNum(c) && index < len(text) {
			lexeme = lexeme + string(c)
			index++
			if index < len(text) {
				c = rune(text[index])
			}
		}
		tok.value = TokenValue{0, lexeme}
		tok.tocat = TokenCat(TokenCats[IDENT])
		index--
	} else if index < len(text) && IsDigit(c) {
		var intLit int = 0
		for IsDigit(c) && index < len(text) {
			intLit = intLit * 10
			intLit = intLit + ToDigit(c)
			index++
			if index < len(text) {
				c = rune(text[index])
			}
		}

		tok.value = TokenValue{intLit, strconv.Itoa(intLit)}
		tok.tocat = TokenCat(TokenCats[INTLIT])
		index--
	} else if index >= len(text) {
		tok.value = TokenValue{0, "EOF"}
		tok.tocat = TokenCats[EOF]
		index = 0
	} else {
		tok.value = TokenValue{int(c), string(c)}
		switch c {
		case '(':
			tok.tocat = TokenCat(TokenCats[LPAREN])
		case ')':
			tok.tocat = TokenCat(TokenCats[RPAREN])
		case '+':
			tok.tocat = TokenCat(TokenCats[ADDOP])
		case '-':
			tok.tocat = TokenCat(TokenCats[SUBOP])
		case '{':
			tok.tocat = TokenCat(TokenCats[LCBRAC])
		case '}':
			tok.tocat = TokenCat(TokenCats[RCBRAC])
		case '*':
			tok.tocat = TokenCat(TokenCats[MULOP])
		case '/':
			tok.tocat = TokenCat(TokenCats[DIVOP])
		case '=':
			if index+1 < len(text) && text[index+1] == '=' {
				tok.tocat = TokenCat(TokenCats[EQOP])
                tok.value = TokenValue{int(c) * 2, "=="}
				index++
			} else {
				tok.tocat = TokenCat(TokenCats[ASNOP])
			}
		default:
			fmt.Printf("Unrecognized symbol: \"%c\"\n", c)
			lexSuccess = false
		}
	}
	if lexSuccess {
		CurTok = tok
	}
	index++
}

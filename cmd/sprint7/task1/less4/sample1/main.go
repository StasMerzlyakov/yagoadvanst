package main

import (
	"fmt"
	"sort"
	"unicode"
)

type TokenType int

const (
	TNumber = iota
	TIdent
)

const (
	StateMain = iota
	StateNumber
	StateIdent
)

// Token — информация о токене.
type Token struct {
	Type  TokenType
	Value string
}

// State — интерфейс состояния. Next передаёт очередной символ.
// Если символ разобран, то возвращается true.
type State interface {
	Next(rune) bool
}

// Number определяет число.
type Number struct {
	buf   []rune
	lexer *Lexer
}

func (l *Number) Next(r rune) bool {
	if unicode.IsDigit(r) {
		l.buf = append(l.buf, r)
		return true
	}
	l.lexer.NewToken(TNumber, l.buf)
	l.lexer.SetState(StateMain)
	l.buf = l.buf[:0]
	return false
}

type Ident struct {
	buf   []rune
	lexer *Lexer
}

func (id *Ident) Next(r rune) bool {
	switch {
	case len(id.buf) == 0 && unicode.IsLetter(r):
		{
			id.buf = append(id.buf, r)
			return true
		}
	case len(id.buf) > 0 && (unicode.IsLetter(r) || unicode.IsDigit(r)):
		{
			id.buf = append(id.buf, r)
			return true
		}

	}
	id.lexer.NewToken(TIdent, id.buf)
	id.lexer.SetState(StateMain)
	id.buf = id.buf[:0]
	return false
}

// Main — состояние по умолчанию.
type Main struct {
	lexer *Lexer
}

func (l *Main) Next(r rune) bool {
	if unicode.IsDigit(r) {
		l.lexer.SetState(StateNumber)
		return false
	}
	if unicode.IsLetter(r) {
		l.lexer.SetState(StateIdent)
		return false
	}
	return true
}

// Lexer содержит список состояний и полученные токены.
type Lexer struct {
	states []State
	state  State
	tokens []Token
}

// SetState изменяет состояние.
func (lex *Lexer) SetState(state int) {
	if state >= len(lex.states) {
		panic("unknown state")
	}
	lex.state = lex.states[state]
}

// NewToken добавляет токен.
func (lex *Lexer) NewToken(t TokenType, value []rune) {
	lex.tokens = append(lex.tokens, Token{
		Type:  t,
		Value: string(value),
	})
}

func main() {

	sort.StringSlice

	var lex Lexer

	// определяем состояния
	lex.states = []State{&Main{lexer: &lex}, &Number{lexer: &lex}, &Ident{lexer: &lex}}
	lex.SetState(StateMain)

	// пробуем разобрать эту строку
	s := "line778, 5 + 35 равно 40"
	for _, ch := range s {
		for !lex.state.Next(ch) {
		}
	}
	// завершаем разбор последнего токена, если он начат
	lex.state.Next(0)

	fmt.Println(lex.tokens)
	expected := "[{1 line778} {0 5} {0 35} {1 равно} {0 40}]"
	actual := fmt.Sprintf("%v", lex.tokens)
	fmt.Printf("equals: %v", expected == actual)
}

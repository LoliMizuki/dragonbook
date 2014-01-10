package main

type Lexer struct {
	line  int
	peek  int
	words map[string]Word
}

func NewLexer() Lexer {
	l := Lexer{1, ' ', nil}
	l.reserve(NewWordWithLexme(kTag_TRUE, "true"))
	l.reserve(NewWordWithLexme(kTag_FALSE, "false"))

	return l
}

// TODO:
// func (l *Lexer) Scan() Token {

// }

func (l *Lexer) reserve(w Word) {
	if l.words == nil {
		l.words = make(map[string]Word)
	}
	l.words[w.lexme] = w
}

package main

// 因為 ASCII 為 0~255, 所以這邊用 > 255 的數字來定義
const (
	kTag_NUM   = 256
	kTag_ID    = 257
	kTag_TRUE  = 258
	kTag_FALSE = 259
)

type Tag int

type Token struct {
	tag Tag
}

type Num struct {
	Token
	value int
}

// 保留字或識別字
type Word struct {
	Token
	lexme string
}

func NewTokenWithTag(tag Tag) Token {
	return Token{tag}
}

func NewNumWithValue(value int) Num {
	return Num{Token{kTag_NUM}, value}
}

func NewWordWithLexme(tag Tag, lexme string) Word {
	return Word{Token{tag}, lexme}
}

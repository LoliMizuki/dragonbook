/*
expr -> term rest

rest -> + term {print '+'} rest |
        - term {print '+'} rest |
        e
    // rest 套用了 tail recursive


term -> 0 { print '0' } |
        1 { print '1' } |
        ...
        9 { print '9' } |
*/

package main

import (
	"fmt"
	"os"
)

var sourceCode string
var currentIndex int

var lookhead int
var eof bool

func InitReaderWithData(srcCode string) {
	sourceCode = srcCode
	currentIndex = -1
	eof = false
}

func readNext() int {
	currentIndex++
	if currentIndex >= len(sourceCode) {
		eof = true
		return -1
	}

	return int(sourceCode[currentIndex])
}

func expr() {
	term()
	rest()
}

func rest() {
	// rest, use tail recursive -> expand to loop
	for {
		if lookhead == '+' {
			match('+')
			term()
			fmt.Println("+")

			if eof {
				os.Exit(0)
			}
		} else if lookhead == '-' {
			match('-')
			term()
			fmt.Println("-")

			if eof {
				os.Exit(0)
			}
		} else {
			return
		}
	}
}

func term() {
	if isDigit(lookhead) {
		fmt.Println(string(lookhead))
		match(lookhead)
	} else {
		panic("syntax error")
	}
}

func match(t int) {
	if lookhead == t {
		lookhead = readNext()
	} else {
		panic("syntax error")
	}
}

func isDigit(t int) bool {
	return ('0' <= t && t <= '9')
}

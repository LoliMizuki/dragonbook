package main

import ()

func main() {
	InitReaderWithData("7+9-4")

	lookhead = readNext()
	expr()
}

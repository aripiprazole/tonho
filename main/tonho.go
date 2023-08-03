package main

import (
	"fmt"
	"tonho"
)

func main() {
	tokens := tonho.Lex("test", "fun main() { println(\"hello world\") }")
	fmt.Printf("%v", tokens)
}

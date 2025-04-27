package parser

import (
	"fmt"

	"github.com/4am-png/Nexa-Next-Generation-Systems-Architecture/lexer"
)

func Parse(tokens []lexer.Token) {
	for _, token := range tokens {
		fmt.Println("Parsed Token:", token.Type, "with value", token.Literal)
	}
}

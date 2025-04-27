package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Використання: nexa <файл.nexa>")
		return
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Помилка при зчитуванні файлу: %v\n", err)
		return
	}

	tokens := Lex(string(content))
	ast := Parse(tokens)
	output := Translate(ast)

	fmt.Println("Згенерований код:")
	fmt.Println(output)
}

package main

// Node представляє елемент дерева AST
type Node struct {
	Type  string
	Value string
}

// Parse приймає токени і повертає AST
func Parse(tokens []Token) []Node {
	var nodes []Node

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Type {
		case "PRINT":
			if i+1 < len(tokens) {
				nodes = append(nodes, Node{Type: "PRINT", Value: tokens[i+1].Literal})
				i++
			}
		}
	}

	return nodes
}

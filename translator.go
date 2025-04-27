package main

// Translate приймає AST і генерує Bash код
func Translate(nodes []Node) string {
	var output string

	for _, node := range nodes {
		switch node.Type {
		case "PRINT":
			output += "echo \"" + node.Value + "\"\n"
		}
	}

	return output
}

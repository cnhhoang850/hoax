package metaAST

import (
	"fmt"
	"os"
	"strings"
)

// metaAST generates Go code for creating an AST (Abstract Syntax Tree) based on given types.
func metaAST() {
	// Check for command-line arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <output_path> <base_name> <type1> <type2> ...")
		return
	}

	// Extract command-line arguments
	outputPath := os.Args[1]
	baseName := os.Args[2]
	types := os.Args[3:]

	defineAST(outputPath, baseName, types)
}

// defineAST generates code for the base class and its subclasses based on provided types.
func defineAST(path string, baseName string, types []string) {
	// Open the file for writing, creating if it doesn't exist
	file, err := os.Create(path + "/" + baseName + ".go")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	basenameContent := fmt.Sprintf(
		"package parser\n\n type %s struct {\n}\n\n",
		baseName,
	)

	// Iterate through provided types and generate code for each subclass
	for _, typeStr := range types {
		parts := strings.Split(typeStr, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid type string", typeStr)
			return
		}

		className := strings.TrimSpace(parts[0])
		fields := strings.TrimSpace(parts[1])

		defineSubClass(baseName, className, fields)
	}

	_, err = fmt.Fprintf(file, basenameContent, baseName)
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}
}

// defineSubClass generates code for a subclass with the specified fields.
func defineSubClass(structName string, className string, fields string) string {
	fieldSplits := strings.Split(fields, ",")

	// Add fields into struct
	structString := fmt.Sprintf(
		"type %s struct {\n%s\n}\n\n",
		structName,
		strings.Join(fieldSplits, "\n"),
	)

	return structString
}

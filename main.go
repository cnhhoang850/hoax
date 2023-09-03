package main

import (
	"bufio"
	"fmt"
	"hoax/abstractSyntaxTree"
	"hoax/scanner"
	"hoax/token"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	switch {
	case length == 1:
		fmt.Println("Run file")
	case length > 1:
		fmt.Println("Run file with args")
	case length == 0:
		fmt.Println("Run repl")
	}

	scanner := scanner.Scanner{
		Source:  " thisIsAnIdentifier and 'this is a string' ( ) ! != 1",
		Start:   0,
		Pointer: 0,
		Line:    1,
		Tokens:  []token.Token{},
	}

	scanner.ScanTokens()
	fmt.Println(scanner.Tokens)

	//Test generateAST
	abstractSyntaxTree.GenerateAST("./parser", "Expr", []string{
		"Expression : Expression Expr",
		"Binary   : Left Expr, Operator token.Token, Right Expr",
		"Grouping : Expression Expr",
		"Literal  : Value token.Token", // interface{} is a placeholder for any type
		"Unary    : Operator token.Token, Right Expr",
	})
}

func runFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(65)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Read into byte slices
	byteSlice := make([]byte, 0)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer) // read into buffer
		fmt.Println(n)
		if err != nil && err != bufio.ErrBufferFull && err != io.EOF {
			fmt.Println("Error reading file")
			os.Exit(65)
		}

		byteSlice = append(byteSlice, buffer[:n]...)
		fmt.Println(byteSlice)

		if err == io.EOF {
			break
		}
	}

	run(string(byteSlice))
}

func runPrompt() {
	inputStreamReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")

		input, err := inputStreamReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			os.Exit(65)
		}

		input = strings.TrimSpace(input)
		run(input)

		if input == "exit" {
			break
		}
	}
}

func run(data string) {

}

func lex(rawData string) {
	fmt.Println("lexing", rawData)
	// turn raw data in to lexemes
}

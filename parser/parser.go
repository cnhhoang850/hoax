package parser

import (
	"fmt"
	"hoax/token"
)

type Expression struct {
	Left     *Expression
	Right    *Expression
	Operator token.Token
}

type Binary struct {
	Expression
}

type Unary struct {
	Expression
}

type Literal struct {
	Expression
}

type Parser struct {
	Tokens []token.Token
}

// invoke a Parser instance with a list of tokens

// loop through the list one by one

// create a new Production instance on an equality

// syntax trees holds the intermmediate representation of the program between the list of tokens vs being parsed by the interpreter

// => they should just be data

func (p *Parser) Parse() {
	fmt.Println("Parsing tokens")
	for _, t := range p.Tokens {
		if t.Type == token.EQUAL {
			p.Evaluate(t)
		}
	}
}

func (p *Parser) Evaluate(t token.Token) {

}

// NOTES: In fact, these types exist to enable the parser and interpreter to communicate. That lends itself to types that are simply data with no associated behavior. This style is very natural in functional languages like Lisp and ML where all data is separate from behavior, but it feels odd in Java.

//Functional programming aficionados right now are jumping up to exclaim “See! Object-oriented languages are a bad fit for an interpreter!” I won’t go that far. You’ll recall that the scanner itself was admirably suited to object-orientation. It had all of the mutable state to keep track of where it was in the source code, a well-defined set of public methods, and a handful of private helpers.

//My feeling is that each phase or part of the interpreter works fine in an object-oriented style. It is the data structures that flow between them that are stripped of behavior.

package parser

import (
	"fmt"
	"hoax/token"
)

// We declare interface to wrap function signatures around the concrete structs
type ExpressionInterface interface {
	Accept(visitor VisitorInterface)
}

type Expression struct{}
type Binary struct {
	Expression
	Left     *Expression
	Right    *Expression
	Operator token.Token
}

func (b *Binary) Accept(visitor VisitorInterface) {
	visitor.VisitBinary(b)
}

type Unary struct {
	Expression
	Operator token.Token
	Right    *Expression
}

func (u *Unary) Accept(visitor VisitorInterface) {
	visitor.VisitUnary(u)
}

type Literal struct {
	Expression
	Value interface{}
}

func (l *Literal) Accept(visitor VisitorInterface) {
	visitor.VisitLiteral(l)
}

type Grouping struct {
	Left  token.Token
	Body  *Expression
	Right token.Token
}

func (g *Grouping) Accept(visitor VisitorInterface) {
	visitor.VisitGrouping(g)
}

type VisitorInterface interface {
	VisitBinary(expr *Binary)
	VisitUnary(expr *Unary)
	VisitLiteral(expr *Literal)
	VisitGrouping(expr *Grouping)
}

// Visitors implementations
type Visitor struct{}

func (v *Visitor) VisitBinary(expr *Binary) {
	fmt.Println("Visiting Binary")
}
func (v *Visitor) VisitUnary(expr *Unary) {
	fmt.Println("Visiting Unary")
}
func (v *Visitor) VisitLiteral(expr *Literal) {
	fmt.Println("Visiting Literal")
}
func (v *Visitor) VisitGrouping(expr *Grouping) {
	fmt.Println("Visiting Grouping")
}

// NOTES: In fact, these types exist to enable the parser and interpreter to communicate. That lends itself to types that are simply data with no associated behavior. This style is very natural in functional languages like Lisp and ML where all data is separate from behavior, but it feels odd in Java.

//Functional programming aficionados right now are jumping up to exclaim “See! Object-oriented languages are a bad fit for an interpreter!” I won’t go that far. You’ll recall that the scanner itself was admirably suited to object-orientation. It had all of the mutable state to keep track of where it was in the source code, a well-defined set of public methods, and a handful of private helpers.

//My feeling is that each phase or part of the interpreter works fine in an object-oriented style. It is the data structures that flow between them that are stripped of behavior.

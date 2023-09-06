package abstractSyntaxTree

import (
	"fmt"
	"hoax/parser"
	"hoax/token"
	"strings"
)

// print AST tree in LISP style (* (- 12) (/ 2 3)) = -12 * 2 / 3

type PrettyPrinter struct {
	depth int
}

func (p *PrettyPrinter) VisitExpression(expr *parser.Expression) {
	expr.Expression.Accept(p)
}

func (p *PrettyPrinter) VisitBinary(expr *parser.Binary) {
	p.depth++
	fmt.Println()
	fmt.Print(strings.Repeat("\t", p.depth))
	fmt.Print("(")
	fmt.Print(expr.Operator.Lexeme + " ")
	expr.Left.Accept(p)
	fmt.Print(" ")
	expr.Right.Accept(p)
	fmt.Print(")")
}

func (p *PrettyPrinter) VisitGrouping(expr *parser.Grouping) {
	fmt.Print("(")
	expr.Expression.Accept(p)
	fmt.Print(")")
}

func (p *PrettyPrinter) VisitLiteral(expr *parser.Literal) {
	fmt.Print(expr.Value.Literal)
}

func (p *PrettyPrinter) VisitUnary(expr *parser.Unary) {
	fmt.Print("(")
	fmt.Print(expr.Operator.Lexeme + " ")
	expr.Right.Accept(p)
	fmt.Print(")")
}

func Printer() {
	expression := parser.Expression{
		Expression: &parser.Binary{
			Left: &parser.Literal{
				Value: token.NewToken(token.NUMBER, "12", 12, 1),
			},
			Operator: token.NewToken(token.MINUS, "-", nil, 1),
			Right: &parser.Expression{
				Expression: &parser.Binary{
					Left: &parser.Literal{
						Value: token.NewToken(token.NUMBER, "12", 12, 1),
					},
					Operator: token.NewToken(token.MINUS, "-", nil, 1),
					Right: &parser.Literal{
						Value: token.NewToken(token.NUMBER, "12", 12, 1),
					},
				},
			},
		},
	}

	PrettyPrint(&expression)
}

func PrettyPrint(expression parser.Expr) {
	printer := &PrettyPrinter{}
	expression.Accept(printer)
}

// Don't understand whether to use value receiver or pointer receiver
// &parser.Binary instantiate AND give a pointer with &
// (+ 1 )
//
//

package abstractSyntaxTree

import (
	"fmt"
)

// print AST tree in LISP style (* (- 12) (/ 2 3)) = -12 * 2 / 3

type PrettyPrinter struct{}

func (v *PrettyPrinter) VisitBinary(expr *parser.Binary) {
	fmt.Println("Visiting Binary")
}

func (v *PrettyPrinter) VisitGrouping(expr *parser.Grouping) {
	fmt.Println("Visiting Grouping")
}

func (v *PrettyPrinter) VisitLiteral(expr *parser.Literal) {
	fmt.Println("Visiting Literal")
}

func (v *PrettyPrinter) VisitUnary(expr *parser.Unary) {
	fmt.Println("Visiting Unary")
}

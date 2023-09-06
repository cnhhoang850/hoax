package parser

import (
	"fmt"
	"hoax/token")
 type Expr interface {
	Accept(visitor VisitorInterface)
}

type Expression struct {
	Expr Expr
	Expression Expr
}

func (x *Expression) Accept(visitor VisitorInterface) {
	visitor.VisitExpression(x)
}
type Binary struct {
	Expr Expr
	Left Expr
	 Operator token.Token
	 Right Expr
}

func (x *Binary) Accept(visitor VisitorInterface) {
	visitor.VisitBinary(x)
}
type Grouping struct {
	Expr Expr
	Expression Expr
}

func (x *Grouping) Accept(visitor VisitorInterface) {
	visitor.VisitGrouping(x)
}
type Literal struct {
	Expr Expr
	Value token.Token
}

func (x *Literal) Accept(visitor VisitorInterface) {
	visitor.VisitLiteral(x)
}
type Unary struct {
	Expr Expr
	Operator token.Token
	 Right Expr
}

func (x *Unary) Accept(visitor VisitorInterface) {
	visitor.VisitUnary(x)
}

type VisitorInterface interface {
	VisitExpression(expr *Expression)
	VisitBinary(expr *Binary)
	VisitGrouping(expr *Grouping)
	VisitLiteral(expr *Literal)
	VisitUnary(expr *Unary)
}

type Visitor struct{}


func (v *Visitor) VisitExpression(expr *Expression) {
	fmt.Println("Visiting Expression")
}

func (v *Visitor) VisitBinary(expr *Binary) {
	fmt.Println("Visiting Binary")
}

func (v *Visitor) VisitGrouping(expr *Grouping) {
	fmt.Println("Visiting Grouping")
}

func (v *Visitor) VisitLiteral(expr *Literal) {
	fmt.Println("Visiting Literal")
}

func (v *Visitor) VisitUnary(expr *Unary) {
	fmt.Println("Visiting Unary")
}

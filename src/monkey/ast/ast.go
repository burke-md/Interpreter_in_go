package ast

// Base Node interface
type Node interface {
	TokenLiteral() string
	String() string
}

// All statement nodes require this interface
type Statement interface {
	Node
	statementNode()
}

// All expression nodes require this interface
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statment
}

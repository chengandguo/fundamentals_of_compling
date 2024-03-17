package ast

import (
	"monkey/token"
)

type Node struct {
	Type string
	Kind string
}

type Ast struct {
	tokenList []token.Token
}

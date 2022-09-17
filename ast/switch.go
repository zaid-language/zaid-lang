package ast

import (
	"zaidlang.tech/x/zaid/token"
)

type Switch struct {
	ExpressionNode
	Token token.Token    // The "switch" token
	Value ExpressionNode // The value that will be used to determine the case
	Cases []*Case        // The cases this switch statement will handle
}

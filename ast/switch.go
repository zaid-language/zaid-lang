package ast

import (
	"github.com/zaid-language/zaid-lang/token"
)

type Switch struct {
	ExpressionNode
	Token token.Token    // The "switch" token
	Value ExpressionNode // The value that will be used to determine the case
	Cases []*Case        // The cases this switch statement will handle
}

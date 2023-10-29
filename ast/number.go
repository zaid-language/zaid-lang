package ast

import (
	"zaidlang.org/x/zaid/token"

	"github.com/shopspring/decimal"
)

type Number struct {
	ExpressionNode
	Token token.Token
	Value decimal.Decimal
}

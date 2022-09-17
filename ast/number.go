package ast

import (
	"github.com/zaid-language/zaid/token"

	"github.com/shopspring/decimal"
)

type Number struct {
	ExpressionNode
	Token token.Token
	Value decimal.Decimal
}

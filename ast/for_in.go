package ast

import "github.com/zaid-language/zaid-lang/token"

type ForIn struct {
	ExpressionNode
	Token    token.Token    // for
	Key      *Identifier    // key
	Value    *Identifier    // value
	Iterable ExpressionNode // list, map
	Block    *Block         // { ... }
}

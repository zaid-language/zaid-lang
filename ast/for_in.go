package ast

import "zaidlang.tech/x/zaid/token"

type ForIn struct {
	ExpressionNode
	Token    token.Token    // for
	Key      *Identifier    // key
	Value    *Identifier    // value
	Iterable ExpressionNode // list, map
	Block    *Block         // { ... }
}

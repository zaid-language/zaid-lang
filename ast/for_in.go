package ast

import "zaidlang.org/x/zaid/token"

type ForIn struct {
	ExpressionNode
	Token    token.Token    // for
	Key      *Identifier    // key
	Value    *Identifier    // value
	Iterable ExpressionNode // list, map
	Block    *Block         // { ... }
}

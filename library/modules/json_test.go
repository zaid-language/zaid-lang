package modules

import (
	"testing"

	"github.com/shopspring/decimal"
	"zaidlang.org/x/zaid/object"
	"zaidlang.org/x/zaid/token"
)

func TestJsonDecode(t *testing.T) {
	input := `{"name": "John", "age": 30, "city": "New York"}`

	expected := &object.Map{Pairs: map[object.MapKey]object.MapPair{
		(&object.String{Value: "name"}).MapKey(): {Key: &object.String{Value: "name"}, Value: &object.String{Value: "Zaid"}},
		(&object.String{Value: "age"}).MapKey():  {Key: &object.String{Value: "age"}, Value: &object.Number{Value: decimal.NewFromInt(21)}},
	}}

	result := jsonDecode(nil, token.Token{}, &object.String{Value: input})

	if result.Type() != expected.Type() {
		t.Errorf("wrong result type. got=%s, expected=%s", result.Type(), expected.Type())
	}
}

func TestJsonEncode(t *testing.T) {
	input := &object.Map{Pairs: map[object.MapKey]object.MapPair{
		(&object.String{Value: "name"}).MapKey(): {Key: &object.String{Value: "name"}, Value: &object.String{Value: "Zaid"}},
		(&object.String{Value: "age"}).MapKey():  {Key: &object.String{Value: "age"}, Value: &object.Number{Value: decimal.NewFromInt(21)}},
	}}

	expected := `{"age":21,"name":"Zaid"}`

	result := jsonEncode(nil, token.Token{}, input)

	if result.String() != expected {
		t.Errorf("wrong result. got=%s, expected=%s", result.String(), expected)
	}
}

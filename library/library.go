package library

import (
	"zaidlang.org/x/zaid/library/functions"
	"zaidlang.org/x/zaid/library/modules"
	"zaidlang.org/x/zaid/object"
)

var Functions = map[string]*object.LibraryFunction{}
var Modules = map[string]*object.LibraryModule{}

func init() {
	RegisterModule("console", modules.ConsoleMethods, modules.ConsoleProperties)
	RegisterModule("zaid", modules.ZaidMethods, modules.ZaidProperties)
	RegisterModule("http", modules.HttpMethods, modules.HttpProperties)
	RegisterModule("io", modules.IoMethods, modules.IoProperties)
	RegisterModule("math", modules.MathMethods, modules.MathProperties)
	RegisterModule("os", modules.OsMethods, modules.OsProperties)
	RegisterModule("random", modules.RandomMethods, modules.RandomProperties)
	RegisterModule("time", modules.TimeMethods, modules.TimeProperties)
	RegisterModule("json", modules.JsonMethods, modules.JsonProperties)

	RegisterFunction("printftw", functions.Print)
	RegisterFunction("type", functions.Type)
}

func RegisterFunction(name string, function object.GoFunction) {
	Functions[name] = &object.LibraryFunction{Name: name, Function: function}
}

func RegisterModule(name string, methods map[string]*object.LibraryFunction, properties map[string]*object.LibraryProperty) {
	Modules[name] = &object.LibraryModule{Name: name, Methods: methods, Properties: properties}
}

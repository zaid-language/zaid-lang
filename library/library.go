package library

import (
	"github.com/zaid-language/zaid/library/functions"
	"github.com/zaid-language/zaid/library/modules"
	"github.com/zaid-language/zaid/object"
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

	RegisterFunction("printftw", functions.Print)
	RegisterFunction("type", functions.Type)
}

func RegisterFunction(name string, function object.GoFunction) {
	Functions[name] = &object.LibraryFunction{Name: name, Function: function}
}

func RegisterModule(name string, methods map[string]*object.LibraryFunction, properties map[string]*object.LibraryProperty) {
	Modules[name] = &object.LibraryModule{Name: name, Methods: methods, Properties: properties}
}

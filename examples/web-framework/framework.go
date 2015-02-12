package main

import (
	"go-callbacks/examples/web-framework/core"

	// register modules, is required since go is static
	"go-callbacks/examples/web-framework/modules/contact"
	"go-callbacks/examples/web-framework/modules/session"
)

func main() {
	var context core.AppContext
	registerCallbacks(&context)
	context.Init()
	context.Run()
}

// registerCallbacks calls the .RegisterCallbacks method for each desired module
func registerCallbacks(context *core.AppContext) {
	// add callbacks method for each module, naimg does not matter
	session.RegisterCallbacks(context)
	contact.RegisterCallbacks(context)

	// fire hook, modules might be interested to check if other modules exists
	context.CallbacksCall("init.callbacks_loaded", context)
}

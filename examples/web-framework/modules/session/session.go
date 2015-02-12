package session

import (
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/schorsch/go-callbacks"
	"go-callbacks/examples/web-framework/core"
)

// register the hook into init.middleware
func RegisterCallbacks(context *core.AppContext) {
	fmt.Println("register session module")
	// define a single callback
	cb := callbacks.Callback{"init.middleware", Init}
	context.Callbacks = append(context.Callbacks, cb)

}

// Init session module with:
// - session middleware provided by negroni-sessions (based on gorilla/session)
// - simple cookie store
func Init(context *core.AppContext) {

	fmt.Println("Init session module")
	// TODO config secret & session key
	session_store := cookiestore.New([]byte("secret123"))
	context.Middleware.Use(sessions.Sessions("my_session", session_store))
}

package core

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
	"github.com/schorsch/go-callbacks"
)

//AppContext is created on boot and keeps app wide configurations.
//On startup each module receives the app context and can use or edit the config values
type AppContext struct {
	Router     *pat.Router
	Middleware *negroni.Negroni
	callbacks.Callbacks
	// Db, Es, Memcache ..other App-Global configs
}

//Run starts the web server via negroni
func (c *AppContext) Run() {
	c.Middleware.Run(":3000")
}

//Init runs the parts required for the core
func (c *AppContext) Init() {
	c.initRouter()
	c.initModules()
	c.initMiddleware()
}

func (c *AppContext) initModules() {
	c.CallbacksCall("init.modules", c)
}

func (c *AppContext) initRouter() {
	c.Router = pat.New()
}

func (c *AppContext) initMiddleware() {
	// init negroni with default middlewares
	c.Middleware = negroni.Classic()
	// let others register middleware
	c.CallbacksCall("init.middleware", c)
	// important add self.routes as LAST middleware, c.Router must be init'ed
	c.Middleware.UseHandler(c.Router)
	fmt.Printf("%v middlewares registered\n", len(c.Middleware.Handlers()))

}

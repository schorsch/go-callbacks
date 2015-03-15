package callbacks

import (
	"errors"
	"log"
	"reflect"
)

//Callback used by modules to register hooks with a name and a function
type Callback struct {
	Name   string
	Method interface{}
}

//Call the method of a callback. Detects if the given params are valid
func (c *Callback) Call(params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(c.Method)
	if len(params) != f.Type().NumIn() {
		// TODO add method name and param numbers
		log.Println("ERROR: callback.Call parameter error.")
		err = errors.New("Callback Call: wrong number of params.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

//Callbacks are a simple list of callbacks. The type is meant to be used in your own structs by composition.
// Afterwards just add callbacks to the list and than later on call them
//
//		type MyAppContext struct {
//			Router     *pat.Router
//			Middleware *negroni.Negroni
//			Callbacks
//		}
// 		// init context + add callbacks
//		// run all named callbacks
// 		an_app_context.CallbacksCall("init.middlewares")
//
type Callbacks []Callback

//CallbacksCall calls all callbacks in the list by the given name with the passed params
func (c *Callbacks) CallbacksCall(name string, params ...interface{}) {
	log.Println("callbacks call:", name)
	for _, cb := range c.CallbacksFind(name) {
		// TODO catch errors ary ?
		// TODO return values ary if any?
		cb.Call(params...)
	}
}

//CallbacksFind all callbacks in the list by name.
// Returns a new instance of Callbacks list
func (c *Callbacks) CallbacksFind(name string) (cbs Callbacks) {
	for _, cb := range *c {
		if cb.Name == name {
			cbs = append(cbs, cb)
		}
	}
	return cbs
}

package callbacks_test

import (
//	"fmt"
//	"strings"
	"testing"

	"github.com/schorsch/go-callbacks"
)

// setup
type StructWithCallbacks struct {
	callbacks.Callbacks
	CallResult string
}

func (c *StructWithCallbacks) SimpleMethod() {
	c.CallResult = "Hello"
}
func (c *StructWithCallbacks) MethodWithParam(res string) {
	c.CallResult = res
}

func TestCallbacksCall_OnStruct(t *testing.T) {
	var context StructWithCallbacks
	// add callback method 'SimpleMethod' run when 'a.hooks' is called
	cb := callbacks.Callback{Name: "a.hooks", Method: context.SimpleMethod}
	context.Callbacks = append(context.Callbacks, cb)
	// add 'MethodWithParam' executed when 'b.callbacks' is called
	cb1 := callbacks.Callback{"b.callbacks", context.MethodWithParam}
	context.Callbacks = append(context.Callbacks, cb1)

	context.CallbacksCall("a.hooks")
	if context.CallResult != "Hello" {
		t.Error("Expected Hello, got ", context.CallResult)
	}
	context.CallbacksCall("b.callbacks", "Hi")
	if context.CallResult != "Hi" {
		t.Error("Expected hi, got ", context.CallResult)
	}
}

func TestCallbacksCall_OnInlineMethod(t *testing.T) {
	counter := 0
	//	simple func incrementing the counter by given value
	inline_callback := func(incr int){
		counter += incr
	}
	var cbs callbacks.Callbacks
	cb := callbacks.Callback{Name: "a.run", Method: inline_callback}
	cbs = append(cbs, cb)

	cbs.CallbacksCall("a.run", 5)
	if counter != 5 {
		t.Error("Expected 5, got ", counter)
	}
}

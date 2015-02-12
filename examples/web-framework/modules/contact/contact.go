package contact

import (
	"fmt"
	"github.com/schorsch/go-callbacks"

	// the core context needed to register & init hooks
	"go-callbacks/examples/web-framework/core"

	// needed for controller
	"encoding/json"
	"net/http"

	// still need to get session
	"github.com/goincremental/negroni-sessions"
)

// ------------- Init

func RegisterCallbacks(context *core.AppContext) {
	fmt.Println("register contact module")
	// define this modules callbacks
	cbs := []callbacks.Callback{
		// hook into 2 core callbacks
		{Name: "init.modules", Method: Init},
		{Name: "init.callbacks_loaded", Method: TestCallback},
		// define own callbacks for other modules to use
		//{Name: "contact.delete", Method: TestCallback},
	}
	// push them onto the App Context
	context.Callbacks = append(context.Callbacks, cbs...)
}

// Init contact module with:
// - adding routes
func Init(context *core.AppContext) {
	fmt.Println("Init contact module")
	// add a route to the apps router, here we use pat
	context.Router.Get("/contacts", GetContacts)
}

func TestCallback(context *core.AppContext) {
	fmt.Println("Hello From the contact module")
}

// ------------- the model structure

type Contact struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func ContactsFindAll() []Contact {
	return []Contact{
		{"1", "peter@me.com"},
		{"2", "mary@me.com"},
	}
}

// ------------- Controller methods

func GetContacts(w http.ResponseWriter, r *http.Request) {
	// uses session module
	session := sessions.GetSession(r)
	session.Set("hello", "world")

	var result = ContactsFindAll()

	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(js)
}

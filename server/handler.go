package server

import (
	"encoding/json"
	"net/http"

	"github.com/billinghamj/go-petition-example/app"
)

type context struct {
	App     *app.App
	Request *http.Request
}

type handlerFn func(context, map[string]interface{}) (interface{}, error)

type handler struct {
	app *app.App
	fn  handlerFn
}

func createHandler(app *app.App, fn handlerFn) handler {
	return handler{app, fn}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context{h.app, r}
	var out interface{}
	var err error

	// todo: handle missing content (assume empty map)
	// todo: check Content-Type
	// todo: support urlencoded form
	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(400)
		return
	}

	out, err = h.fn(ctx, data)

	if err != nil {
		// todo: output meaningful error
		w.WriteHeader(500)
		return
	}

	if out == nil {
		w.WriteHeader(204)
		return
	}

	// todo: check Accept

	if data, err := json.Marshal(out); err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
		w.Write(data)
	}
}

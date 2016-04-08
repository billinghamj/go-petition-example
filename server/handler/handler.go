package handler

import (
	"net/http"

	"github.com/billinghamj/go-petition-example/app"
)

type handlerFn func(Context, map[string]interface{}) (interface{}, error)

// Handler allows a semantic function to be used as an HTTP handler
type Handler struct {
	app *app.App
	fn  handlerFn
}

// Create returns the function wrapped into an HTTP handler
func Create(app *app.App, fn handlerFn) Handler {
	return Handler{app, fn}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var outputType string
	var input map[string]interface{}
	var output interface{}
	var status int

	if outputType, err = getOutputType(r); err != nil {
		w.WriteHeader(406)
		return
	}

	if input, err = getInput(r); err != nil {
		w.WriteHeader(415)
		return
	}

	if result, err := h.fn(Context{h.app, r}, input); err != nil {
		output = err
		status = 500 // todo: other status codes
	} else {
		if result == nil {
			status = 204
		} else {
			output = result
			status = 200
		}
	}

	if output == nil {
		w.WriteHeader(status)
		return
	}

	if err = writeOutput(w, status, outputType, output); err != nil {
		w.WriteHeader(500)
		return
	}
}

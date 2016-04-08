package handler

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

func getInput(r *http.Request) (data map[string]interface{}, err error) {
	header := r.Header.Get("Content-Type")

	if header == "" {
		data = make(map[string]interface{})
		return
	}

	mediaType, _, err := mime.ParseMediaType(header)

	if err != nil {
		return
	}

	// todo: support application/x-www-form-urlencoded

	switch mediaType {
	default:
		err = fmt.Errorf("unsupported media type")

	case "application/json":
		data = make(map[string]interface{})
		err = json.NewDecoder(r.Body).Decode(&data)
	}

	return
}

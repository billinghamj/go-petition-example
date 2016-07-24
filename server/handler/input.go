package handler

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/billinghamj/go-petition-example/log"
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

	switch mediaType {
	default:
		err = log.CreateError("unsupported_media_type", nil)

	case "application/json":
		data = make(map[string]interface{})
		err = json.NewDecoder(r.Body).Decode(&data)

	case "application/x-www-form-urlencoded":
		// currently very primitive - no support for maps/arrays/slices/etc.
		// only able to output `map[string]string`
		// todo: support proper structured `a[0][abc]=1234` format

		if err = r.ParseForm(); err == nil {
			data = make(map[string]interface{})

			for k := range r.PostForm {
				data[k] = r.PostForm.Get(k)
			}
		}
	}

	return
}

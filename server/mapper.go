package server

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/mitchellh/mapstructure"
)

func mapInput(input map[string]interface{}, output interface{}) *log.Error {
	mapper, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           output,
		TagName:          "json",
		WeaklyTypedInput: true,
	})

	if err != nil {
		panic(err)
	}

	if err = mapper.Decode(input); err != nil {
		return log.CreateError("mapper_failed", map[string]interface{}{"error": err})
	}

	return nil
}

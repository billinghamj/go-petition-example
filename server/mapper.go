package server

import "github.com/mitchellh/mapstructure"

func mapInput(input map[string]interface{}, output interface{}) error {
	mapper, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           output,
		TagName:          "json",
		WeaklyTypedInput: true,
	})

	if err != nil {
		panic(err)
	}

	return mapper.Decode(input)
}

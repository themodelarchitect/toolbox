package tools

import "encoding/json"

func DecodeJSON(in []byte) (interface{}, error) {
	var parsed any
	err := json.Unmarshal(in, &parsed)
	return parsed, err
}

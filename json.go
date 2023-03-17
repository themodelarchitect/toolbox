package tools

import "encoding/json"

func DecodeJSON(in []byte) (any, error) {
	var parsed any
	err := json.Unmarshal(in, &parsed)
	return parsed, err
}

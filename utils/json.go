package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) string {
	str, err := json.MarshalIndent(&v, "", " ")

	if err != nil {
		panic(err)
	}

	return string(str)
}

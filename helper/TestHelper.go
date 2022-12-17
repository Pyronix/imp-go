package helper

import (
	"encoding/json"
)

func StructToJson(class interface{}) string {
	out, err := json.Marshal(class)
	if err != nil {
		panic(err)
	}
	return string(out)
}

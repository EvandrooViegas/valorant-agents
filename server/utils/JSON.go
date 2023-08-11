package utils

import (
	"encoding/json"
)

func JSONstringify(value interface{}) string {
	v, err := json.Marshal(value)
	HandleErr(err)
	return string(v)
}

func JSONparse[T any](value []byte, data *T)   {
	err := json.Unmarshal(value, &data)
	HandleErr(err)
}


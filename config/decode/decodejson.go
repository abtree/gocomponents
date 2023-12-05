package decode

import (
	"encoding/json"
	"log"
)

func decodeJson(dat string, v interface{}) {
	err := json.Unmarshal([]byte(dat), v)
	if err != nil {
		log.Panicf(err.Error())
	}
}

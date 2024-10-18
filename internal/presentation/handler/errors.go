package handler

import (
	"encoding/json"
	"log"
)

func BadRequest(flag string) []byte {
	errorMessage := map[string]string{
		"error": "bad request",
		"details": "you missing flag: " + flag,
	}
	
	res, err := json.Marshal(errorMessage)
	if err != nil {
		log.Println(err)
		return nil
	}
	return res
}

func InternalError(err error) []byte {
	errorMessage := map[string]string{
		"error": "internal server error",
		"details": err.Error(),
	}

	res, marshalErr := json.Marshal(errorMessage)
	if marshalErr != nil {
		log.Println(marshalErr)
		return nil
	}
	return res
}

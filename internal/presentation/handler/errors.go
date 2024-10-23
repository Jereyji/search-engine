package handler

import (
	"encoding/json"
	"log"
)

const (
	MissingFlag   = "missing required flag: "
	IncorrectData = "incorrect data: "
)

func BadRequest(errMessage, value string) []byte {
	errorMessage := map[string]string{
		"error":   "bad request",
		"details": errMessage + value,
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
		"error":   "internal server error",
		"details": err.Error(),
	}

	res, marshalErr := json.Marshal(errorMessage)
	if marshalErr != nil {
		log.Println(marshalErr)
		return nil
	}
	return res
}

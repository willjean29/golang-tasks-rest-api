package utils

import (
	"encoding/json"
	"errors"
	"io"
)

func TransformBody(body io.Reader, params ...interface{}) error {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	for _, param := range params {
		err = selectStructType(param, reqBody)
	}
	return err
}

func selectStructType(model interface{}, body []byte) error {
	err := json.Unmarshal(body, &model)
	if err != nil {
		return errors.New("Error to unmarshal body")
	}
	return nil
}

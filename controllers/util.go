package controllers

import (
	"encoding/json"
	"fmt"
)

const (
	headerToken         = "Token"
	apiAccessController = "access_controller"
)

func fetchInputPayloadData(input []byte, info interface{}) *failedApiResult {
	if err := json.Unmarshal(input, info); err != nil {
		return newFailedApiResult(
			400, errParsingApiBody, fmt.Errorf("invalid input payload: %s", err.Error()),
		)
	}
	return nil
}

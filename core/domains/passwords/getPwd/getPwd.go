package getPwd

import (
	"encoding/json"
)

type Passwords struct {
	Id string  `dynamodbav:"id"`
	Hash string  `dynamodbav:"hash"`
	Identifier  string  `dynamodbav:"identifier"`
	Tags []string `dynamodbav:"tags"`
	Size float32 `dynamodbav:"size"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParsePassword(obj string) Passwords {
	var passwords Passwords
	json.Unmarshal([]byte(obj), &passwords)
	return passwords
}

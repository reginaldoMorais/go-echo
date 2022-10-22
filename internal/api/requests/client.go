package requests

import (
	"encoding/json"
)

type ClientRequest struct {
	Name  string `json:"name,omitempty" example:"John Doe"  validate:"required"`
	Email string `json:"email,omitempty" example:"john.doe@email.com"  validate:"required"`
}

func (c ClientRequest) ToString() string {
	clientJson, _ := json.Marshal(c)
	return string(clientJson)
}

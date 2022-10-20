package responses

import (
	"encoding/json"
	"go-echo/types"
)

type ClientErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ClientResponse struct {
	Id    types.ID `json:"id"`
	Name  string   `json:"name,omitempty"`
	Email string   `json:"email,omitempty"`
}

func (c ClientResponse) ToString() string {
	clientJson, _ := json.Marshal(c)
	return string(clientJson)
}

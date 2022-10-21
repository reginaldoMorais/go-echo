package responses

import (
	"encoding/json"
	"go-echo/internal/domain/types"
)

type ClientResponse struct {
	Id     types.ID     `json:"id"`
	Name   string       `json:"name,omitempty"`
	Email  string       `json:"email,omitempty"`
	Status types.Status `json:"status"`
}

func (c ClientResponse) ToString() string {
	clientJson, _ := json.Marshal(c)
	return string(clientJson)
}

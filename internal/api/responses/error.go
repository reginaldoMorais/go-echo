package responses

import "encoding/json"

type ClientErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c ClientErrorResponse) ToString() string {
	clientJson, _ := json.Marshal(c)
	return string(clientJson)
}

package types

import (
	"strings"
	"time"
)

const ISO8601 string = "2006-01-02T15:04:05.999Z"

type RequestTime time.Time

func (c *RequestTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05.999+0000", value)
	if err != nil {
		return err
	}
	*c = RequestTime(t)
	return nil
}

func (c RequestTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02T15:04:05.999+0000") + `"`), nil
}

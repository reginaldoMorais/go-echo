package types

import "github.com/google/uuid"

type ID string

var NilID = ID("00000000-0000-0000-0000-000000000000")

func NewID() ID {
	return ID(uuid.New().String())
}

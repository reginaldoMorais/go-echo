package models

import (
	"go-echo/types"
)

type Client struct {
	Id    types.ID `json:"id" bson:"_id" example:"e416dcc0-1ad5-4a7a-94cb-f46f3974950f"  validate:"required"`
	Name  string   `json:"name" bson:"name" example:"John Doe"  validate:"required"`
	Email string   `json:"email" bson:"email" example:"john.doe@email.com"  validate:"required"`
}

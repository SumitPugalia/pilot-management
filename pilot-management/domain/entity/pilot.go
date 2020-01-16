package entity

import (
	guuid "github.com/google/uuid"
)

type Pilot struct {
	Id         guuid.UUID
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
	CreatedAt  int64
	UpdatedAt  int64
	DeletedAt  int64
}

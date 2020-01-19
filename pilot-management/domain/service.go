package domain

import (
	"pilot-management/domain/entity"

	guuid "github.com/google/uuid"
)

type Service interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id guuid.UUID) (entity.Pilot, error)
	CreatePilot(params CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(params UpdatePilotParams) (entity.Pilot, error)
	ChangeStatePilot(id guuid.UUID, state string) (entity.Pilot, error)
	DeletePilot(id guuid.UUID) error
}

type CreatePilotParams struct {
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

type UpdatePilotParams struct {
	Id         guuid.UUID
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

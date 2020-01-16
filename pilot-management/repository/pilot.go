package repository

import (
	"pilot-management/domain"
	"pilot-management/domain/entity"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id string) (entity.Pilot, error)
	CreatePilot(param domain.CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(param domain.UpdatePilotParams) (entity.Pilot, error)
	DeletePilot(id string) error
	StatePilot(id string, state string) (entity.Pilot, error)
}

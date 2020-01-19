package repository

import (
	"pilot-management/domain"
	"pilot-management/domain/entity"

	"github.com/google/uuid"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id uuid.UUID) (entity.Pilot, error)
	CreatePilot(param domain.CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(param domain.UpdatePilotParams) (entity.Pilot, error)
	DeletePilot(id uuid.UUID) error
	ChangeStatePilot(id uuid.UUID, state string) (entity.Pilot, error)
}

package repository

import (
	"pilot-management/domain"
	"pilot-management/domain/entity"

	guuid "github.com/google/uuid"
)

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	GetPilot(id guuid.UUID) (entity.Pilot, error)
	CreatePilot(param domain.CreatePilotParams) (entity.Pilot, error)
	UpdatePilot(param domain.UpdatePilotParams) (entity.Pilot, error)
	DeletePilot(id guuid.UUID) error
}

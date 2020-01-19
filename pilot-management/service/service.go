package service

import (
	"errors"
	"pilot-management/domain"
	"pilot-management/domain/entity"
	"pilot-management/repository"
	"pilot-management/repository/impl/postgresql"

	"github.com/google/uuid"
)

type ServiceImpl struct {
	pilotRepo repository.PilotRepo
}

func MakeServiceImpl() ServiceImpl {
	pilotRepo := postgresql.MakePostgresPilotRepo()
	return ServiceImpl{pilotRepo: &pilotRepo}
}

func (s ServiceImpl) ListPilots() ([]entity.Pilot, error) {
	return s.pilotRepo.ListPilots()
}

func (s ServiceImpl) GetPilot(id uuid.UUID) (entity.Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

func (s ServiceImpl) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.CreatePilot(params)
}

func (s ServiceImpl) UpdatePilot(params domain.UpdatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.UpdatePilot(params)
}

func (s ServiceImpl) DeletePilot(id uuid.UUID) error {
	return s.pilotRepo.DeletePilot(id)
}

func (s ServiceImpl) ChangeStatePilot(id uuid.UUID, state string) (entity.Pilot, error) {
	switch state {
	case "idle":
		return s.pilotRepo.ChangeStatePilot(id, "IDLE")
	case "active":
		return s.pilotRepo.ChangeStatePilot(id, "ACTIVE")
	case "offduty":
		return s.pilotRepo.ChangeStatePilot(id, "OFFDUTY")
	case "break":
		return s.pilotRepo.ChangeStatePilot(id, "BREAK")
	case "suspend":
		return s.pilotRepo.ChangeStatePilot(id, "SUSPEND")
	default:
		return entity.Pilot{}, errors.New("invalid request to change the state")
	}
}

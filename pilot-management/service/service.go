package service

import (
	"pilot-management/domain"
	"pilot-management/domain/entity"
	"pilot-management/repository"
	"pilot-management/repository/impl/postgresql"

	guuid "github.com/google/uuid"
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

func (s ServiceImpl) GetPilot(id guuid.UUID) (entity.Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

func (s ServiceImpl) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.CreatePilot(params)
}

func (s ServiceImpl) UpdatePilot(params domain.UpdatePilotParams) (entity.Pilot, error) {
	return s.pilotRepo.UpdatePilot(params)
}

func (s ServiceImpl) DeletePilot(id guuid.UUID) error {
	return s.pilotRepo.DeletePilot(id)
}

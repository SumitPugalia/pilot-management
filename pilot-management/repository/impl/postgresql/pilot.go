package postgresql

import (
	"math/rand"
	"pilot-management/domain"
	"pilot-management/domain/entity"
	"strconv"
	"time"

	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id         string `db:"id,omitempty"`
	UserId     string `db:"user_id,omitempty"`
	CodeName   string `db:"code_name,omitempty"`
	SupplierId string `db:"supplier_id,omitempty"`
	MarketId   string `db:"market_id,omitempty"`
	ServiceId  string `db:"service_id,omitempty"`
	CreatedAt  int64  `db:"created_at,omitempty"`
	UpdatedAt  int64  `db:"updated_at,omitempty"`
	DeletedAt  int64  `db:"deleted_at,omitempty"`
}

func MakePostgresPilotRepo() PilotRepo {
	return PilotRepo{
		readConn:  getReadConn(),
		writeConn: getWriteConn(),
	}
}

func (repo *PilotRepo) ListPilots() ([]entity.Pilot, error) {
	resultSet := make([]Pilot, 0)
	err := repo.readConn.Collection("pilots").Find().All(&resultSet)
	if err != nil {
		return nil, err
	}
	pilots := make([]entity.Pilot, 0)
	for _, pilot := range resultSet {
		pilots = append(pilots, entity.Pilot(pilot))
	}
	return pilots, nil
}

func (repo *PilotRepo) GetPilot(id string) (entity.Pilot, error) {
	var pilot Pilot
	err := repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	if err != nil {
		return entity.Pilot(pilot), err
	}
	return entity.Pilot(pilot), nil
}

func (repo *PilotRepo) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	pilot := Pilot{
		Id:         strconv.Itoa(rand.Int()),
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		CreatedAt:  time.Now().Unix(),
	}

	_, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot(pilot), err
	}

	return entity.Pilot(pilot), nil
}

func (repo *PilotRepo) UpdatePilot(params domain.UpdatePilotParams) (entity.Pilot, error) {
	pilot := Pilot{
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		UpdatedAt:  time.Now().Unix(),
	}

	res := repo.writeConn.Collection("pilots").Find("id", params.Id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot(pilot), err
	}

	err = repo.readConn.Collection("pilots").Find("id", params.Id).One(&pilot)
	if err != nil {
		return entity.Pilot(pilot), err
	}

	return entity.Pilot(pilot), nil
}

func (repo *PilotRepo) DeletePilot(id string) error {
	pilot := Pilot{
		DeletedAt: time.Now().Unix(),
	}
	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)
	return err
}

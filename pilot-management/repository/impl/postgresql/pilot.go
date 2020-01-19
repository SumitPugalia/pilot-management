package postgresql

import (
	"pilot-management/domain"
	"pilot-management/domain/entity"
	"time"

	guuid "github.com/google/uuid"

	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id         guuid.UUID `db:"id,omitempty"`
	UserId     string     `db:"user_id,omitempty"`
	CodeName   string     `db:"code_name,omitempty"`
	SupplierId string     `db:"supplier_id,omitempty"`
	MarketId   string     `db:"market_id,omitempty"`
	ServiceId  string     `db:"service_id,omitempty"`
	State      string     `db:"state,omitempty"`
	CreatedAt  time.Time  `db:"created_at,omitempty"`
	UpdatedAt  time.Time  `db:"updated_at,omitempty"`
	Deleted    bool       `db:"deleted"`
}

func MakePostgresPilotRepo() PilotRepo {
	return PilotRepo{
		readConn:  getReadConn(),
		writeConn: getWriteConn(),
	}
}

func (repo *PilotRepo) ListPilots() ([]entity.Pilot, error) {
	resultSet := make([]Pilot, 0)
	err := repo.readConn.Collection("pilots").Find(db.Cond{"deleted =": false}).All(&resultSet)
	if err != nil {
		return nil, err
	}
	pilots := make([]entity.Pilot, 0)
	for _, pilot := range resultSet {
		pilots = append(pilots, pilotRowToPilot(pilot))
	}
	return pilots, nil
}

func (repo *PilotRepo) GetPilot(id guuid.UUID) (entity.Pilot, error) {
	var pilot Pilot
	err := repo.readConn.Collection("pilots").Find(db.Cond{"id =": id, "deleted =": false}).One(&pilot)
	if err != nil {
		return entity.Pilot{}, err
	}
	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) CreatePilot(params domain.CreatePilotParams) (entity.Pilot, error) {
	now := time.Now()
	pilot := Pilot{
		Id:         genUUID(),
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		State:      "IDLE",
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	_, err := repo.writeConn.Collection("pilots").Insert(pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) UpdatePilot(params domain.UpdatePilotParams) (entity.Pilot, error) {
	pilot := Pilot{
		Id:         params.Id,
		UserId:     params.UserId,
		CodeName:   params.CodeName,
		SupplierId: params.SupplierId,
		MarketId:   params.MarketId,
		ServiceId:  params.ServiceId,
		UpdatedAt:  time.Now(),
	}

	res := repo.writeConn.Collection("pilots").Find("id", params.Id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}

	err = repo.readConn.Collection("pilots").Find("id", params.Id).One(&pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) ChangeStatePilot(id guuid.UUID, state string) (entity.Pilot, error) {
	pilot := Pilot{
		Id:        id,
		State:     state,
		UpdatedAt: time.Now(),
	}

	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)

	if err != nil {
		return entity.Pilot{}, err
	}

	err = repo.readConn.Collection("pilots").Find("id", id).One(&pilot)
	if err != nil {
		return entity.Pilot{}, err
	}

	return pilotRowToPilot(pilot), nil
}

func (repo *PilotRepo) DeletePilot(id guuid.UUID) error {
	pilot := Pilot{
		Deleted: true,
	}
	res := repo.writeConn.Collection("pilots").Find("id", id)
	err := res.Update(pilot)
	return err
}

func pilotRowToPilot(row Pilot) entity.Pilot {
	return entity.Pilot{
		Id:         row.Id,
		UserId:     row.UserId,
		SupplierId: row.SupplierId,
		MarketId:   row.MarketId,
		ServiceId:  row.ServiceId,
		CodeName:   row.CodeName,
		State:      entity.PilotState(row.State),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func genUUID() guuid.UUID {
	id := guuid.New()
	return id
}

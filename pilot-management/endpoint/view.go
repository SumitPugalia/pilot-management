package endpoint

import (
	"pilot-management/domain/entity"

	"github.com/google/uuid"
)

type PilotView struct {
	Id         uuid.UUID `json:"id"`
	UserId     string    `json:"userId"`
	CodeName   string    `json:"codeName"`
	SupplierId string    `json:"supplierId"`
	MarketId   string    `json:"marketId"`
	ServiceId  string    `json:"serviceId"`
	State      string    `json:"state"`
	CreatedAt  int64     `json:"CreatedAt"`
	UpdatedAt  int64     `json:"UpdatedAt"`
}

func toPilotView(pilot entity.Pilot) PilotView {
	return PilotView{
		Id:         pilot.Id,
		UserId:     pilot.UserId,
		CodeName:   pilot.CodeName,
		SupplierId: pilot.SupplierId,
		MarketId:   pilot.MarketId,
		ServiceId:  pilot.ServiceId,
		State:      string(pilot.State),
		CreatedAt:  pilot.CreatedAt.Unix(),
		UpdatedAt:  pilot.UpdatedAt.Unix(),
	}
}

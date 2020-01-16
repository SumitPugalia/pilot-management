package endpoint

import (
	"pilot-management/domain/entity"
)

type PilotView struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
	State      string `json:"state"`
	CreatedAt  int64  `json:"CreatedAt"`
	UpdatedAt  int64  `json:"UpdatedAt"`
	DeletedAt  int64  `json:"DeletedAt"`
}

func toPilotView(pilot entity.Pilot) PilotView {
	return PilotView{
		Id:         pilot.Id,
		UserId:     pilot.UserId,
		CodeName:   pilot.CodeName,
		SupplierId: pilot.ServiceId,
		MarketId:   pilot.MarketId,
		ServiceId:  pilot.ServiceId,
		State:      pilot.State,
		CreatedAt:  pilot.CreatedAt,
		UpdatedAt:  pilot.UpdatedAt,
		DeletedAt:  pilot.DeletedAt,
	}
}

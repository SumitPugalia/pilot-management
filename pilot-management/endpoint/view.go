package endpoint

import (
	"pilot-management/domain/entity"

	guuid "github.com/google/uuid"
)

type PilotView struct {
	Id         guuid.UUID `json:"id"`
	UserId     string     `json:"userId"`
	CodeName   string     `json:"codeName"`
	SupplierId string     `json:"supplierId"`
	MarketId   string     `json:"marketId"`
	ServiceId  string     `json:"serviceId"`
	CreatedAt  int64      `json:"CreatedAt"`
	UpdatedAt  int64      `json:"UpdatedAt"`
	DeletedAt  int64      `json:"DeletedAt"`
}

func toPilotView(pilot entity.Pilot) PilotView {
	return PilotView{
		Id:         pilot.Id,
		UserId:     pilot.UserId,
		CodeName:   pilot.CodeName,
		SupplierId: pilot.ServiceId,
		MarketId:   pilot.MarketId,
		ServiceId:  pilot.ServiceId,
		CreatedAt:  pilot.CreatedAt,
		UpdatedAt:  pilot.UpdatedAt,
		DeletedAt:  pilot.DeletedAt,
	}
}

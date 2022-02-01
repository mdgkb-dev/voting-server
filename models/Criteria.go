package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Criteria struct {
	bun.BaseModel `bun:"criterias,alias:criterias"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Stars         int           `json:"stars"`
	Event         *Event        `bun:"rel:belongs-to" json:"event"`
	EventID       uuid.NullUUID `bun:"type:uuid" json:"eventId"`
	Ratings       Ratings       `bun:"rel:has-many" json:"ratings"`
}

type Criterias []*Criteria

//
//func (item *Speeches) SetIdForChildren() {
//	for i := range item.FieldValues {
//		item.FieldValues[i].EventApplicationID = item.ID
//	}
//}

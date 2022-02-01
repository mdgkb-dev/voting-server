package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Speech struct {
	bun.BaseModel `bun:"speeches,alias:speeches"`
	ID            uuid.UUID              `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description   string                 `json:"description"`
	Author        string                 `json:"author"`
	Meta          map[string]interface{} `json:"meta"`
	Event         *Event                 `bun:"rel:belongs-to" json:"event"`
	EventID       uuid.NullUUID          `bun:"type:uuid" json:"eventId"`
	Ratings       Ratings                `bun:"rel:has-many" json:"ratings"`
	Attachments   Attachments            `bun:"rel:has-many" json:"attachments"`
}

type Speeches []*Speech

//
//func (item *Speeches) SetIdForChildren() {
//	for i := range item.FieldValues {
//		item.FieldValues[i].EventApplicationID = item.ID
//	}
//}

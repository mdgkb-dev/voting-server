package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Attachment struct {
	bun.BaseModel `bun:"attachments,alias:attachments"`
	ID            uuid.UUID              `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Filename      string                 `json:"filename"`
	Ext           string                 `json:"ext"`
	Path          string                 `json:"path"`
	Pivot         string                 `json:"pivot"`
	Meta          map[string]interface{} `json:"meta"`
	Speech        *Speech                `bun:"rel:belongs-to" json:"speech"`
	SpeechID      uuid.UUID              `bun:"type:uuid" json:"speechId"`
}

type Attachments []*Attachment

//
//func (item *Speeches) SetIdForChildren() {
//	for i := range item.FieldValues {
//		item.FieldValues[i].EventApplicationID = item.ID
//	}
//}

package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Rating struct {
	bun.BaseModel `bun:"ratings,alias:ratings"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Speech        *Speech   `bun:"rel:belongs-to" json:"speech"`
	SpeechID      uuid.UUID `bun:"type:uuid" json:"speechId"`
	Rating        int       `json:"rating"`
	Criteria      *Criteria `bun:"rel:belongs-to" json:"criteria"`
	CriteriaID    uuid.UUID `bun:"type:uuid" json:"criteriaId"`
	User          *User     `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.UUID `bun:"type:uuid" json:"userId"`
}

type Ratings []*Rating

package models

import (
	"time"

	"github.com/google/uuid"
)

type Kudos struct {
	UserId    uuid.UUID `json:"user_id" gorm:"primaryKey"`
	PostId    uuid.UUID `json:"post_id" gorm:"primaryKey;index"`
	CreatedAt time.Time `json:"created_at"`
}

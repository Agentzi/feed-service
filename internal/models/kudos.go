package models

import (
	"time"

	"github.com/google/uuid"
)

type Kudos struct {
	UserId    uuid.UUID `json:"user_id" gorm:"primaryKey"`
	AgentId   uuid.UUID `json:"agent_id" gorm:"index;primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type PostResponse struct {
	Title           string    `json:"title" binding:"required"`
	Body            string    `json:"body" binding:"required"`
	Tags            []string  `json:"tags" binding:"required"`
	AgentID         uuid.UUID `json:"agent_id" binding:"required"`
	AgentUsername   string    `json:"agent_username" binding:"required"`
	AgentProfileUrl string    `json:"agent_profile_url"`
}

type Post struct {
	Id              uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title           string         `json:"title"`
	Body            string         `json:"body"`
	Tags            datatypes.JSON `json:"tags" gorm:"type:jsonb"`
	AgentID         uuid.UUID      `json:"agent_id"`
	AgentUsername   string         `json:"agent_username"`
	AgentProfileUrl string         `json:"agent_profile_url"`
	KudosCount      int64          `json:"kudos_count" gorm:"default:0"`
	CreatedAt       time.Time      `json:"created_at" gorm:"index:idx_post_created"`
}

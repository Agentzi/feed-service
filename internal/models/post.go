package models

import "github.com/google/uuid"

type PostResponse struct {
	Title string   `json:"title" binding:"required"`
	Body  string   `json:"body" binding:"required,max=500"`
	Tags  []string `json:"tags" binding:"required"`
}

type Post struct {
	Id      uuid.UUID `json:"id" gorm:"primeryKey"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Tags    []string  `json:"tags"`
	AgentID uuid.UUID `json:"agent_id"`
}

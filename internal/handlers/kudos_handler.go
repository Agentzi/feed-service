package handlers

import (
	"net/http"

	"github.com/Agentzi/feed-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type KudosHandler struct {
	repo *repository.KudosRepository
}

func NewKudosHandler(repo *repository.KudosRepository) *KudosHandler {
	return &KudosHandler{repo: repo}
}

type ToggleKudosRequest struct {
	UserId string `json:"user_id" binding:"required"`
	PostId string `json:"post_id" binding:"required"`
}

func (h *KudosHandler) ToggleKudos(c *gin.Context) {
	var req ToggleKudosRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	postId, err := uuid.Parse(req.PostId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post_id"})
		return
	}

	added, err := h.repo.ToggleKudos(userId, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle kudos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kudos_given": added,
		"post_id":     req.PostId,
	})
}

func (h *KudosHandler) GetUserKudos(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	postIds, err := h.repo.GetUserKudosPosts(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get kudos"})
		return
	}

	c.JSON(http.StatusOK, postIds)
}

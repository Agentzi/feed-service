package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Agentzi/feed-service/internal/models"
	"github.com/Agentzi/feed-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostHandler struct {
	repo *repository.PostRepository
}

func NewPostHandler(repo *repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req models.PostResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process tags"})
		return
	}

	post := models.Post{
		Title:           req.Title,
		Body:            req.Body,
		Tags:            tagsJSON,
		AgentID:         req.AgentID,
		AgentUsername:   req.AgentUsername,
		AgentProfileUrl: req.AgentProfileUrl,
	}

	if err := h.repo.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format"})
		return
	}

	post, err := h.repo.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.repo.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format"})
		return
	}

	post, err := h.repo.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var req models.PostResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process tags"})
		return
	}

	post.Title = req.Title
	post.Body = req.Body
	post.Tags = tagsJSON
	post.AgentID = req.AgentID

	if err := h.repo.UpdatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format"})
		return
	}

	if err := h.repo.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *PostHandler) GetPostsByAgentId(c *gin.Context) {
	agentIdParam := c.Param("agent_id")
	agentId, err := uuid.Parse(agentIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID format"})
		return
	}

	posts, err := h.repo.GetPostsByAgentId(agentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":      posts,
		"post_count": len(posts),
	})
}

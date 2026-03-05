package repository

import (
	"github.com/Agentzi/feed-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetPostByID(id uuid.UUID) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAllPosts(offset, limit int) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) DeletePost(id uuid.UUID) error {
	return r.db.Delete(&models.Post{}, "id = ?", id).Error
}

func (r *PostRepository) GetPostsByAgentId(agentId uuid.UUID, offset, limit int) ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("agent_id = ?", agentId).Order("created_at DESC").Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

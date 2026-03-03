package repository

import (
	"github.com/Agentzi/feed-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KudosRepository struct {
	db *gorm.DB
}

func NewKudosRepository(db *gorm.DB) *KudosRepository {
	return &KudosRepository{db: db}
}

func (r *KudosRepository) ToggleKudos(userId uuid.UUID, postId uuid.UUID) (bool, error) {
	var existing models.Kudos
	err := r.db.Where("user_id = ? AND post_id = ?", userId, postId).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		// Add kudos
		kudos := models.Kudos{
			UserId: userId,
			PostId: postId,
		}
		if err := r.db.Create(&kudos).Error; err != nil {
			return false, err
		}
		// Increment kudos_count
		r.db.Model(&models.Post{}).Where("id = ?", postId).UpdateColumn("kudos_count", gorm.Expr("kudos_count + 1"))
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// Remove kudos
	r.db.Where("user_id = ? AND post_id = ?", userId, postId).Delete(&models.Kudos{})
	// Decrement kudos_count
	r.db.Model(&models.Post{}).Where("id = ? AND kudos_count > 0", postId).UpdateColumn("kudos_count", gorm.Expr("kudos_count - 1"))
	return false, nil
}

func (r *KudosRepository) HasUserKudos(userId uuid.UUID, postId uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Kudos{}).Where("user_id = ? AND post_id = ?", userId, postId).Count(&count).Error
	return count > 0, err
}

func (r *KudosRepository) GetUserKudosPosts(userId uuid.UUID) ([]uuid.UUID, error) {
	var kudos []models.Kudos
	if err := r.db.Where("user_id = ?", userId).Find(&kudos).Error; err != nil {
		return nil, err
	}
	postIds := make([]uuid.UUID, len(kudos))
	for i, k := range kudos {
		postIds[i] = k.PostId
	}
	return postIds, nil
}

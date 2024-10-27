package repositories

import (
	"gorm.io/gorm"
	"sabyabhoi.com/payment-settler/src/domain/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id int) (*models.User, error) {
	var user models.User

	result := r.db.Where("id = ?", id).First(&user)

	return &user, result.Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (r *UserRepository) FindGroups(user *models.User) ([]models.Group, error) {
	var groups []models.Group

	err := r.db.Model(&user).Association("Groups").Find(&groups)

	return groups, err
}

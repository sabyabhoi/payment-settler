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

func (r *UserRepository) FindUsers() ([]models.User, error) {
	var users []models.User

	result := r.db.Find(&users)

	return users, result.Error
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

func (r *UserRepository) CreateUser(user *models.User) (uint, error) {
	result := r.db.Create(user)

	return user.ID, result.Error
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	result := r.db.First(user, user.ID)
	if result.Error != nil {
		return user, result.Error
	}

	result = r.db.Save(user)
	return user, result.Error
}

func (r *UserRepository) DeleteUser(uid uint) error {
	result := r.db.Delete(&models.User{}, uid)

	return result.Error
}

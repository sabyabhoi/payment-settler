package repositories

import (
	"gorm.io/gorm"
	"sabyabhoi.com/payment-settler/src/domain/models"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) FindAllGroups() ([]models.Group, error) {
	var groups []models.Group

	result := r.db.Find(&groups)

	return groups, result.Error
}

func (r *GroupRepository) FindById(id int) (*models.Group, error) {
	var group models.Group

	result := r.db.Where("id = ?", id).First(&group)

	return &group, result.Error
}

func (r *GroupRepository) FindAllUsersInGroup(group *models.Group) ([]models.User, error) {
	var users []models.User

	err := r.db.Model(&group).Association("Users").Find(&users)

	return users, err
}

func (r *GroupRepository) CreateGroup(group *models.Group) (uint, error) {
	result := r.db.Create(group)

	return group.ID, result.Error
}

func (r *GroupRepository) UpdateGroup(group *models.Group) (*models.Group, error) {
	result := r.db.First(group, group.ID)
	if result.Error != nil {
		return group, result.Error
	}

	result = r.db.Save(group)
	return group, result.Error
}

func (r *GroupRepository) DeleteGroup(groupId uint) error {
	result := r.db.Delete(&models.Group{}, groupId)

	return result.Error
}

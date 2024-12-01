package services

import (
	"sabyabhoi.com/payment-settler/src/domain/models"
	"sabyabhoi.com/payment-settler/src/infrastructure/repositories"
)

type GroupService struct {
	repo repositories.GroupRepository
}

func NewGroupService(repo *repositories.GroupRepository) *GroupService {
	return &GroupService{repo: *repo}
}

func (s *GroupService) GetAllGroups() ([]models.Group, error) {
	return s.repo.FindAllGroups()
}

func (s *GroupService) GetGroupById(id int) (*models.Group, error) {
	return s.repo.FindById(id)
}

func (s *GroupService) CreateGroup(group *models.Group) (uint, error) {
	return s.repo.CreateGroup(group)
}

func (s *GroupService) UpdateGroup(group *models.Group) (*models.Group, error) {
	return s.repo.UpdateGroup(group)
}

func (s *GroupService) DeleteGroup(groupId uint) error {
	return s.repo.DeleteGroup(groupId)
}

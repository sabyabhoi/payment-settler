package services

import (
	"sabyabhoi.com/payment-settler/src/domain/models"
	"sabyabhoi.com/payment-settler/src/infrastructure/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: *repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAllUsers()
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetAllGroupsForUserId(id int) ([]models.Group, error) {
	u, err := s.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindAllGroupsForUser(u)
}

func (s *UserService) CreateUser(user *models.User) (uint, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) AddGroupToUser(user *models.User, group *models.Group) error {
	return s.repo.AddGroupToUser(user, group)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(uid uint) error {
	return s.repo.DeleteUser(uid)
}

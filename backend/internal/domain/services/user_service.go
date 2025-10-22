package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"errors"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetUser(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Email != nil {
		// Check if email is already taken by another user
		existingUser, _ := s.userRepo.GetByEmail(*req.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already taken")
		}
		user.Email = *req.Email
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

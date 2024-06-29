package service

import (
	models "bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	logger "bbdk/infrastructure/log"
	"bbdk/utils/encrypt"
	"errors"
)

var (
	ErrInvalidPage = errors.New("invalid page or page size")
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetAllUsers(page, pageSize int) ([]*models.User, error)
}
type userService struct {
	userRepo userRepo.Repository
	logger   logger.Logger
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo userRepo.Repository, logger logger.Logger) UserService {
	return &userService{userRepo: userRepo, logger: logger}
}

func (s *userService) CreateUser(user *models.User) error {
	user.Password = encrypt.HashSHA256(user.Password)

	if err := s.userRepo.CreateUser(user); err != nil {
		if errors.Is(err, userRepo.ErrAlreadyExist) {
			return userRepo.ErrAlreadyExist
		}
		s.logger.Errorf("failed to create user:%s", err.Error())
		return err
	}
	return nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return nil, err
		}
		s.logger.Errorf("failed to get user:%s", err.Error())
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
	user.Password = encrypt.HashSHA256(user.Password)
	if err := s.userRepo.UpdateUser(user); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return err
		} else if errors.Is(err, userRepo.ErrAlreadyExist) {
			return userRepo.ErrAlreadyExist
		}
		s.logger.Errorf("failed to update user:%s", err.Error())
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id uint) error {
	if err := s.userRepo.DeleteUser(id); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return err
		}
		s.logger.Errorf("failed to delete user:%s", err.Error())
		return err
	}
	return nil
}

func (s *userService) GetAllUsers(page, pageSize int) ([]*models.User, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, ErrInvalidPage
	}
	offset := (page - 1) * pageSize
	limit := pageSize

	users, err := s.userRepo.GetAll(offset, limit)
	if err != nil {
		s.logger.Errorf("failed to get all users: %s", err.Error())
		return nil, err
	}
	return users, nil
}

package user

import (
	"bbdk/domain/entity"
	"errors"
)

var (
	ErrAlreadyExist     = errors.New("already exist")
	ErrNotFound         = errors.New("not found")
	ErrNotEnoughBalance = errors.New("not enough balance")
)

type Repository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	GetAll(offset, limit int) ([]*entity.User, error)
	DepositCrypto(cu *entity.CoinUser) error
	WithdrawCrypto(cu *entity.CoinUser) error
}

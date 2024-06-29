package coin

import (
	"bbdk/domain/entity"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

type Repository interface {
	CreateCoin(coin *entity.Coin) error
	GetAll() ([]*entity.Coin, error)
}

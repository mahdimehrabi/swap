package coin_price

import (
	"bbdk/domain/entity"
	"context"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

type Repository interface {
	SetCoin(ctx context.Context, coin *entity.Coin) error
	GetCoin(ctx context.Context, id uint) (*entity.Coin, error)
}

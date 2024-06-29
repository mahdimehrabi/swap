package transaction

import (
	"bbdk/domain/entity"
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	GetTransaction(ctx context.Context, id uuid.UUID) (*entity.Transaction, error)
}

package redis

import (
	"bbdk/domain/entity"
	transRepo "bbdk/domain/repository/transaction"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type TransactionRepository struct {
	redis *redis.Client
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(rdb *redis.Client) *TransactionRepository {
	return &TransactionRepository{redis: rdb}
}

func (t TransactionRepository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	return t.redis.Set(ctx, transaction.ID.String(), transaction, time.Minute).Err()
}

func (t TransactionRepository) GetTransaction(ctx context.Context, id uuid.UUID) (*entity.Transaction, error) {
	var transaction entity.Transaction
	cmd := t.redis.Get(ctx, id.String())
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, transRepo.ErrNotFound
		}
		return nil, err
	}
	if err := cmd.Scan(&transaction); err != nil {
		return nil, err
	}
	return &transaction, nil
}

package redis

import (
	"bbdk/domain/entity"
	transRepo "bbdk/domain/repository/coin"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type CoinRepository struct {
	redis *redis.Client
}

// NewCoinRepository creates a new instance of CoinRepository
func NewCoinRepository(rdb *redis.Client) *CoinRepository {
	return &CoinRepository{redis: rdb}
}

//I don't have time to use uuid for coin entity too so just I just use normal uint id (we won't get duplicate problem because we only store coin by id on redis) :(
//I don't have time to use uuid for coin entity too so just I just use normal uint id (we won't get duplicate problem because we only store coin by id on redis) :(
//I don't have time to use uuid for coin entity too so just I just use normal uint id (we won't get duplicate problem because we only store coin by id on redis) :(
//I don't have time to use uuid for coin entity too so just I just use normal uint id (we won't get duplicate problem because we only store coin by id on redis) :(

func (t CoinRepository) SetCoin(ctx context.Context, coin *entity.Coin) error {
	return t.redis.Set(ctx, fmt.Sprintf("%d", coin.ID), coin, time.Minute*3).Err()
}

func (t CoinRepository) GetCoin(ctx context.Context, id uint) (*entity.Coin, error) {
	var coin entity.Coin
	cmd := t.redis.Get(ctx, fmt.Sprintf("%d", id))
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, transRepo.ErrNotFound
		}
		return nil, err
	}
	if err := cmd.Scan(&coin); err != nil {
		return nil, err
	}
	return &coin, nil
}

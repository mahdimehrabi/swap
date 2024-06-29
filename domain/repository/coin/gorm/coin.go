package gorm

import (
	"bbdk/domain/entity"
	coinRepo "bbdk/domain/repository/coin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type CoinRepository struct {
	db *gorm.DB
}

// NewCoinRepository creates a new instance of CoinRepository
func NewCoinRepository(db *gorm.DB) *CoinRepository {
	return &CoinRepository{db: db}
}

func (r *CoinRepository) CreateCoin(coin *entity.Coin) error {
	if err := r.db.Create(coin).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return coinRepo.ErrAlreadyExist
			}
		}
		return err
	}
	return nil
}

func (r *CoinRepository) GetAll() ([]*entity.Coin, error) {
	var coins []*entity.Coin
	if err := r.db.Find(&coins).Error; err != nil {
		return nil, err
	}
	return coins, nil
}

package gorm

import (
	"bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"math/big"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) DepositCrypto(coinUser *entity.CoinUser) error {
	//there is a faster way directly using sql without ORM
	return r.db.Transaction(func(tx *gorm.DB) error {
		//transaction to handle race condition
		currentCUser := entity.NewCoinUser(coinUser.CoinID, coinUser.UserID)
		if err := tx.Find(currentCUser).Error; err != nil {
			return err
		}
		coinUser.I.Add(coinUser.I, currentCUser.I)
		coinUser.UpdateAmount()
		if err := tx.Save(coinUser).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *UserRepository) Swap(transaction *entity.Transaction) (*entity.CoinUser, *entity.CoinUser, error) {
	currentSrc := entity.NewCoinUser(transaction.SrcCoinID, transaction.UserID)
	currentDest := entity.NewCoinUser(transaction.DestCoinID, transaction.UserID)

	//there is a faster way directly using sql without ORM
	err := r.db.Transaction(func(tx *gorm.DB) error {
		//transaction to handle race condition
		if err := tx.Find(currentSrc).Error; err != nil {
			return err
		}
		if err := currentSrc.FromIntString(currentSrc.Amount); err != nil {
			return err
		}

		currentSrc.I.Sub(currentSrc.I, transaction.SrcCoinA.I)
		fmt.Println(currentSrc.I, transaction.SrcCoinA.I)
		if currentSrc.I.Cmp(big.NewInt(0)) == -1 {
			return userRepo.ErrNotEnoughBalance
		}
		currentSrc.Amount = currentSrc.I.String()

		if err := tx.Save(currentSrc).Error; err != nil {
			return err
		}

		if err := tx.Find(currentDest).Error; err != nil {
			return err
		}
		if err := currentDest.FromIntString(currentDest.Amount); err != nil {
			return err
		}

		currentDest.I.Add(transaction.DestCoinA.I, currentDest.I)
		currentDest.Amount = currentDest.I.String()

		if err := tx.Save(currentDest).Error; err != nil {
			return err
		}
		return nil
	})
	return currentSrc, currentDest, err
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return userRepo.ErrAlreadyExist
			}
		}
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, userRepo.ErrNotFound
		}
		return nil, err

	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	tx := r.db.Where("id", user.ID).UpdateColumns(user)

	if err := tx.Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return userRepo.ErrAlreadyExist
			}
		}
		return err
	}
	if tx.RowsAffected < 1 {
		return userRepo.ErrNotFound

	}
	return nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	tx := r.db.Delete(&entity.User{}, id)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected < 1 {
		return userRepo.ErrNotFound

	}
	return nil
}

func (r *UserRepository) GetAll(offset, limit int) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

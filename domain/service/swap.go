package service

import (
	models "bbdk/domain/entity"
	"bbdk/domain/repository/coin_price"
	transactionRepo "bbdk/domain/repository/transaction"
	logger "bbdk/infrastructure/log"
	"context"
	"errors"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, user *models.User, srcCoin *models.Coin, destCoin *models.Coin) (*models.Transaction, error)
	CommitTransaction(ctx context.Context, transaction *models.Transaction) error
}
type transactionService struct {
	transactionRepo transactionRepo.Repository
	cpr             coin_price.Repository
	logger          logger.Logger
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo transactionRepo.Repository, cpr coin_price.Repository, logger logger.Logger) TransactionService {
	return &transactionService{transactionRepo: transactionRepo, logger: logger, cpr: cpr}
}

func (t transactionService) CreateTransaction(ctx context.Context, user *models.User,
	srcCoin *models.Coin, destCoin *models.Coin) (*models.Transaction, error) {
	transaction := models.NewTransaction(user.ID, srcCoin.ID, destCoin.ID)

	srcCoin, err := t.cpr.GetCoin(ctx, srcCoin.ID)
	if err != nil {
		t.logger.Errorf("failed to coin %s", err.Error())
		return nil, err
	}
	destCoin, err = t.cpr.GetCoin(ctx, destCoin.ID)
	if err != nil {
		t.logger.Errorf("failed to coin %s", err.Error())
		return nil, err
	}
	if srcCoin.LastPriceUpdate.Unix() < time.Now().Unix()-180 || destCoin.LastPriceUpdate.Unix() < time.Now().Unix()-180 {
		t.logger.Errorf("prices are not up to date")

		return nil, errors.New("prices are not up to date")
	}

	transaction.SrcCoinPrice = srcCoin.USDPrice
	transaction.DestCoinPrice = destCoin.USDPrice
	if err := transaction.SrcCoin.FromIntString(srcCoin.USDPrice); err != nil {
		return nil, err
	}
	if err := transaction.DestCoin.FromIntString(destCoin.USDPrice); err != nil {
		return nil, err
	}

	if err := t.transactionRepo.CreateTransaction(ctx, transaction); err != nil {
		t.logger.Errorf("failed to create transaction %s", err.Error())
		return nil, err
	}
	return transaction, nil
}

func (t transactionService) CommitTransaction(ctx context.Context, transaction *models.Transaction) error {

	//TODO implement me
	panic("implement me")
}

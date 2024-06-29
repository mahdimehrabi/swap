package service

import (
	models "bbdk/domain/entity"
	"bbdk/domain/repository/coin_price"
	transactionRepo "bbdk/domain/repository/transaction"
	"bbdk/domain/repository/user"
	logger "bbdk/infrastructure/log"
	"context"
	"errors"
	"math"
	"math/big"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, user *models.User, srcCoin *models.Coin, destCoin *models.Coin) (*models.Transaction, error)
	CommitTransaction(ctx context.Context, transaction *models.Transaction) error
}
type transactionService struct {
	transactionRepo transactionRepo.Repository
	cpr             coin_price.Repository
	userRepository  user.Repository
	logger          logger.Logger
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo transactionRepo.Repository, cpr coin_price.Repository,
	logger logger.Logger, userRepo user.Repository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo,
		logger: logger, cpr: cpr, userRepository: userRepo}
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
	transaction.SrcCoinP.I = srcCoin.I
	transaction.DestCoinPrice = destCoin.USDPrice
	transaction.DestCoinP.I = destCoin.I

	price := big.NewFloat(transaction.SrcCoinP.ToFloat())
	temp := big.NewFloat(transaction.SrcCoinP.ToFloat())
	price = price.Quo(price, temp)
	temp = temp.SetFloat64(transaction.SrcCoinA.ToFloat())
	destAmount := price.Mul(price, temp)
	destAmount = destAmount.Mul(destAmount, big.NewFloat(math.Pow10(18)))
	destAmount.Int(transaction.DestCoinA.I)

	if err := t.transactionRepo.CreateTransaction(ctx, transaction); err != nil {
		t.logger.Errorf("failed to create transaction %s", err.Error())
		return nil, err
	}
	return transaction, nil
}

func (t transactionService) CommitTransaction(ctx context.Context, transaction *models.Transaction) error {
	panic("not implemented")
	//	transaction, err := t.transactionRepo.GetTransaction(ctx, transaction.ID)
	//	if err != nil {
	//		if errors.Is(err, transactionRepo.ErrNotFound) {
	//			return err
	//		}
	//		t.logger.Errorf("failed to get transaction err:%s", err.Error())
	//		return err
	//	}
	//	price := transaction.SrcCoinP.I.Int64() / transaction.DestCoinP.I
	//	coinSrc := models.NewCoinUser(transaction.SrcCoinID, transaction.UserID)
	//
	//	return nil
}

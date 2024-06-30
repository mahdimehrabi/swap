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
	CreateTransaction(ctx context.Context, user *models.User,
		srcCoin *models.Coin, destCoin *models.Coin, srcCoinAmount float64) (*models.Transaction, error)
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
	srcCoin *models.Coin, destCoin *models.Coin, srcCoinAmount float64) (*models.Transaction, error) {
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
	transaction.SrcCoinA.FromFloat(srcCoinAmount)
	transaction.SrcCoinAmount = transaction.SrcCoinA.ToIntString()

	transaction.SrcCoinPrice = srcCoin.USDPrice
	transaction.SrcCoinP.I = srcCoin.I
	transaction.DestCoinPrice = destCoin.USDPrice
	transaction.DestCoinP.I = destCoin.I

	srcPrice := big.NewFloat(0).SetInt(transaction.SrcCoinP.I)   //10^2 srcPrice
	destPrice := big.NewFloat(0).SetInt(transaction.DestCoinP.I) // 10^2 dest srcPrice
	srcPrice = srcPrice.Quo(srcPrice, big.NewFloat(math.Pow10(2)))
	destPrice = destPrice.Quo(destPrice, big.NewFloat(math.Pow10(2)))

	wholePrice := srcPrice.Mul(srcPrice, big.NewFloat(transaction.SrcCoinA.ToFloat()))

	destAmount := big.NewFloat(0)
	destAmount = destAmount.Quo(wholePrice, destPrice)
	destAmount = destAmount.Mul(destAmount, big.NewFloat(math.Pow10(18)))
	destAmount.Int(transaction.DestCoinA.I)
	transaction.DestCoinAmount = transaction.DestCoinA.ToIntString()

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

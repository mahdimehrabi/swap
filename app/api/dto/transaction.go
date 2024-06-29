package dto

import (
	"bbdk/domain/entity"
	"github.com/google/uuid"
)

type Transaction struct {
	ID            uuid.UUID `json:"ID"`
	CreatedAt     uint64    `json:"createdAt"` //the transaction would be accepted after one minute duration of this time
	DestCoinID    uint      `json:"destCoinID"`
	SrcCoinID     uint      `json:"srcCoinID"`
	UserID        uint      `json:"userID"`
	SrcCoinPrice  float64   `json:"srcCoinPrice"`
	DestCoinPrice float64   `json:"destCoinPrice"`
}

func (t *Transaction) FromEntity(transaction *entity.Transaction) {
	t.ID = transaction.ID
	t.CreatedAt = transaction.CreatedAt
	t.DestCoinID = transaction.DestCoinID
	t.SrcCoinID = transaction.SrcCoinID
	t.UserID = transaction.UserID
	t.SrcCoinPrice = transaction.SrcCoin.ToFloat()
	t.DestCoinPrice = transaction.DestCoin.ToFloat()
}

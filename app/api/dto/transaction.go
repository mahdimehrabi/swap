package dto

import (
	"bbdk/domain/entity"
	"github.com/google/uuid"
)

type Transaction struct {
	ID             uuid.UUID `json:"ID"`
	CreatedAt      uint64    `json:"createdAt"` //the transaction would be accepted after one minute duration of this time
	DestCoinID     uint      `json:"destCoinID"`
	SrcCoinID      uint      `json:"srcCoinID"`
	UserID         uint      `json:"userID"`
	SrcCoinPrice   float64   `json:"srcCoinPrice"`
	DestCoinPrice  float64   `json:"destCoinPrice"`
	SrcCoinAmount  float64   `json:"srcCoinAmount"`
	DestCoinAmount float64   `json:"destCoinAmount"`
}

func (t *Transaction) FromEntity(transaction *entity.Transaction) {
	t.ID = transaction.ID
	t.CreatedAt = transaction.CreatedAt
	t.DestCoinID = transaction.DestCoinID
	t.SrcCoinID = transaction.SrcCoinID
	t.UserID = transaction.UserID
	t.SrcCoinPrice = transaction.SrcCoinP.ToFloat()
	t.DestCoinPrice = transaction.DestCoinP.ToFloat()
	t.SrcCoinAmount = transaction.SrcCoinA.ToFloat()
	t.DestCoinAmount = transaction.DestCoinA.ToFloat()
}

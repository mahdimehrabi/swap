package entity

import (
	"bbdk/domain/entity/currency"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// only redis
type Transaction struct {
	ID            uuid.UUID     `json:"ID"`
	CreatedAt     uint64        `json:"createdAt"` //the transaction would be accepted after one minute duration of this time
	DestCoinID    uint          `json:"destCoinID"`
	SrcCoinID     uint          `json:"srcCoinID"`
	UserID        uint          `json:"userID"`
	SrcCoinPrice  string        `json:"srcCoinPrice"`
	DestCoinPrice string        `json:"destCoinPrice"`
	SrcCoin       *currency.USD `json:"-"`
	DestCoin      *currency.USD `json:"-"`
}

func NewTransaction(userID uint, srcCoinID uint, destCoinID uint) *Transaction {
	t := &Transaction{
		CreatedAt:  uint64(time.Now().Unix()),
		SrcCoinID:  srcCoinID,
		DestCoinID: destCoinID,
		UserID:     userID,
		SrcCoin:    currency.NewUSD(),
		DestCoin:   currency.NewUSD(),
		ID:         uuid.New(),
	}
	return t
}

func (t *Transaction) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *Transaction) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

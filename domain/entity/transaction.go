package entity

import (
	"bbdk/domain/entity/currency"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// only redis
type Transaction struct {
	ID             uuid.UUID        `json:"ID"`
	CreatedAt      uint64           `json:"createdAt"` //the transaction would be accepted after one minute duration of this time
	DestCoinID     uint             `json:"destCoinID"`
	SrcCoinID      uint             `json:"srcCoinID"`
	UserID         uint             `json:"userID"`
	SrcCoinPrice   string           `json:"srcCoinPrice"`
	DestCoinPrice  string           `json:"destCoinPrice"`
	SrcCoinP       *currency.USD    `json:"-"`
	DestCoinP      *currency.USD    `json:"-"`
	SrcCoinAmount  string           `json:"srcCoinAmount"`
	DestCoinAmount string           `json:"destCoinAmount"`
	SrcCoinA       *currency.Crypto `json:"-"`
	DestCoinA      *currency.Crypto `json:"-"`
}

func NewTransaction(userID uint, srcCoinID uint, destCoinID uint) *Transaction {
	t := &Transaction{
		CreatedAt:  uint64(time.Now().Unix()),
		SrcCoinID:  srcCoinID,
		DestCoinID: destCoinID,
		UserID:     userID,
		SrcCoinP:   currency.NewUSD(),
		DestCoinP:  currency.NewUSD(),
		SrcCoinA:   currency.NewCrypto(),
		DestCoinA:  currency.NewCrypto(),
		ID:         uuid.New(),
	}
	return t
}

func (t *Transaction) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *Transaction) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	t.SrcCoinP = currency.NewUSD()
	t.DestCoinP = currency.NewUSD()
	t.SrcCoinA = currency.NewCrypto()
	t.DestCoinA = currency.NewCrypto()
	if err := t.SrcCoinP.FromIntString(t.SrcCoinPrice); err != nil {
		return err
	}

	if err := t.DestCoinP.FromIntString(t.DestCoinPrice); err != nil {
		return err
	}

	if err := t.SrcCoinA.FromIntString(t.SrcCoinAmount); err != nil {
		return err
	}
	if err := t.DestCoinA.FromIntString(t.DestCoinAmount); err != nil {
		return err
	}
	return nil
}

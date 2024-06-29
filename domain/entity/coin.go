package entity

import (
	"bbdk/domain/entity/currency"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Coin struct {
	gorm.Model
	*currency.USD   `gorm:"-" json:"-"`
	Symbol          string     `gorm:"unique"`
	CoinUsers       []CoinUser `gorm:"foreignKey:CoinID"`
	LastPriceUpdate time.Time  `gorm:"-"`  //just redis
	USDPrice        string     `gorm:"-" ` //just redis
}

func (c *Coin) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	c.USD = currency.NewUSD()
	if err := c.FromIntString(c.USDPrice); err != nil {
		return err
	}
	return nil
}

func (c *Coin) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func NewCoin(symbol string) *Coin {
	return &Coin{USD: currency.NewUSD(), Symbol: symbol}
}

// SetAmount only use this function to update amount
func (c *Coin) SetAmount(amount float64) {
	c.FromFloat(amount)
	c.UpdateAmount()
}

// UpdateAmount always use this after updating amount
func (c *Coin) UpdateAmount() {
	c.USDPrice = c.ToIntString()
}

func (c *Coin) AddAmount(amount float64) {
	c.Add(amount)
	c.UpdateAmount()
}

func (c *Coin) DivideAmount(amount float64) {
	c.Divide(amount)
	c.UpdateAmount()
}

func (c *Coin) MultiplyAmount(amount float64) {
	c.Multiply(amount)
	c.UpdateAmount()
}

func (c *Coin) SubAmount(amount float64) {
	c.Sub(amount)
	c.UpdateAmount()
}

type CoinUser struct {
	*currency.Crypto `gorm:"-"`
	CoinID           uint   `gorm:"primaryKey"`
	UserID           uint   `gorm:"primaryKey"`
	Amount           string `gorm:"not null"`
	Coin             Coin   `gorm:"foreignKey:CoinID;constraint:OnDelete:CASCADE;"`
	User             User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func NewCoinUser(coinID uint, userID uint) *CoinUser {
	return &CoinUser{
		Crypto: currency.NewCrypto(),
		CoinID: coinID,
		UserID: userID,
	}
}

// SetAmount only use this function to update amount
func (c *CoinUser) SetAmount(amount float64) {
	c.FromFloat(amount)
	c.UpdateAmount()
}

// UpdateAmount always use this after updating amount
func (c *CoinUser) UpdateAmount() {
	c.Amount = c.ToIntString()
}

func (c *CoinUser) AddAmount(amount float64) {
	c.Add(amount)
	c.UpdateAmount()
}

func (c *CoinUser) DivideAmount(amount float64) {
	c.Divide(amount)
	c.UpdateAmount()
}

func (c *CoinUser) MultiplyAmount(amount float64) {
	c.Multiply(amount)
	c.UpdateAmount()
}

func (c *CoinUser) SubAmount(amount float64) {
	c.Sub(amount)
	c.UpdateAmount()
}

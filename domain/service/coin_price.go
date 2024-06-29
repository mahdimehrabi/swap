package service

import (
	crypto_compare "bbdk/domain/dto/crypto-compare"
	"bbdk/domain/entity"
	"bbdk/domain/entity/currency"
	"bbdk/domain/repository/coin"
	"bbdk/domain/repository/coin_price"
	logger "bbdk/infrastructure/log"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	queueLen    = 100
	workerCount = 500
)

type CoinPrice interface {
	Start() //starts updating price
}
type coinPrice struct {
	logger          logger.Logger
	apiURL          string
	queue           chan *entity.Coin
	coinRepository  coin.Repository
	priceRepository coin_price.Repository
}

func NewCoinPrice(logger logger.Logger, apiURL string, coinRepository coin.Repository, cpr coin_price.Repository) CoinPrice {
	return &coinPrice{logger: logger, apiURL: apiURL, coinRepository: coinRepository,
		queue: make(chan *entity.Coin, queueLen), priceRepository: cpr}
}

func (c *coinPrice) Start() {
	go c.allocator()
	for i := 0; i < workerCount; i++ {
		go c.worker()
	}
}

func (c *coinPrice) allocator() {
	for {
		//add coin update price task for each coin to queue every 5 sec
		time.Sleep(3 * time.Second)
		coins, err := c.coinRepository.GetAll()
		if err != nil {
			c.logger.Errorf("err getting coins %s", err.Error())
			continue
		}
		for i := 0; i < len(coins); i++ {
			c.queue <- coins[i]
		}
	}
}

func (c *coinPrice) worker() {
	for coin := range c.queue {
		apiURL := strings.Replace(c.apiURL, "XXCoinSymbolXX", coin.Symbol, -1)

		resp, err := http.Get(apiURL)
		if err != nil {
			c.logger.Errorf("error in connecting to price api:%s", err.Error())
			return
		}
		if resp.StatusCode != http.StatusOK {
			errStr := fmt.Sprintf("http status code %d from price api", resp.StatusCode)
			c.logger.Errorf(errStr)
		}
		response := crypto_compare.Response{}
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&response); err != nil {
			c.logger.Errorf("failed to decode api response %s", err.Error())
			return
		}
		if response.Raw == nil {
			c.logger.Errorf("failed to decode api response")
			return
		}
		coin.USD = currency.NewUSD()
		coin.SetAmount(response.Raw.Price)
		coin.LastPriceUpdate = time.Now()
		if err := c.priceRepository.SetCoin(context.Background(), coin); err != nil {
			c.logger.Errorf("failed to set coin price %s", err.Error())
			return
		}
		fmt.Printf("updated price for %s=%f\n", coin.Symbol, coin.ToFloat())
	}
}

package seeder

import (
	"bbdk/domain/entity"
	coinGorm "bbdk/domain/repository/coin/gorm"
	userGorm "bbdk/domain/repository/user/gorm"
	"bbdk/infrastructure/godotenv"
	"bbdk/utils/encrypt"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func Seed() error {
	env := godotenv.NewEnv()
	env.Load()
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// Seed the random number generator
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	userRepo := userGorm.NewUserRepository(db)
	coinRepo := coinGorm.NewCoinRepository(db)

	coinStrings := []string{"BTC", "ETH", "DOGE", "XRP", "USDT"}
	coins := make([]entity.Coin, 0)
	for _, coin := range coinStrings {
		err = coinRepo.CreateCoin(&entity.Coin{
			Symbol: coin,
		})
		if err != nil {
			log.Printf("Failed to create coin: %s", err.Error())
			return err
		}
		var coinEnt entity.Coin
		err = db.Where("symbol=?", coin).First(&coinEnt).Error
		if err != nil {
			log.Printf("Failed to find coin: %s", err.Error())
			return err
		}
		coins = append(coins, coinEnt)
	}

	for i := 0; i < 10; i++ {
		user := entity.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: encrypt.HashSHA256("a12345678"),
		}
		err := userRepo.CreateUser(&user)
		if err != nil {
			log.Printf("Failed to create user: %s", err.Error())
			return err
		}
		fmt.Printf("Created User: %v\n", user)

		//give users random amount crypto credentials
		for _, coin := range coins {
			cu := entity.NewCoinUser(coin.ID, user.ID)
			cu.SetAmount(r.Float64() * 100)
			err = userRepo.DepositCrypto(cu)
			if err != nil {
				log.Printf("Failed to create coin: %s", err.Error())
				return err
			}
		}
	}

	return nil
}

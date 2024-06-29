package main

import (
	"bbdk/app/api"
	"bbdk/cmd/seeder"
	coinGorm "bbdk/domain/repository/coin/gorm"
	coinRedis "bbdk/domain/repository/coin_price/redis"
	"bbdk/domain/service"
	"bbdk/infrastructure/godotenv"
	"bbdk/infrastructure/log/zerolog"
	"flag"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	seed := flag.Bool("seed", false, "Seed the database with initial data")

	flag.Parse()

	if *seed {
		seeder.Seed()
		return
	}
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	if err != nil {
		logger.Fatalf(err.Error())
	}
	coinRepository := coinGorm.NewCoinRepository(db)
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: "",
		DB:       0,
	})
	coinRedisRepo := coinRedis.NewCoinRepository(rdb)
	cp := service.NewCoinPrice(logger, env.PriceAPI, coinRepository, coinRedisRepo)

	go cp.Start() // start background updating price service

	api.Boot()
}

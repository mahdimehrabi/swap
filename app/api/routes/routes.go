package routes

import (
	coinRedis "bbdk/domain/repository/coin_price/redis"
	redisTran "bbdk/domain/repository/transaction/redis"
	gormUserRepo "bbdk/domain/repository/user/gorm"
	"bbdk/domain/service"
	"bbdk/infrastructure/godotenv"
	logger "bbdk/infrastructure/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Router interface {
	SetupRoutes(engine *gin.Engine)
}

func CreateRouters(env *godotenv.Env, logger logger.Logger) []Router {
	db, err := gorm.Open(postgres.Open(env.DATABASE_HOST), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to database error:%s", err.Error())
	}
	userRepo := gormUserRepo.NewUserRepository(db)
	userService := service.NewUserService(userRepo, logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: "",
		DB:       0,
	})
	tranRepo := redisTran.NewTransactionRepository(rdb)
	coinRedisRepo := coinRedis.NewCoinRepository(rdb)

	tranService := service.NewTransactionService(tranRepo, coinRedisRepo, logger)

	return []Router{NewUserRouter(userService), NewSwapRouter(tranService)}
}

func HandleRouters(e *gin.Engine, routers []Router) {
	for _, router := range routers {
		router.SetupRoutes(e)
	}
}

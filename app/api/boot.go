package api

import (
	"bbdk/app/api/routes"
	"bbdk/infrastructure/godotenv"
	"bbdk/infrastructure/log/zerolog"
	"github.com/gin-gonic/gin"
)

func Boot() {
	r := gin.Default()
	logger := zerolog.NewLogger()
	env := godotenv.NewEnv()
	env.Load()
	routes.HandleRouters(r, routes.CreateRouters(env, logger))
	if err := r.Run(env.ServerAddr); err != nil {
		logger.Fatalf("error running gin server error:%s", err.Error())
	}
}

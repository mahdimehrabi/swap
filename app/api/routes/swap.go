package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type SwapRouter struct {
	swapController controller.SwapController
}

func NewSwapRouter(tranService service.TransactionService) *SwapRouter {
	swapController := controller.NewSwapController(tranService)
	return &SwapRouter{swapController: *swapController} //transient controller injection to improve performance
}

func (rh *SwapRouter) SetupRoutes(router *gin.Engine) {
	//I must implement like /api/v1 but no time :(
	router.POST("/swap/commit/:id", rh.swapController.CommitTransaction)
	router.POST("/swap/:userID/:srcCoinID/:destCoinID", rh.swapController.CreateTransaction)

}

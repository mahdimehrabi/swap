package routes

import (
	controller "bbdk/app/api/controllers"
	"bbdk/domain/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userController controller.UserController
}

func NewUserRouter(userService service.UserService) *UserRouter {
	userController := controller.NewUserController(userService)
	return &UserRouter{userController: *userController} //transient controller injection to improve performance
}

func (rh *UserRouter) SetupRoutes(router *gin.Engine) {
	router.POST("/users", rh.userController.CreateUser)
	router.GET("/users/:id", rh.userController.GetUserByID)
	router.PUT("/users/:id", rh.userController.UpdateUser)
	router.DELETE("/users/:id", rh.userController.DeleteUser)
	router.GET("/users", rh.userController.GetAllUsers)

	//we must use authentication but its ok for test  purpose

}

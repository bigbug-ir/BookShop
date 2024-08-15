package router

import (
	controller "bmacharia/jwt-go-rbac/controllers"
	util "bmacharia/jwt-go-rbac/utils"

	"github.com/gin-gonic/gin"
)

func setUpSupportRoutes(supportRout *gin.RouterGroup) {
	// jwt role support route controller
	supportRout.Use(util.JWTAuthSupport())
	/*****************************************************************/
	// support users
	supportRout.GET("/users", controller.GetUsers)
	supportRout.GET("/users/:id", controller.GetUser)
	supportRout.PUT("/users/:id", controller.UpdateUserByAdmin)
	/*****************************************************************/
	//suport books route
	supportRout.GET("/books", controller.GetBooks)
	supportRout.GET("/books/:id", controller.GetBook)
	/*****************************************************************/
	//support order route
	supportRout.GET("/orders", controller.GetOrders)
	supportRout.GET("/orders/:id", controller.GetOrder)
	supportRout.PUT("/orders/:id", controller.UpdateOrder)
	/*****************************************************************/
	//support user profile route
	supportRout.GET("/", controller.Auth)
	supportRout.PUT("/", controller.UpdateUserAuth)
	supportRout.GET("/profile", controller.GetProfile)
	supportRout.POST("/profile", controller.AddProfile)
	supportRout.PUT("/profile", controller.UpdateProfile)
	supportRout.PUT("/password", controller.UpdatePassword)
	/*****************************************************************/
}

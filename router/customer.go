package router

import (
	controller "bmacharia/jwt-go-rbac/controllers"
	util "bmacharia/jwt-go-rbac/utils"

	"github.com/gin-gonic/gin"
)

func setUpCustomerRoutes(customerRout *gin.RouterGroup) {
	// jwt role customer route controller
	customerRout.Use(util.JWTAuthCustomer())
	/*****************************************************************/
	//customer orders route
	customerRout.GET("/orders", controller.GetAllOrderCustomerAuth)
	customerRout.POST("/orders", controller.AddOrder)
	customerRout.GET("/orders/:id", controller.GetOrderCustomerAuth)
	/*****************************************************************/
	//customer user profile route
	customerRout.GET("/", controller.Auth)
	customerRout.PUT("/", controller.UpdateUserAuth)
	customerRout.GET("/profile", controller.GetProfile)
	customerRout.POST("/profile", controller.AddProfile)
	customerRout.PUT("/profile", controller.UpdateProfile)
	customerRout.PUT("/password", controller.UpdatePassword)
	/*---------------------------------------------------------------*/
}

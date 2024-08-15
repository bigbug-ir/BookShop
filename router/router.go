package router

import (
	controller "bmacharia/jwt-go-rbac/controllers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	/*---------------------------------------------------------------*/
	app := gin.Default()
	/*****************************************************************/
	app.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.1"})
	/*---------------------------------------------------------------*/
	// auth routing
	authRoutes := app.Group("/auth/user")
	/*****************************************************************/
	//registration route
	authRoutes.POST("/register", controller.Register)
	/*****************************************************************/
	//login route
	authRoutes.POST("/login", controller.Login)
	/*---------------------------------------------------------------*/
	//only super admin rout
	superAdminRoute := app.Group("/super")
	/*****************************************************************/
	setUpSuperAdminRoutes(superAdminRoute)
	/*---------------------------------------------------------------*/
	// only admin route
	adminRoutes := app.Group("/admin")
	/*****************************************************************/
	// set up admin router
	setUpAdminRoutes(adminRoutes)
	/*---------------------------------------------------------------*/
	//only customer routes
	customerRout := app.Group("/customer")
	/*****************************************************************/
	//setup customer routs
	setUpCustomerRoutes(customerRout)
	/*---------------------------------------------------------------*/
	//supportRout
	supportRout := app.Group("/support")
	/*****************************************************************/
	// setup support routes
	setUpSupportRoutes(supportRout)
	/*****************************************************************/
	app.Run(os.Getenv("PORT"))
	fmt.Println("Server running on port :", os.Getenv("PORT"))
	/*****************************************************************/
}

package router

import (
	controller "bmacharia/jwt-go-rbac/controllers"
	util "bmacharia/jwt-go-rbac/utils"

	"github.com/gin-gonic/gin"
)

func setUpSuperAdminRoutes(superAdminRoute *gin.RouterGroup) {
	//jwt role super admin route controller
	superAdminRoute.Use(util.JWTAuth())
	/*****************************************************************/
	// super admin users roles route
	superAdminRoute.POST("/user/role", controller.CreateRole)
	superAdminRoute.GET("/user/roles", controller.GetRoles)
	superAdminRoute.GET("/user/role/:id", controller.GetRole)
	superAdminRoute.PUT("/user/role/:id", controller.UpdateRole)
	superAdminRoute.DELETE("/user/role/:id", controller.DeleteRole)
	/*****************************************************************/
	// super admin users route
	superAdminRoute.POST("/users", controller.CreateUser)
	superAdminRoute.GET("/users", controller.GetUsers)
	superAdminRoute.GET("/user/:id", controller.GetUser)
	superAdminRoute.PUT("/user/:id", controller.UpdateUser)
	superAdminRoute.DELETE("/user/:id", controller.DeleteUser)
	/*****************************************************************/
	// super admin orders route
	superAdminRoute.GET("/orders", controller.GetOrders)
	superAdminRoute.GET("/orders/:id", controller.GetOrder)
	superAdminRoute.PUT("/orders/:id", controller.UpdateOrder)
	superAdminRoute.DELETE("/orders/:id", controller.DeleteOrder)
	/*****************************************************************/
	// super admin user profile route
	superAdminRoute.GET("/", controller.Auth)
	superAdminRoute.PUT("/", controller.UpdateUserAuth)
	superAdminRoute.GET("/profile", controller.GetProfile)
	superAdminRoute.POST("/profile", controller.AddProfile)
	superAdminRoute.PUT("/profile", controller.UpdateProfile)
	superAdminRoute.PUT("/password", controller.UpdatePassword)
	/*****************************************************************/
	// super admin category route
	superAdminRoute.GET("/categories", controller.GetCategories)
	superAdminRoute.GET("/category/:id", controller.GetCategory)
	superAdminRoute.POST("/category", controller.AddCategory)
	superAdminRoute.PUT("/category/:id", controller.UpdateCategory)
	superAdminRoute.DELETE("/category/:id", controller.DeleteCategory)
	/*****************************************************************/
	// admin author route
	superAdminRoute.GET("/authors", controller.GetAuthors)
	superAdminRoute.GET("/author/:id", controller.GetAuthor)
	superAdminRoute.POST("/author", controller.AddAuthor)
	superAdminRoute.PUT("/author/:id", controller.UpdateAuthor)
	superAdminRoute.DELETE("/author/:id", controller.DeleteAuthor)
	/*****************************************************************/

}

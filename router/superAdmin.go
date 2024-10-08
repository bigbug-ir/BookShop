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
	superAdminRoute.POST("/user/roles", controller.CreateRole)
	superAdminRoute.GET("/user/roles", controller.GetRoles)
	superAdminRoute.GET("/user/roles/:id", controller.GetRole)
	superAdminRoute.PUT("/user/roles/:id", controller.UpdateRole)
	superAdminRoute.DELETE("/user/roles/:id", controller.DeleteRole)
	/*****************************************************************/
	// super admin users route
	superAdminRoute.POST("/users", controller.CreateUser)
	superAdminRoute.GET("/users", controller.GetUsers)
	superAdminRoute.GET("/users/:id", controller.GetUser)
	superAdminRoute.PUT("/users/:id", controller.UpdateUser)
	superAdminRoute.DELETE("/users/:id", controller.DeleteUser)
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
	superAdminRoute.GET("/category", controller.GetCategories)
	superAdminRoute.GET("/category/:id", controller.GetCategory)
	superAdminRoute.POST("/category", controller.AddCategory)
	superAdminRoute.PUT("/category/:id", controller.UpdateCategory)
	superAdminRoute.DELETE("/category/:id", controller.DeleteCategory)
	/*****************************************************************/
	// admin author route
	superAdminRoute.GET("/authors", controller.GetAuthors)
	superAdminRoute.GET("/authors/:id", controller.GetAuthor)
	superAdminRoute.POST("/authors", controller.AddAuthor)
	superAdminRoute.PUT("/authors/:id", controller.UpdateAuthor)
	superAdminRoute.DELETE("/authors/:id", controller.DeleteAuthor)
	/*****************************************************************/
	// admin books route
	superAdminRoute.GET("/books", controller.GetBooks)
	superAdminRoute.GET("/books/:id", controller.GetBook)
	superAdminRoute.POST("/books", controller.AddBook)
	superAdminRoute.PUT("/books/:id", controller.UpdateBook)
	superAdminRoute.DELETE("/books/:id", controller.DeleteBook)
	/*****************************************************************/

}

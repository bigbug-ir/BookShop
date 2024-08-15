package router

import (
	controller "bmacharia/jwt-go-rbac/controllers"
	util "bmacharia/jwt-go-rbac/utils"

	"github.com/gin-gonic/gin"
)

func setUpAdminRoutes(adminRoutes *gin.RouterGroup) {
	/*****************************************************************/
	// jwt role admin route controller
	adminRoutes.Use(util.JWTAuthAdmin())
	/*****************************************************************/
	// admin user route
	adminRoutes.POST("/users", controller.CreateUserByAdmin)
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/users/:id", controller.GetUser)
	adminRoutes.PUT("/users/:id", controller.UpdateUserByAdmin)
	/*****************************************************************/
	// admin books route
	adminRoutes.GET("/books", controller.GetBooks)
	adminRoutes.GET("/books/:id", controller.GetBook)
	adminRoutes.POST("/books", controller.AddBook)
	adminRoutes.PUT("/books/:id", controller.UpdateBook)
	adminRoutes.DELETE("/books/:id", controller.DeleteBook)
	/*****************************************************************/
	// admin order route
	adminRoutes.GET("/orders", controller.GetOrders)
	adminRoutes.GET("/orders/:id", controller.GetOrder)
	adminRoutes.PUT("/orders/:id", controller.UpdateOrder)
	adminRoutes.DELETE("/orders/:id", controller.DeleteOrder)
	/*****************************************************************/
	// admin user profile
	adminRoutes.GET("/", controller.Auth)
	adminRoutes.PUT("/", controller.UpdateUserAuth)
	adminRoutes.GET("/profile", controller.GetProfile)
	adminRoutes.POST("/profile", controller.AddProfile)
	adminRoutes.PUT("/profile", controller.UpdateProfile)
	adminRoutes.PUT("/password", controller.UpdatePassword)
	/*****************************************************************/
	// admin  category route
	adminRoutes.GET("/category", controller.GetCategories)
	adminRoutes.GET("/category/:id", controller.GetCategory)
	adminRoutes.POST("/category", controller.AddCategory)
	adminRoutes.PUT("/category/:id", controller.UpdateCategory)
	adminRoutes.DELETE("/category/:id", controller.DeleteCategory)
	/*****************************************************************/
	// admin author route
	adminRoutes.GET("/authors", controller.GetAuthors)
	adminRoutes.GET("/authors/:id", controller.GetAuthor)
	adminRoutes.POST("/authors", controller.AddAuthor)
	adminRoutes.PUT("/authors/:id", controller.UpdateAuthor)
	adminRoutes.DELETE("/authors/:id", controller.DeleteAuthor)
	/*****************************************************************/
}

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
	adminRoutes.POST("/users", controller.CreateUser)
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	/*****************************************************************/
	// admin books route
	adminRoutes.GET("/books", controller.GetBooks)
	adminRoutes.GET("/book/:id", controller.GetBook)
	adminRoutes.POST("/book", controller.AddBook)
	adminRoutes.PUT("/book/:id", controller.UpdateBook)
	adminRoutes.DELETE("/book/:id", controller.DeleteBook)
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
	adminRoutes.GET("/categories", controller.GetCategories)
	adminRoutes.GET("/category/:id", controller.GetCategory)
	adminRoutes.POST("/category", controller.AddCategory)
	adminRoutes.PUT("/category/:id", controller.UpdateCategory)
	adminRoutes.DELETE("/category/:id", controller.DeleteCategory)
	/*****************************************************************/
	// admin author route
	adminRoutes.GET("/authors", controller.GetAuthors)
	adminRoutes.GET("/author/:id", controller.GetAuthor)
	adminRoutes.POST("/author", controller.AddAuthor)
	adminRoutes.PUT("/author/:id", controller.UpdateAuthor)
	adminRoutes.DELETE("/author/:id", controller.DeleteAuthor)
	/*****************************************************************/
}

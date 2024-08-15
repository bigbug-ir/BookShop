package initializers

import (
	"bmacharia/jwt-go-rbac/database"
	model "bmacharia/jwt-go-rbac/models"
	"bmacharia/jwt-go-rbac/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*****************************************************************/

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

/*****************************************************************/

func LoadDatabase() {
	database.ConnectDB()
	log.Println("Database connected successfully")
	err := database.Database.DB.AutoMigrate(
		&model.Profile{},
		&model.Role{},
		&model.User{},
		&model.Profile{},
		&model.Book{},
		&model.Author{},
		&model.Order{},
		&model.OrderBook{},
		&model.Category{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	log.Println("Database migrated successfully")
}

/*****************************************************************/

func SetRole() {

	var roles = []model.Role{
		{Name: "superadmin", Description: "Super administrator role"},
		{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "support", Description: "Authenticated support role"},
		{Name: "anonymous", Description: "Unauthenticated customer role"}}
	database.Database.DB.Save(&roles)
	log.Println("Successfully added roles")

}

/*****************************************************************/

func SetSuperAdimin() {
	var user = []model.User{
		{Username: os.Getenv("SUPER_ADMIN_USERNAME"), Email: os.Getenv("SUPER_ADMIN_EMAIL"), Password: os.Getenv("SUPER_ADMIN_PASSWORD"), RoleID: 1}}
	database.Database.DB.Save(&user)
	log.Println("Successfully added  Admins")

}

/*****************************************************************/

func ServeRouter() {
	router.SetupRouter()
}

/*****************************************************************/

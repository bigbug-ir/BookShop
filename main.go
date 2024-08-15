package main

import (
	"bmacharia/jwt-go-rbac/initializers"
)

func main() {
	initializers.LoadEnv()
	initializers.LoadDatabase()
	initializers.SetRole()
	initializers.SetSuperAdimin()
	initializers.ServeRouter()
}

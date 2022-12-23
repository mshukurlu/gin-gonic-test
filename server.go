package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mshukurlu/gin-gonic-test/config"
	"github.com/mshukurlu/gin-gonic-test/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuhtController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}

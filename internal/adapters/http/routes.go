package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", RegisterUser) // From auth_handler.go
			auth.POST("/login", LoginUser)
		}

		user := api.Group("/user")
		user.Use(AuthMiddleware())
		{
			user.GET("/dashboard", UserDashboard) // From user_handler.go
		}

		txns := api.Group("/transactions")
		txns.Use(AuthMiddleware())
		{
			txns.POST("/transfer", TransferFunds) // From transaction_handler.go
			txns.GET("/history", TransactionHistory)
		}
	}
}

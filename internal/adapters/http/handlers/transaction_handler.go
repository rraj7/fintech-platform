package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransferFunds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Funds transferred"})
}

func TransactionHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"transactions": []string{}})
}

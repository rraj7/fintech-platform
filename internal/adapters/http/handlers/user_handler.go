package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"balance": 1000})
}

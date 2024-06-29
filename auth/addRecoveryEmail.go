package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecoveryEmailStr struct {
	Email string `json:"email"`
}

func RecoveryEmail(c *gin.Context) {
	var req RecoveryEmailStr
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", req)
}

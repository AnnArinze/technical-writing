package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordChange struct {
	Password string `json:"password"`
}

func ChangePassword(c *gin.Context) {
	var req PasswordChange
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", req)
}

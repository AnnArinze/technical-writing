package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPassword struct {
	Email string `json:"email"`
}

func ForgotPasswordHandler(c *gin.Context) {
	var req Register
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", req)
}

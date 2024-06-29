package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangeUserName struct {
	Password string `json:"password"`
}

func ChangeUsername(c *gin.Context) {
	var req ChangeUserName
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", req)
}

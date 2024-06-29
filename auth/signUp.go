package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ID        string `json:"id"`
}

func SignUp(c *gin.Context) {
	var req Register
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("%v", req)
}

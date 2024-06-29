package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginStr struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginStr
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if c.Request.Header.Get("x-seoiw-token-sdw") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are so dumb",
		})
		return
	}

	fmt.Println(c.Request.Header, ">>>>>>>>>>>>>>>>")

	fmt.Printf("%v", req)
	c.JSON(http.StatusOK, req)
}

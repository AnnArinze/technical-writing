package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	tony := c.Param("tony")
	zara := c.Param("zara")
	fmt.Printf("somethignboring %s and %s \n", tony, zara)
	fmt.Printf("%v", id)
}

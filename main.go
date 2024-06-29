package main

import (
	"github.com/ann-arinze/technical-writing-001/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authGrp := r.Group("/auth")
	authGrp.POST("/register", auth.SignUp)
	authGrp.POST("/login", auth.Login)
	authGrp.POST("/forgotpassword", auth.ForgotPasswordHandler)
	authGrp.POST("/resetpassword", auth.ChangePassword)
	authGrp.POST("/changeusername", auth.ChangeUsername)
	authGrp.POST("/recoveryemail", auth.RecoveryEmail)
	authGrp.DELETE("/delete/:id/:tonie/:zara", auth.DeleteUser)

	r.Run(":8081") //
}

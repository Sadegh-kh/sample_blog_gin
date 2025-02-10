package web

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {

	r.GET("", index)

	r.GET("/login", login)
	r.GET("/register", register)
}

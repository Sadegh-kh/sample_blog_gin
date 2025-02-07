package web

import "github.com/gin-gonic/gin"

func Register(g gin.RouterGroup) {

	g.GET("", index)

	g.GET("/login", login)
	g.GET("/register", register)
}

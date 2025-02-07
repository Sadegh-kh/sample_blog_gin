package web

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func register(c *gin.Context) {
	c.HTML(200, "register.html", nil)

}

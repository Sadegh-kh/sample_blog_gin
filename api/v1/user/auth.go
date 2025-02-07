package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "login",
	})

}

func register(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "register"})
}

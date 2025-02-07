package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"name": "blog"})
}

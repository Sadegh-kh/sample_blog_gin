package api

import (
	v1 "blog/api/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	g := r.Group("/api")
	v1.Register(g)

}

package v1

import (
	"blog/api/v1/user"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	g := r.Group("/v1")
	user.Register(g)
}

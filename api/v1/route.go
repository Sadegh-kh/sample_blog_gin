package v1

import (
	"github.com/gin-gonic/gin"
)

func (h V1Handller) Register(r *gin.RouterGroup) {
	g := r.Group("/v1")
	h.UserHandller.Register(g)
}

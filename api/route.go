package api

import (
	"github.com/gin-gonic/gin"
)

func (h ApiHandller) Register(r *gin.Engine) {

	g := r.Group("/api")
	h.V1Handller.Register(g)

}

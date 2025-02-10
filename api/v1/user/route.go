package user

import "github.com/gin-gonic/gin"

func (uh UserHandller) Register(r *gin.RouterGroup) {

	r.POST("register", uh.register)
	r.POST("login", uh.login)

}

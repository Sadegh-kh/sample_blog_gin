package user

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {

	r.POST("register", register)
	r.POST("login", login)

}

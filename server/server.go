package server

import (
	"blog/api"
	"blog/web"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s Server) Run() {

	r := gin.Default()

	templates, _ := filepath.Glob("templates/**/*.html")
	templates = append(templates, "templates/index.html")
	r.LoadHTMLFiles(templates...)

	web.Register(r)
	s.ApiHanddler.Register(r)

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"messgae": "ok"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

type Server struct {
	ApiHanddler api.ApiHandller
}

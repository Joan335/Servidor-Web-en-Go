package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine)  {
	
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola, Mundo")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola, %s", nombre)
	})

	r.Run(":8080")

}
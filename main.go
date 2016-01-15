package main

import (
	"net/http"

	"github.com/Depado/jagon-api/jagon"
	"github.com/gin-gonic/gin"
)

const version = "1"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/static", "./assets")
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	apir := r.Group("/api/v" + version)
	{
		apir.POST("/pray", jagon.Pray)
		apir.POST("/praise", jagon.Praise)
		apir.POST("/confess", jagon.Confess)
		apir.POST("/sacrifice", jagon.DoSacrifice)
		apir.GET("/info/prophet", jagon.Prophet)
		apir.GET("/info/apostles", jagon.Apostotles)
	}
	r.Run(":8080")
}

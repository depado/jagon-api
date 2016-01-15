package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Depado/jagon-api/conf"
	"github.com/Depado/jagon-api/jagon"
	"github.com/gin-gonic/gin"
)

const version = "1"

func main() {
	var err error

	if err = conf.Load("conf.yml"); err != nil {
		log.Fatal(err)
	}

	if !conf.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
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
	r.Run(":" + strconv.Itoa(conf.C.Port))
}

package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kelvintaywl/ramenipsum/lib"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func paragraphHandler(c *gin.Context) {
	para := c.Param("num")
	locale := c.DefaultQuery("lc", "en_UK")
	n, err := strconv.Atoi(para)
	if err != nil {
		n = 1
	}
	c.String(http.StatusOK, lib.RandomText(n, locale))
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl",
		gin.H{
			"title": "Ramen Ipsum: A brothier Lorem Ipsum generator",
		})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/assets", "assets")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)
	router.GET("/paragraphs/:num", paragraphHandler)

	router.Run(":" + port)
}

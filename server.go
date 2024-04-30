package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.LoadHTMLGlob("pages/*.html")
	router.Static("images", "./images")
	router.Static("fonts", "./fonts")

	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "arcanist.html", "")
	})

	router.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "dragonknight.html", "")
	})

	router.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "necro.html", "")
	})

	router.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "nightblade.html", "")
	})

	router.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "sorc.html", "")
	})

	router.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "templar.html", "")
	})

	router.GET("warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "warden.html", "")
	})

	router.Run();
}
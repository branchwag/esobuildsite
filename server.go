package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("images", "./images")
	r.Static("fonts", "./fonts")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "ESO Builds",
			"Content": "index.html",
		})
	})

	r.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Arcanist",
			"Content": "arcanist.html",
		})
	})

	r.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Dragonknight",
			"Content": "dragonknight.html",
		})
	})

	r.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Necromancer",
			"Content": "necromancer.html",
		})
	})

	r.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Nightblade",
			"Content": "nightblade.html",
		})
	})

	r.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Sorcerer",
			"Content": "sorcerer.html",
		})
	})

	r.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Templar",
			"Content": "templar.html",
		})
	})

	r.GET("/warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Warden",
			"Content": "warden.html",
		})
	})

	r.Run();
}

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
			"Content": "index",
		})
	})

	r.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Arcanist",
			"Content": "arcanist",
		})
	})

	r.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Dragonknight",
			"Content": "dragonknight",
		})
	})

	r.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Necromancer",
			"Content": "necromancer",
		})
	})

	r.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Nightblade",
			"Content": "nightblade",
		})
	})

	r.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Sorcerer",
			"Content": "sorcerer",
		})
	})

	r.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Templar",
			"Content": "templar",
		})
	})

	r.GET("/warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Warden",
			"Content": "warden",
		})
	})

	r.Run();
}

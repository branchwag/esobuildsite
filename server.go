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
			"Title":   "Page",
			"Content": "index.html",
		})
	})

	r.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "arcanist.html",
		})
	})

	r.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "dragonknight.html",
		})
	})

	r.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "necromancer.html",
		})
	})

	r.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "nightblade.html",
		})
	})

	r.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "sorcerer.html",
		})
	})

	r.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "templar.html",
		})
	})

	r.GET("/warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "Page",
			"Content": "warden.html",
		})
	})

	r.Run();
}

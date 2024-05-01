package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {

	buildSections := []string{"cp", "gear", "food", "skills", "race"}

	router := gin.Default()

	router.LoadHTMLGlob("pages/*.html")
	router.Static("images", "./images")
	router.Static("fonts", "./fonts")

	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "arcanist.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "dragonknight.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "necro.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "nightblade.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "sorc.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "templar.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.GET("warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "warden.html", gin.H{
            "BuildSections": buildSections,
        })
	})

	router.Run();
}
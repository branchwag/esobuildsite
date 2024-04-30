package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.LoadHTMLGlob("pages/*.html")

	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.Run();
}
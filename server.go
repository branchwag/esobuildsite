package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()
	r.SetHTMLTemplate(loadTemplates("templates"))

	r.Static("images", "./images")
	r.Static("fonts", "./fonts")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.GET("/arcanist", func(c *gin.Context){
		c.HTML(http.StatusOK, "arcanist.html", "")
	})

	r.GET("/dragonknight", func(c *gin.Context){
		c.HTML(http.StatusOK, "dragonknight.html", "")
	})

	r.GET("/necromancer", func(c *gin.Context){
		c.HTML(http.StatusOK, "necro.html", "")
	})

	r.GET("/nightblade", func(c *gin.Context){
		c.HTML(http.StatusOK, "nightblade.html", "")
	})

	r.GET("/sorcerer", func(c *gin.Context){
		c.HTML(http.StatusOK, "sorc.html", "")
	})

	r.GET("/templar", func(c *gin.Context){
		c.HTML(http.StatusOK, "templar.html", "")
	})

	r.GET("warden", func(c *gin.Context){
		c.HTML(http.StatusOK, "warden.html", "")
	})

	r.Run();
}

func loadTemplates(templatesDir string) *template.Template {
    templ := template.New("")

    filepath.WalkDir(templatesDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !d.IsDir() && filepath.Ext(path) == ".html" {
            content, err := os.ReadFile(path)
            if err != nil {
                return err
            }
			name := filepath.Base(path) // Get the base file name
            _, err = templ.New(name).Parse(string(content)) // Use the base file name as the template name
            if err != nil {
                return err
            }
        }
        return nil
    })

    return templ
}
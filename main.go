package main

import (
	"net/http"
	"strings"

	"strconv"

	"github.com/dkallman13/Vicky3Optimizer/initial"
	"github.com/dkallman13/Vicky3Optimizer/types"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.GetEnvVars()
	initial.SaveFileLocSetter()
	initial.ConnectToDB()
	initial.SaveFileLister()
	initial.Migrate()
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base/base.html", "templates/index.html")
	r.AddFromFiles("/save", "templates/base/base.html", "templates/save/save.html")
	r.AddFromFiles("/db/state", "templates/base/base.html", "templates/db/state.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createRenderer()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title":     "home",
			"savefiles": initial.SaveFiles,
		})
	})
	for _, savFile := range initial.SaveFiles {
		var savRouteBuilder strings.Builder
		savRouteBuilder.WriteString("save/")
		savRouteBuilder.WriteString(savFile)
		router.GET(savRouteBuilder.String(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "/save", gin.H{
				"title":     savFile,
				"firstline": initial.TokenizeSave(savFile),
			})
		})
	}
	router.GET("/db/state", func(c *gin.Context) {
		var results []types.State
		stateArray := initial.DB.Table("states").Select("id", "name").Find(&results)
		print(stateArray)
		c.HTML(http.StatusOK, "/db/state", gin.H{
			"title": "states",
			"states": stateArray,
		})
	})
	router.POST("/db/state", func(c *gin.Context) {
		c.Request.ParseForm()
		idstring :=c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err!=nil {
			panic(err)
		}
		newstate := types.NewState(id, c.Request.Form.Get("stateName"))
		initial.DB.Create(&newstate)
	})
	router.Run() // listen and serve on localhost
}

package main

import (
	"net/http"
	"strconv"
	"strings"

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
		initial.DB.Table("states").Select("*").Find(&results)
		c.HTML(http.StatusOK, "/db/state", gin.H{
			"title":  "states",
			"states": results,
		})
	})
	router.POST("/db/state", func(c *gin.Context) {
		c.Request.ParseForm()
		idstring := c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		newstate := types.NewState(id, c.Request.Form.Get("stateName"))
		initial.DB.Omit("Provinces.*").Create(&newstate)
	})
	router.POST("/db/stateU", func(c *gin.Context) {
		c.Request.ParseForm()
		idstring := c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		var state types.State
		initial.DB.Table("states").Select("Id", "Name", "Provinces").Where(&types.State{Id: id}).Find(&state)
		if c.Request.Form.Get("ProvinceIds") != "" {
			provIds := strings.Split(c.Request.Form.Get("ProvinceIds"), ",")
			for i := 0; i < len(provIds); i++ {
				provId, err := strconv.Atoi(provIds[i])
				if err != nil {
					panic(err)
				}
				var province *types.Province
				initial.DB.First(&province, provId)
				if province != nil {
					initial.DB.Model(&state).Association("Provinces").Append(&province)
				}
			}
		}
		initial.DB.Model(&state).Update("Name", c.Request.Form.Get("stateName"))
	})
	router.Run() // listen and serve on localhost
}

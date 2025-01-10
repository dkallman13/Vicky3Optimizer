package main

import (
	"strconv"
	"strings"
	"github.com/dkallman13/Vicky3Optimizer/controllers"
	"github.com/dkallman13/Vicky3Optimizer/model"
	"github.com/dkallman13/Vicky3Optimizer/types"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func init() {
	model.GetEnvVars()
	model.SaveFileLocSetter()
	model.ConnectToDB()
	model.SaveFileLister()
	model.Migrate()
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

	router.GET("/", controllers.Home)
	controllers.Save(router)
	router.GET("/db/state", controllers.StateGet)
	router.POST("/db/state", func(c *gin.Context) {
		c.Request.ParseForm()
		idstring := c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		newstate := types.NewState(id, c.Request.Form.Get("stateName"))
		model.DB.Omit("Provinces").Create(&newstate)
	})
	router.POST("/db/stateU", func(c *gin.Context) {
		c.Request.ParseForm()
		idstring := c.Request.Form.Get("stateId")
		id, err := strconv.Atoi(idstring)
		if err != nil {
			panic(err)
		}
		var state types.State
		model.DB.Table("states").Select("Id", "Name").Where(&types.State{Id: id}).Find(&state)
		var provinces []*types.Province
		if c.Request.Form.Get("ProvinceIds") != "" {
			provIds := strings.Split(c.Request.Form.Get("ProvinceIds"), ",")
			for i := 0; i < len(provIds); i++ {
				provId, err := strconv.Atoi(provIds[i])
				if err != nil {
					panic(err)
				}
				var province *types.Province
				model.DB.Model(&province).Where(&types.Province{Id: provId}).Find(&province)
				if province != nil {
					provinces = append(provinces, province)
				}
			}
			model.DB.Model(&state).Association("Province").Replace(provinces)
		}
		if c.Request.Form.Get("airableLand") != "" {
			airableland := c.Request.Form.Get("airableLand")
			land, err := strconv.Atoi(airableland)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("AirableLand", land)
		}
		if c.Request.Form.Get("Iron") != "" {
			ironstr := c.Request.Form.Get("Iron")
			iron, err := strconv.Atoi(ironstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("IronCap", iron)
		}
		if c.Request.Form.Get("Coal") != "" {
			coalstr := c.Request.Form.Get("Coal")
			coal, err := strconv.Atoi(coalstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("CoalCap", coal)
		}
		if c.Request.Form.Get("Sulfur") != "" {
			sulfurstr := c.Request.Form.Get("Sulfur")
			sulfur, err := strconv.Atoi(sulfurstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("SulfurCap", sulfur)
		}
		if c.Request.Form.Get("Lead") != "" {
			leadstr := c.Request.Form.Get("Lead")
			lead, err := strconv.Atoi(leadstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("LeadCap", lead)
		}
		if c.Request.Form.Get("Wood") != "" {
			woodstr := c.Request.Form.Get("Wood")
			wood, err := strconv.Atoi(woodstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("WoodCap", wood)
		}
		if c.Request.Form.Get("Fish") != "" {
			fishstr := c.Request.Form.Get("Fish")
			fish, err := strconv.Atoi(fishstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("FishCap", fish)
		}
		if c.Request.Form.Get("Whales") != "" {
			whalestr := c.Request.Form.Get("Whales")
			whale, err := strconv.Atoi(whalestr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("WhaleCap", whale)
		}
		if c.Request.Form.Get("Oil") != "" {
			oilstr := c.Request.Form.Get("Oil")
			oil, err := strconv.Atoi(oilstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("OilCap", oil)
		}
		if c.Request.Form.Get("Rubber") != "" {
			rubberstr := c.Request.Form.Get("Rubber")
			rubber, err := strconv.Atoi(rubberstr)
			if err != nil {
				panic(err)
			}
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("RubberCap", rubber)
		}
		if c.Request.Form.Get("stateName") != ""{
			model.DB.Model(&state).Where(&types.State{Id: id}).Update("Name", c.Request.Form.Get("stateName"))
		}
	})
	router.Run() // listen and serve on localhost
}

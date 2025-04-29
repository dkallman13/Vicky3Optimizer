package controllers
import(
	"net/http"
	"github.com/gin-gonic/gin"
)
func StateFile(c* gin.Context) {
	c.HTML(http.StatusOK, "/db/statefile", gin.H{
		"title": "State File Upload",
	})
}
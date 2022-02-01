package ratings

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/mdgkb-server/handlers/ratings"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
	//r.GET("/current", h.GetAll)
	//r.POST("/current-full", h.GetAll)
	//r.GET("/:id", h.Get)
	//r.GET("/create-slugs", h.CreateSlugs)
	//r.GET("/search", h.Search)
	//r.GET("/:slug", h.Get)
	//r.GET("/division/:id", h.GetByDivisionID)
	//r.POST("", h.Create)
	//r.DELETE("/:id", h.Delete)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/comment/:id", h.RemoveComment)
	//r.PUT("/comment/:id", h.UpdateComment)
	//r.POST("/comment", h.CreateComment)
}

package routing

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/handlers/auth"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/ratings"
	"mdgkb/mdgkb-server/helpers"
	authRouter "mdgkb/mdgkb-server/routing/auth"
	eventsRouter "mdgkb/mdgkb-server/routing/events"
	ratingsRouter "mdgkb/mdgkb-server/routing/ratings"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	//localUploader := helpers.NewLocalUploader(&config.UploadPath)
	//localUploaderNew := uploadHelper.NewLocalUploader(&config.UploadPath)
	helper := helpers.NewHelper(config)

	r.Static("/static", "./static/")

	api := r.Group("/api/voting")
	//m := middleware.CreateMiddleware(helper)
	//api.Use(m.Authentication())

	authRouter.Init(api.Group("/auth"), auth.CreateHandler(db, helper))
	eventsRouter.Init(api.Group("/events"), events.CreateHandler(db, helper))
	ratingsRouter.Init(api.Group("/ratings"), ratings.CreateHandler(db, helper))
}

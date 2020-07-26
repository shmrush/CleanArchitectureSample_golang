package infrastructure

import (
	"CleanArchitectureSample_golang/common"
	"CleanArchitectureSample_golang/infrastructure/middleware"
	"CleanArchitectureSample_golang/infrastructure/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// NewRouter initializes gin.Engine
func NewRouter(c *dig.Container) *gin.Engine {
	r := gin.Default()

	setupSession(r)

	apiRouter := r.Group("/api", middleware.CSRFProtection)
	authorizedRouter := r.Group("/api", middleware.RequireAuth, middleware.CSRFProtection)
	router.AuthRouter(apiRouter, c)
	router.TodoRouter(authorizedRouter, c)

	apiRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "root"})
	})
	return r
}

func setupSession(r *gin.Engine) {
	kvsEnv := common.KvsEnv
	store, _ := redis.NewStore(10, kvsEnv.Protocol, kvsEnv.Host+":"+kvsEnv.Port, "", []byte("session"))
	r.Use(sessions.Sessions("_clean_architecture_sample_golang_session", store))
}

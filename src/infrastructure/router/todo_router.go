package router

import (
	"CleanArchitectureSample_golang/interfaces/controller"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// TodoRouter is todos router.
func TodoRouter(apiRouter *gin.RouterGroup, container *dig.Container) {
	apiRouter.GET("/todos", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		err := container.Invoke(func(controller controller.TodosController) {
			controller.Index(ctx, session)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	apiRouter.GET("/todos/:id", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		err := container.Invoke(func(controller controller.TodosController) {
			controller.Show(ctx, session)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	apiRouter.POST("/todos", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		err := container.Invoke(func(controller controller.TodosController) {
			controller.Create(ctx, session)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})
}

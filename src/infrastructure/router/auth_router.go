package router

import (
	"CleanArchitectureSample_golang/interfaces/controller"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// AuthRouter is sign in and sign out router.
func AuthRouter(apiRouter *gin.RouterGroup, container *dig.Container) {
	apiRouter.POST("/auth/sign_up", func(ctx *gin.Context) {
		err := container.Invoke(func(controller controller.UsersController) {
			controller.SignUp(ctx)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	apiRouter.POST("/auth/sign_in", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		err := container.Invoke(func(controller controller.UsersController) {
			controller.SignIn(ctx, session)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	apiRouter.DELETE("/auth/sign_out", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		err := container.Invoke(func(controller controller.UsersController) {
			controller.SignOut(ctx, session)
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	})
}

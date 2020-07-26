package di

import (
	"CleanArchitectureSample_golang/interfaces/controller"
	"CleanArchitectureSample_golang/interfaces/database/dao"
	"CleanArchitectureSample_golang/interfaces/database/repository"

	"go.uber.org/dig"
)

// InterfacesModule injects interfaces dependencies.
func InterfacesModule(c *dig.Container) {
	// DAO
	c.Provide(dao.NewUserDao)
	c.Provide(dao.NewTodoDao)

	// Repository
	c.Provide(repository.NewUserDaoRepository)
	c.Provide(repository.NewTodoDaoRepository)

	// Controller
	c.Provide(controller.NewUsersController)
	c.Provide(controller.NewTodosController)
}

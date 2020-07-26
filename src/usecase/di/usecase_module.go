package di

import (
	"CleanArchitectureSample_golang/usecase"

	"go.uber.org/dig"
)

// UseCaseModule injects usecase dependencies.
func UseCaseModule(c *dig.Container) {
	c.Provide(usecase.NewSignUpUseCase)
	c.Provide(usecase.NewSignInUseCase)
	c.Provide(usecase.NewGetOwnTodoUseCase)
	c.Provide(usecase.NewGetOwnTodosUseCase)
	c.Provide(usecase.NewCreateTodoUseCase)
}

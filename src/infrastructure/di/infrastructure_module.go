package di

import (
	"CleanArchitectureSample_golang/infrastructure"

	"go.uber.org/dig"
)

// InfrastructureModule injects interfaces dependencies.
func InfrastructureModule(c *dig.Container) {
	c.Provide(infrastructure.NewDatabase)
}

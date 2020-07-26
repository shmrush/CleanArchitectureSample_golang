package main

import (
	"CleanArchitectureSample_golang/infrastructure"
	infrastructureDI "CleanArchitectureSample_golang/infrastructure/di"
	interfacesDI "CleanArchitectureSample_golang/interfaces/di"
	usecaseDI "CleanArchitectureSample_golang/usecase/di"

	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	inject(c)
	infrastructure.NewRouter(c).Run(":8000")
	c.Invoke(func(database infrastructure.Database) {
		defer database.Conn.Close()
	})
}

func inject(c *dig.Container) {
	infrastructureDI.InfrastructureModule(c)
	interfacesDI.InterfacesModule(c)
	usecaseDI.UseCaseModule(c)
}

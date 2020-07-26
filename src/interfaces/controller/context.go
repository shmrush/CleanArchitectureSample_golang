package controller

// Context is interface of gin.Context
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}

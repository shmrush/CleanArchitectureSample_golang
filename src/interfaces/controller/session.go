package controller

// Session is interface of gin.Sessions.session
type Session interface {
	Save() error
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Delete(key interface{})
}

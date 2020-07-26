package repository

import "CleanArchitectureSample_golang/domain/model"

// UserDaoRepository is interface of UserDaoRepositoryImpl
type UserDaoRepository interface {
	Create(name string, email string, password string) (model.User, error)
	Get(id uint) (model.User, error)
	FindByEmailAndPassword(email string, password string) (model.User, error)
}

package repository

import (
	"CleanArchitectureSample_golang/domain/model"
	"time"
)

// TodoDaoRepository is interface of TodoDaoRepositoryImpl
type TodoDaoRepository interface {
	FindByIDAndUserID(id uint, userID uint) (model.Todo, error)
	Create(userID uint, title string, description string, deadlineAt time.Time) (model.Todo, error)
	FindByUserID(userID uint) ([]model.Todo, error)
	UpdateCompleted(id uint, completed bool) (model.Todo, error)
}

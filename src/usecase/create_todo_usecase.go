package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"time"
)

// CreateTodoUseCase is interface of CreateTodoUseCaseImpl.
type CreateTodoUseCase interface {
	Execute(userID uint, title string, description string, deadlineAt time.Time) (model.Todo, error)
}

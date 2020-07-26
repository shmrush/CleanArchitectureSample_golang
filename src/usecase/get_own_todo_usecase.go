package usecase

import "CleanArchitectureSample_golang/domain/model"

// GetOwnTodoUseCase is interface of GetOwnTodoUseCaseImpl.
type GetOwnTodoUseCase interface {
	Execute(id uint, userID uint) (model.Todo, error)
}

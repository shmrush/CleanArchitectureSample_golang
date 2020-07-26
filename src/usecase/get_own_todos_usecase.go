package usecase

import "CleanArchitectureSample_golang/domain/model"

// GetOwnTodosUseCase is interface of GetOwnTodosUseCaseImpl.
type GetOwnTodosUseCase interface {
	Execute(userID uint) ([]model.Todo, error)
}

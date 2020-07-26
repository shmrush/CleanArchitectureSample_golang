package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
)

// GetOwnTodosUseCaseImpl provides to get own todos.
type GetOwnTodosUseCaseImpl struct {
	todoDaoRepository repository.TodoDaoRepository
}

// NewGetOwnTodosUseCase initializes GetOwnTodosUseCaseImpl.
func NewGetOwnTodosUseCase(todoDaoRepository repository.TodoDaoRepository) GetOwnTodosUseCase {
	return &GetOwnTodosUseCaseImpl{todoDaoRepository}
}

// Execute gets own todos.
func (r *GetOwnTodosUseCaseImpl) Execute(userID uint) (todos []model.Todo, err error) {
	todos, err = r.todoDaoRepository.FindByUserID(userID)
	return
}

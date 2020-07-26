package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
)

// GetOwnTodoUseCaseImpl provides to get own todo.
type GetOwnTodoUseCaseImpl struct {
	todoDaoRepository repository.TodoDaoRepository
}

// NewGetOwnTodoUseCase initializes GetOwnTodoUseCaseImpl.
func NewGetOwnTodoUseCase(todoDaoRepository repository.TodoDaoRepository) GetOwnTodoUseCase {
	return &GetOwnTodoUseCaseImpl{todoDaoRepository}
}

// Execute gets todo by id and userID.
func (r *GetOwnTodoUseCaseImpl) Execute(id uint, userID uint) (todo model.Todo, err error) {
	todo, err = r.todoDaoRepository.FindByIDAndUserID(id, userID)
	return
}

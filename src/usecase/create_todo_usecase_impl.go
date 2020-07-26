package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
	"time"
)

// CreateTodoUseCaseImpl provides to create todo.
type CreateTodoUseCaseImpl struct {
	todoDaoRepository repository.TodoDaoRepository
}

// NewCreateTodoUseCase initializes CreateTodoUseCaseImpl.
func NewCreateTodoUseCase(todoDaoRepository repository.TodoDaoRepository) CreateTodoUseCase {
	return &CreateTodoUseCaseImpl{todoDaoRepository}
}

// Execute creates todo.
func (u *CreateTodoUseCaseImpl) Execute(userID uint, title string, description string, deadlineAt time.Time) (todo model.Todo, err error) {
	todo, err = u.todoDaoRepository.Create(userID, title, description, deadlineAt)
	return
}

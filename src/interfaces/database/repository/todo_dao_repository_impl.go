package repository

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
	"CleanArchitectureSample_golang/interfaces/database/dao"
	"time"
)

// TodoDaoRepositoryImpl is TodoDao interface.
type TodoDaoRepositoryImpl struct {
	todoDao dao.TodoDao
}

// NewTodoDaoRepository initializes TodoDaoRepositoryImpl.
func NewTodoDaoRepository(todoDao dao.TodoDao) repository.TodoDaoRepository {
	return &TodoDaoRepositoryImpl{todoDao}
}

// FindByIDAndUserID finds todo by id and userID.
func (r *TodoDaoRepositoryImpl) FindByIDAndUserID(id uint, userID uint) (todo model.Todo, err error) {
	e, err := r.todoDao.FindByIDAndUserID(id, userID)
	if err != nil {
		return
	}
	todo = model.Todo{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		DeadlineAt:  e.DeadlineAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
	return
}

// Create creates new todo.
func (r *TodoDaoRepositoryImpl) Create(userID uint, title string, description string, deadlineAt time.Time) (todo model.Todo, err error) {
	e, err := r.todoDao.Create(userID, title, description, deadlineAt)
	if err != nil {
		return
	}
	todo = toModel(e)
	return
}

// FindByUserID finds todos by userID.
func (r *TodoDaoRepositoryImpl) FindByUserID(userID uint) (todos []model.Todo, err error) {
	todos = []model.Todo{}
	e, err := r.todoDao.FindByUserID(userID)
	if err != nil {
		return
	}
	for _, todo := range e {
		model := toModel(todo)
		todos = append(todos, model)
	}
	return
}

// UpdateCompleted updates completed.
func (r *TodoDaoRepositoryImpl) UpdateCompleted(id uint, completed bool) (todo model.Todo, err error) {
	e, err := r.todoDao.UpdateCompleted(id, completed)
	if err != nil {
		return
	}
	todo = toModel(e)
	return
}

func toModel(e dao.TodoEntity) (m model.Todo) {
	m = model.Todo{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		DeadlineAt:  e.DeadlineAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
	return
}

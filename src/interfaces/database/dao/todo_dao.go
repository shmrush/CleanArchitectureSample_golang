package dao

import (
	"CleanArchitectureSample_golang/infrastructure"
	"time"

	"github.com/jinzhu/gorm"
)

// TodoEntity is POGO of todos table.
type TodoEntity struct {
	gorm.Model
	UserID      uint
	Title       string
	Description string
	DeadlineAt  time.Time
	Completed   bool
}

// TodoDao provides todos table manipulation.
type TodoDao struct {
	database infrastructure.Database
}

// NewTodoDao initializes TodoDao.
func NewTodoDao(database infrastructure.Database) TodoDao {
	return TodoDao{database}
}

// TableName sets custom table name.
func (TodoEntity) TableName() string {
	return "todos"
}

// FindByIDAndUserID finds todo by id and user_id.
func (d *TodoDao) FindByIDAndUserID(id uint, userID uint) (todo TodoEntity, err error) {
	todo = TodoEntity{}
	if result := d.database.Conn.Where("id = ? AND user_id = ?", id, userID).First(&todo); result.Error != nil {
		err = result.Error
	}
	return
}

// Create creates new todo.
func (d *TodoDao) Create(userID uint, title string, description string, deadlineAt time.Time) (todo TodoEntity, err error) {
	todo = TodoEntity{UserID: userID, Title: title, Description: description, DeadlineAt: deadlineAt, Completed: false}
	if result := d.database.Conn.Create(&todo); result.Error != nil {
		err = result.Error
	}
	return
}

// FindByUserID finds by user_id.
func (d *TodoDao) FindByUserID(userID uint) (todos []TodoEntity, err error) {
	todos = []TodoEntity{}
	if result := d.database.Conn.Where("user_id = ?", userID).Find(&todos); result.Error != nil {
		err = result.Error
	}
	return
}

// UpdateCompleted updates completed column.
func (d *TodoDao) UpdateCompleted(id uint, completed bool) (todo TodoEntity, err error) {
	todo = TodoEntity{}
	if result := d.database.Conn.Find(&todo, id).Update("completed", completed); result.Error != nil {
		err = result.Error
	}
	return
}

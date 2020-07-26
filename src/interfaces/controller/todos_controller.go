package controller

import (
	"CleanArchitectureSample_golang/common"
	"CleanArchitectureSample_golang/usecase"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TodosController provides todo manipulation controller.
type TodosController struct {
	getOwnTodoUseCase  usecase.GetOwnTodoUseCase
	getOwnTodosUseCase usecase.GetOwnTodosUseCase
	createTodoUseCase  usecase.CreateTodoUseCase
}

type showParams struct {
	ID uint `json:"id" binding:"required"`
}

type createParams struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DeadlineAt  time.Time `json:"deadlineAt" binding:"required"`
}

// NewTodosController initializes TodosController.
func NewTodosController(
	getOwnTodoUseCase usecase.GetOwnTodoUseCase,
	getOwnTodosUseCase usecase.GetOwnTodosUseCase,
	createTodoUseCase usecase.CreateTodoUseCase,
) TodosController {
	return TodosController{getOwnTodoUseCase, getOwnTodosUseCase, createTodoUseCase}
}

// Index gets own todos.
func (c *TodosController) Index(ctx Context, session Session) {
	userID := session.Get(common.SessionUserIDKey).(uint)
	todos, err := c.getOwnTodosUseCase.Execute(userID)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"status": "Internal server error", "message": err.Error()})
		return
	}
	ctx.JSON(200, todos)
}

// Show gets own todo.
func (c *TodosController) Show(ctx Context, session Session) {
	userID := session.Get(common.SessionUserIDKey).(uint)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"status": "Bad request", "message": err.Error()})
		return
	}
	todo, err := c.getOwnTodoUseCase.Execute(uint(id), userID)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"status": "Internal server error", "message": err.Error()})
		return
	}
	ctx.JSON(200, todo)
}

// Create creates todo.
func (c *TodosController) Create(ctx Context, session Session) {
	userID := session.Get(common.SessionUserIDKey).(uint)
	p := createParams{}
	if err := ctx.Bind(&p); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"status": "Bad request", "message": err.Error()})
		return
	}
	todo, err := c.createTodoUseCase.Execute(userID, p.Title, p.Description, p.DeadlineAt)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"status": "Internal server error", "message": err.Error()})
		return
	}
	ctx.JSON(201, todo)
}

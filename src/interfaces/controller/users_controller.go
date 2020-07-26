package controller

import (
	"CleanArchitectureSample_golang/common"
	"CleanArchitectureSample_golang/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

// UsersController provides user manipulation controller.
type UsersController struct {
	signUpUseCase usecase.SignUpUseCase
	signInUseCase usecase.SignInUseCase
}

type signUpParams struct {
	Name                 string `json:"name" binding:"required,max=120"`
	Email                string `json:"email" binding:"required,email,max=120"`
	Password             string `json:"password" binding:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

type signInParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// NewUsersController initializes UsersController.
func NewUsersController(signUpUseCase usecase.SignUpUseCase, signInUseCase usecase.SignInUseCase) UsersController {
	return UsersController{signUpUseCase, signInUseCase}
}

// SignUp creates new user.
func (controller *UsersController) SignUp(c Context) {
	p := signUpParams{}
	if err := c.Bind(&p); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"status": "Bad request", "message": err.Error()})
		return
	}
	_, err := controller.signUpUseCase.Execute(p.Name, p.Email, p.Password, p.PasswordConfirmation)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"status": "Bad request", "message": err.Error()})
		return
	}
	c.Status(201)
}

// SignIn handles sign_in.
func (controller *UsersController) SignIn(ctx Context, session Session) {
	p := signInParams{}
	if err := ctx.Bind(&p); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"status": "Bad request", "message": err.Error()})
		return
	}
	user, err := controller.signInUseCase.Execute(p.Email, p.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"status": "Internal server error", "message": err.Error()})
		return
	}
	userID := session.Get(common.SessionUserIDKey)
	if userID == user.ID {
		ctx.JSON(302, user)
		return
	}
	session.Set(common.SessionUserIDKey, user.ID)
	session.Save()
	ctx.JSON(200, user)
}

// SignOut handles sign_out.
func (*UsersController) SignOut(ctx Context, session Session) {
	userID := session.Get(common.SessionUserIDKey)
	if userID == nil {
		ctx.JSON(401, gin.H{"status": "Unauthorized", "message": "unauthorized"})
	}
	session.Delete(common.SessionUserIDKey)
	session.Save()
	ctx.Status(204)
}

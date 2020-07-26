package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
)

// SignUpUseCase is interface of SignUpUseCaseImpl.
type SignUpUseCase interface {
	Execute(name string, email string, password string, passwordConfirmation string) (model.User, error)
}

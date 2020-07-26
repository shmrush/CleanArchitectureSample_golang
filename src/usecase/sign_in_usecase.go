package usecase

import "CleanArchitectureSample_golang/domain/model"

// SignInUseCase is interface of SignInUseCaseImpl.
type SignInUseCase interface {
	Execute(email string, password string) (model.User, error)
}

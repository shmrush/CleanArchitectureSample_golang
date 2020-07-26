package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
	"errors"
)

// SignUpUseCaseImpl provides sign_up manipulation.
type SignUpUseCaseImpl struct {
	userDaoRepository repository.UserDaoRepository
}

// NewSignUpUseCase initializes SignUpUseCaseImpl.
func NewSignUpUseCase(userDaoRepository repository.UserDaoRepository) SignUpUseCase {
	return &SignUpUseCaseImpl{userDaoRepository}
}

// Execute creates new user.
func (u *SignUpUseCaseImpl) Execute(name string, email string, password string, passwordConfirmation string) (user model.User, err error) {
	if password != passwordConfirmation {
		err = errors.New("invalid password")
	}
	user, err = u.userDaoRepository.Create(name, email, password)
	return
}

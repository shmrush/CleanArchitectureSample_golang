package usecase

import (
	"CleanArchitectureSample_golang/domain/model"
	"CleanArchitectureSample_golang/domain/repository"
)

// SignInUseCaseImpl provides sign_in manipulation.
type SignInUseCaseImpl struct {
	userDaoRepository repository.UserDaoRepository
}

// NewSignInUseCase initializes SignInUseCaseImpl.
func NewSignInUseCase(userDaoRepository repository.UserDaoRepository) SignInUseCase {
	return &SignInUseCaseImpl{userDaoRepository}
}

// Execute find user.
func (u *SignInUseCaseImpl) Execute(email string, password string) (user model.User, err error) {
	user, err = u.userDaoRepository.FindByEmailAndPassword(email, password)
	return
}
